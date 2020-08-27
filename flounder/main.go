package main

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/vvatanabe/git-ha-poc/shared/metadata"

	"github.com/vvatanabe/git-ha-poc/shared/interceptor"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	pbSSH "github.com/vvatanabe/git-ha-poc/proto/ssh"
	"github.com/vvatanabe/git-ssh-test-server/gitssh"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
)

const (
	sshPort = ":2222"
)

var (
	hostPrivateKeySigner ssh.Signer
	gitSSHTransfer       *GitSSHTransfer
)

func init() {
	keyPath := "/root/host_key"

	hostPrivateKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}

	hostPrivateKeySigner, err = ssh.ParsePrivateKey(hostPrivateKey)
	if err != nil {
		panic(err)
	}

	// do not use unary interceptor
	// unaryChain := grpc_middleware.ChainUnaryClient(interceptor.XGitUserUnaryClientInterceptor, interceptor.XGitRepoUnaryClientInterceptor)
	streamChain := grpc_middleware.ChainStreamClient(interceptor.XGitUserStreamClientInterceptor, interceptor.XGitRepoStreamClientInterceptor)
	conn, err := grpc.Dial(os.Getenv("SWORDFISH_ADDR"), grpc.WithStreamInterceptor(streamChain), grpc.WithInsecure())
	if err != nil {
		panic("failed to dial: " + err.Error())
	}

	gitSSHTransfer = &GitSSHTransfer{
		client: pbSSH.NewSSHProtocolServiceClient(conn),
	}
}

func main() {

	log.SetFlags(0)
	log.SetPrefix("[flounder] ")

	sshLis, err := net.Listen("tcp", sshPort)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := gitssh.Server{
		GitRequestTransfer: func(ch ssh.Channel, req *ssh.Request, perms *ssh.Permissions, gitCmd, repoPath string) error {
			user, repo := splitRepoPath(repoPath)
			ctx := context.Background()
			ctx = metadata.ContextWithUser(ctx, user)
			ctx = metadata.ContextWithRepo(ctx, repo)
			_ = req.Reply(true, nil)
			switch gitCmd {
			case "git-receive-pack":
				return gitSSHTransfer.GitReceivePack(ctx, ch)
			case "git-upload-pack":
				return gitSSHTransfer.GitUploadPack(ctx, ch, req)
			default:
				return errors.New("unknown operation " + gitCmd)
			}
		},
		PublicKeyCallback: func(metadata ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			return &ssh.Permissions{
				Extensions: make(map[string]string),
			}, nil
		},
		Signer: hostPrivateKeySigner,
	}
	go func() {
		log.Printf("start ssh server on port%s\n", sshPort)
		if err := s.Serve(sshLis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	<-done

	log.Println("received a signal")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("failed to shutdown: %v\n", err)
	}

	log.Println("completed shutdown")
}

func splitRepoPath(repoPath string) (user, repo string) {
	splitPath := strings.Split(repoPath, "/")
	if len(splitPath) != 3 {
		return
	}
	user = splitPath[1]
	repo = strings.TrimSuffix(splitPath[2], ".git")
	return
}

type GitSSHTransfer struct {
	client pbSSH.SSHProtocolServiceClient
}

func (t *GitSSHTransfer) GitUploadPack(ctx context.Context, ch ssh.Channel, req *ssh.Request) error {

	user := metadata.UserFromContext(ctx)
	repo := metadata.RepoFromContext(ctx)

	stream, err := t.client.PostUploadPack(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err := stream.CloseSend()
		if err != nil {
			log.Println("failed to close send ", err)
		}
	}()

	go func() {
		err = stream.Send(&pbSSH.UploadPackRequest{
			Repository: &pbSSH.Repository{
				User: user,
				Repo: repo,
			},
		})
		if err != nil {
			log.Println("failed to send first ", err)
		}

		buf := make([]byte, 32*1024)

		for {
			n, err := ch.Read(buf)
			if n > 0 {
				err = stream.Send(&pbSSH.UploadPackRequest{
					Repository: &pbSSH.Repository{
						User: user,
						Repo: repo,
					},
					Data: buf[:n],
				})
				if err != nil {
					log.Println("failed to send request ", err)
					return
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("failed to read channel ", err)
				return
			}
		}
	}()

	for {
		c, err := stream.Recv()
		if c != nil {
			if len(c.Data) > 0 {
				ch.Write(c.Data)
			}
			if len(c.Err) > 0 {
				ch.Stderr().Write(c.Data)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *GitSSHTransfer) GitReceivePack(ctx context.Context, ch ssh.Channel) error {

	user := metadata.UserFromContext(ctx)
	repo := metadata.RepoFromContext(ctx)

	stream, err := t.client.PostReceivePack(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err := stream.CloseSend()
		if err != nil {
			log.Println("failed to close send ", err)
		}
	}()

	go func() {
		err = stream.Send(&pbSSH.ReceivePackRequest{
			Repository: &pbSSH.Repository{
				User: user,
				Repo: repo,
			},
		})
		if err != nil {
			log.Println("failed to send first ", err)
		}

		buf := make([]byte, 32*1024)

		for {
			n, err := ch.Read(buf)
			if n > 0 {
				err = stream.Send(&pbSSH.ReceivePackRequest{
					Repository: &pbSSH.Repository{
						User: user,
						Repo: repo,
					},
					Data: buf[:n],
				})
				if err != nil {
					log.Println("failed to send request ", err)
					return
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("failed to read channel ", err)
				return
			}
		}
	}()

	for {
		c, err := stream.Recv()
		if c != nil {
			if len(c.Data) > 0 {
				ch.Write(c.Data)
			}
			if len(c.Err) > 0 {
				ch.Stderr().Write(c.Data)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}
