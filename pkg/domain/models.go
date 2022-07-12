package domain

type Coin struct {
	ID          *uint
	Description string
	Short       string
	Votes       int64
}
