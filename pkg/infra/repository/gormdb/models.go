package gormdb

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	Description string `gorm:"column:description;type:varchar;size:255"`
	Short       string `gorm:"column:short;type:varchar;size:10"`
	Votes       int64  `gorm:"column:votes;type:bigint;default:0"`
}

func (Coin) TableName() string {
	return "coins"
}
