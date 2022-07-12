package gormdb

import (
	"context"

	"gorm.io/gorm"
	"klever.io/interview/pkg/domain"
	"klever.io/interview/pkg/domain/coins"
	"klever.io/interview/utils"
)

type coinRepository struct {
	db *gorm.DB
}

// Delete implements coins.Repository
func (r *coinRepository) Delete(ctx context.Context, coinID uint) error {
	return r.db.Delete(&Coin{}, coinID).Error
}

// Save implements coins.Repository
func (r *coinRepository) Save(ctx context.Context, coin *domain.Coin) (*domain.Coin, error) {
	dbCoin := FromDomain(coin)
	result := r.db.Save(&dbCoin)
	return dbCoin.ToDomain(), result.Error
}

func NewCoinRepository() coins.Repository {
	return &coinRepository{
		db: getConnection(utils.GetEnv("DB_Vendor", "GORM_POSTGRES")),
	}
}
