package coins

import (
	"context"

	"klever.io/interview/pkg/domain"
)

type Repository interface {
	Save(ctx context.Context, coin *domain.Coin) (*domain.Coin, error)
	Delete(ctx context.Context, coinID uint) error
	ListActive(ctx context.Context) ([]domain.Coin, error)
}
