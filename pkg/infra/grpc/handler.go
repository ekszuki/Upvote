package grpc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"klever.io/interview/pkg/application"
	"klever.io/interview/pkg/domain"
	protoCoin "klever.io/interview/pkg/protos/coins"
)

func (s *Server) CreateCoin(
	ctx context.Context, req *protoCoin.CreateCoinRequest,
) (*protoCoin.CreateCoinResponse, error) {
	logCtx := logrus.WithFields(logrus.Fields{"component": "gRPC Server", "method": "CreateCoin"})

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

	return toCreateCoinResponse(newCoin), nil
}

func (s *Server) DeleteCoin(
	ctx context.Context, req *protoCoin.DeleteCoinRequest,
) (*empty.Empty, error) {
	logCtx := logrus.WithFields(logrus.Fields{"component": "gRPCServer", "method": "DeleteCoin"})

	if !req.ProtoReflect().IsValid() {
		logCtx.Warn("invalid delete coin request")
		return &empty.Empty{}, fmt.Errorf("invalid delete coin request")
	}
	coinApp := application.NewCoinApplication(s.Repositories.Coins)
	err := coinApp.DeleteCoin(ctx, uint(req.GetID()))

	return &empty.Empty{}, err
}
