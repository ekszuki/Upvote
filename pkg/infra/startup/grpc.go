package startup

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"klever.io/interview/pkg/infra/grpc"
	"klever.io/interview/pkg/infra/repository/gormdb"
)

type Startup struct {
	ctx context.Context
}

func NewStartup(ctx context.Context) *Startup {
	return &Startup{
		ctx: ctx,
	}
}

func (s *Startup) Initialize() {
	var err error
	time.Local = time.UTC
	logCtx := logrus.WithFields(logrus.Fields{"component": "startup", "function": "Initialize"})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logCtx.Info("Initializing TCP Listiner")
	addr := ":9090"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logCtx.Fatalf("could not create a listener on address: %s", addr)
	}

	// we can to define that repositories we'll want to create.
	repos := createGormDBRepositories()

	logCtx.Infof("gRPC Server running on address: %s", addr)
	server := grpc.NewServer(repos)
	go func() {
		err = server.GRPCServer.Serve(listen)
		if err != nil {
			logCtx.Fatalf("could not initialize the server")
		}
	}()

	<-c

	logCtx.Infof("Waiting gRPC Server graceful shutdown ")
	server.GRPCServer.GracefulStop()
}

func createGormDBRepositories() grpc.Repositories {
	logCtx := logrus.WithFields(logrus.Fields{"component": "startup", "function": "createGormDBRepositories"})
	logCtx.Info("Creating repositories")

	repos := grpc.Repositories{
		Coins: gormdb.NewCoinRepository(),
	}
	return repos
}
