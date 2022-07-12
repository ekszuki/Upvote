package gormdb

import (
	"github.com/AlekSi/pointer"
	"klever.io/interview/pkg/domain"
)

func (c *Coin) ToDomain() *domain.Coin {
	domainCoin := new(domain.Coin)

	domainCoin.ID = pointer.ToUint(c.ID)
	domainCoin.Description = c.Description
	domainCoin.Short = c.Short
	domainCoin.Votes = c.Votes

	return domainCoin
}

func FromDomain(domainCoin *domain.Coin) *Coin {
	coin := new(Coin)
	if domainCoin.ID != nil {
		coin.ID = pointer.GetUint(domainCoin.ID)
	}

	coin.Description = domainCoin.Description
	coin.Short = domainCoin.Short
	coin.Votes = domainCoin.Votes

	return coin
}
