package grpc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"klever.io/interview/pkg/application"
	"klever.io/interview/pkg/domain"
	protoCoin "klever.io/interview/pkg/protos/coins"
)

func (s *Server) CreateCoin(
	ctx context.Context, req *protoCoin.CreateCoinRequest,
) (*protoCoin.CreateCoinResponse, error) {
	logCtx := logrus.WithFields(logrus.Fields{"component": "gRPC Server", "method": "CreateCoin"})
	logCtx.Info("starting gRPC request")

	if !req.ProtoReflect().IsValid() {
		logCtx.Warn("invalid create coin request")
		return &protoCoin.CreateCoinResponse{}, fmt.Errorf("invalid create coin request")
	}

	coinApp := application.NewCoinApplication(s.Repositories.Coins)
	newCoin, err := coinApp.CreateCoin(
		ctx,
		&domain.Coin{
			ID:          new(uint),
			Description: req.GetDescription(),
			Short:       req.GetShort(),
		},
	)
	if err != nil {
		return &protoCoin.CreateCoinResponse{}, fmt.Errorf("could not create new coin")
	}

	logCtx.Info("finishing gRPC request")
	return toCreateCoinResponse(newCoin), nil
}

func (s *Server) DeleteCoin(
	ctx context.Context, req *protoCoin.DeleteCoinRequest,
) (*empty.Empty, error) {
	logCtx := logrus.WithFields(logrus.Fields{"component": "gRPCServer", "method": "DeleteCoin"})
	logCtx.Info("starting gRPC request")
	if !req.ProtoReflect().IsValid() {
		logCtx.Warn("invalid delete coin request")
		return &empty.Empty{}, fmt.Errorf("invalid delete coin request")
	}

	coinApp := application.NewCoinApplication(s.Repositories.Coins)
	err := coinApp.DeleteCoin(ctx, uint(req.GetID()))

	logCtx.Info("finishing gRPC request")
	return &empty.Empty{}, err
}

func (s *Server) ListActiveCoins(
	ctx context.Context, _ *emptypb.Empty,
) (*protoCoin.ActiveCoinsResponse, error) {
	logCtx := logrus.WithFields(logrus.Fields{"component": "gRPCServer", "method": "ListActiveCoins"})
	logCtx.Info("starting gRPC request")

	coinApp := application.NewCoinApplication(s.Repositories.Coins)
	list, err := coinApp.ListActiveCoins(ctx)

	logCtx.Info("finishing gRPC request")
	return toListActiveResponse(list), err
}

func (s *Server) VoteUP(
	ctx context.Context, req *protoCoin.VoteRequest,
) (*emptypb.Empty, error) {
	logCtx := logrus.WithFields(logrus.Fields{"component": "gRPCServer", "method": "VoteUP"})
	logCtx.Info("starting gRPC request")

	if !req.ProtoReflect().IsValid() {
		logCtx.Warn("invalid coin vote request")
		return &empty.Empty{}, fmt.Errorf("invalid vote up request")
	}

	coinApp := application.NewCoinApplication(s.Repositories.Coins)
	err := coinApp.CoinVoteUP(ctx, uint(req.GetCoinID()))
	logCtx.Info("finishing gRPC Request")

	return &emptypb.Empty{}, err
}

func (s *Server) VoteDown(
	ctx context.Context, req *protoCoin.VoteRequest,
) (*emptypb.Empty, error) {
	logCtx := logrus.WithFields(logrus.Fields{"component": "gRPCServer", "method": "VoteDown"})
	logCtx.Info("starting gRPC request")

	if !req.ProtoReflect().IsValid() {
		logCtx.Warn("invalid coin vote request")
		return &empty.Empty{}, fmt.Errorf("invalid vote down request")
	}

	coinApp := application.NewCoinApplication(s.Repositories.Coins)
	err := coinApp.CoinVoteDown(ctx, uint(req.GetCoinID()))
	logCtx.Info("finishing gRPC Request")

	return &emptypb.Empty{}, err
}
