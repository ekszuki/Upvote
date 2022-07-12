package coins

import (
	"context"

	"klever.io/interview/pkg/domain"
)

type Repository interface {
	Save(ctx context.Context, coin *domain.Coin) (*domain.Coin, error)
	Update(ctx context.Context, coin *domain.Coin) error
	Delete(ctx context.Context, coinID uint) error
	ListActive(ctx context.Context) ([]domain.Coin, error)
	VoteUP(ctx context.Context, coinID uint) error
	VoteDown(ctx context.Context, coinID uint) error
	FindByID(ctx context.Context, coinID uint) (*domain.Coin, error)
}
