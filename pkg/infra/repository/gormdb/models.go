package gormdb

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	Description string `gorm:"column:description;type:varchar;size:255"`
	Short       string `gorm:"column:short;type:varchar;size:10"`
}

func (Coin) TableName() string {
	return "coins"
}
