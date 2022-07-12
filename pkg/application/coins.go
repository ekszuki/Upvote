package application

import (
	"context"

	"github.com/sirupsen/logrus"
	"klever.io/interview/pkg/domain"
	"klever.io/interview/pkg/domain/coins"
)

type CoinApplication struct {
	coinRepository coins.Repository
}

func NewCoinApplication(coinRep coins.Repository) *CoinApplication {
	return &CoinApplication{
		coinRepository: coinRep,
	}
}

func (a *CoinApplication) CreateCoin(
	ctx context.Context,
	domainCoin *domain.Coin,
) (*domain.Coin, error) {
	logCtx := logrus.WithFields(logrus.Fields{"component": "CoinApplication", "method": "CreateCoin"})

	newCoin, err := a.coinRepository.Save(ctx, domainCoin)
	if err != nil {
		logCtx.Error("could not register new coin: %v", err)
	}

	return newCoin, err
}

func (a *CoinApplication) DeleteCoin(
	ctx context.Context, coinID uint,
) error {
	logCtx := logrus.WithFields(logrus.Fields{"component": "CoinApplication", "method": "DeleteCoin"})
	err := a.coinRepository.Delete(ctx, coinID)
	if err != nil {
		logCtx.Error("could not delete coin id %d error: %v", coinID, err)
	}

	return err
}
