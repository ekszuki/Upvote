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
		logCtx.Errorf("could not register new coin: %v", err)
	}

	return newCoin, err
}

func (a *CoinApplication) DeleteCoin(
	ctx context.Context, coinID uint,
) error {
	logCtx := logrus.WithFields(logrus.Fields{"component": "CoinApplication", "method": "DeleteCoin"})
	err := a.coinRepository.Delete(ctx, coinID)
	if err != nil {
		logCtx.Errorf("could not delete coin id %d error: %v", coinID, err)
	}

	return err
}

func (a *CoinApplication) ListActiveCoins(ctx context.Context) ([]domain.Coin, error) {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "CoinApplication", "method": "ListActiveCoins"},
	)

	list, err := a.coinRepository.ListActive(ctx)
	if err != nil {
		logCtx.Errorf("could not get active coins %v", err)
	}

	return list, err
}

func (a *CoinApplication) CoinVoteUP(ctx context.Context, coinID uint) error {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "CoinApplication", "method": "CoinVoteUP"},
	)
	err := a.coinRepository.VoteUP(ctx, coinID)
	if err != nil {
		logCtx.Errorf("could not vote on coin id: %d, error: %v", coinID, err)
	}

	return err
}

func (a *CoinApplication) CoinVoteDown(ctx context.Context, coinID uint) error {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "CoinApplication", "method": "CoinVoteDown"},
	)
	err := a.coinRepository.VoteDown(ctx, coinID)
	if err != nil {
		logCtx.Errorf("could not vote on coin id: %d, error: %v", coinID, err)
	}

	return err
}

func (a CoinApplication) FindByID(ctx context.Context, coinID uint) (*domain.Coin, error) {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "CoinApplication", "method": "FindByID"},
	)

	coin, err := a.coinRepository.FindByID(ctx, coinID)
	if err != nil {
		logCtx.Error("could not get coin %v", err)
	}

	return coin, err
}
