package gormdb

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"klever.io/interview/pkg/domain"
	"klever.io/interview/pkg/domain/coins"
	"klever.io/interview/utils"
)

type coinRepository struct {
	db *gorm.DB
}

// VoteUP implements coins.Repository
func (r *coinRepository) VoteUP(ctx context.Context, coinID uint) error {
	dbOpe := r.db.Model(&Coin{}).
		Where("id = ?", coinID).
		Update("votes", gorm.Expr("votes + ?", 1))

	if dbOpe.Error != nil {
		return dbOpe.Error
	}

	if dbOpe.RowsAffected <= 0 {
		return fmt.Errorf("Coin not found")
	}

	return nil
}

// ListActive implements coins.Repository
func (r *coinRepository) ListActive(ctx context.Context) ([]domain.Coin, error) {
	var dbCoins = []Coin{}
	err := r.db.Order("id").Find(&dbCoins).Error

	coins := []domain.Coin{}
	for _, c := range dbCoins {
		coins = append(coins, *c.ToDomain())
	}

	return coins, err
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
