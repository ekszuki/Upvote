package grpc

import (
	"github.com/AlekSi/pointer"
	"klever.io/interview/pkg/domain"
	protoCoin "klever.io/interview/pkg/protos/coins"
)

func toCreateCoinResponse(coin *domain.Coin) *protoCoin.CreateCoinResponse {
	return &protoCoin.CreateCoinResponse{
		ID:          int32(pointer.GetUint(coin.ID)),
		Description: coin.Description,
		Short:       coin.Short,
	}
}
