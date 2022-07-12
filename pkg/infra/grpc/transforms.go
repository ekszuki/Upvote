package grpc

import (
	"time"

	"github.com/AlekSi/pointer"
	"klever.io/interview/pkg/domain"
	protoCoin "klever.io/interview/pkg/protos/coins"
	protoMonitor "klever.io/interview/pkg/protos/monitors"
)

func toCreateCoinResponse(coin *domain.Coin) *protoCoin.CreateCoinResponse {
	return &protoCoin.CreateCoinResponse{
		ID:          int32(pointer.GetUint(coin.ID)),
		Description: coin.Description,
		Short:       coin.Short,
	}
}

func toListActiveResponse(coins []domain.Coin) *protoCoin.ActiveCoinsResponse {
	list := new(protoCoin.ActiveCoinsResponse)
	for _, c := range coins {
		list.ActiveCoins = append(
			list.ActiveCoins,
			&protoCoin.CreateCoinResponse{
				ID:          int32(pointer.GetUint(c.ID)),
				Description: c.Description,
				Short:       c.Short,
			},
		)
	}

	return list
}

func toCoinVoteMonitorResponse(coin *domain.Coin) *protoMonitor.CoinVoteMonitorResponse {
	r := &protoMonitor.CoinVoteMonitorResponse{
		ID:          int32(pointer.GetUint(coin.ID)),
		Description: coin.Description,
		Short:       coin.Short,
		Votes:       coin.Votes,
		LastUpdate:  time.Now().Format(time.RFC3339),
	}

	return r
}
