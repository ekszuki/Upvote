package grpc

import (
	"github.com/sirupsen/logrus"
	gGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"klever.io/interview/pkg/domain/coins"
	protoCoins "klever.io/interview/pkg/protos/coins"
	protoMonitor "klever.io/interview/pkg/protos/monitors"
	"klever.io/interview/utils"
)

type Repositories struct {
	Coins coins.Repository
}

type Server struct {
	protoCoins.UnimplementedCoinServiceServer
	protoMonitor.UnimplementedVoteMonitorServer
	GRPCServer   *gGrpc.Server
	Repositories Repositories
}

func NewServer(repositories Repositories) *Server {
	logCtx := logrus.WithFields(logrus.Fields{"component": "Server", "function": "Run"})
	logCtx.Info("Starting gRPC Server")
	grpcServer := gGrpc.NewServer()
	s := &Server{
		GRPCServer:   grpcServer,
		Repositories: repositories,
	}
	protoCoins.RegisterCoinServiceServer(grpcServer, s)
	protoMonitor.RegisterVoteMonitorServer(grpcServer, s)

	if utils.GetEnv("GRPC_REFLECTION", "S") == "S" {
		reflection.Register(grpcServer)
	}

	return s
}
