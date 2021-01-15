package account

import "context"

type Account struct {
	ID     string
	UserID string
	Total  float64
}

type Filter struct {
	ID     string
	UserID string
	Total  float64
}

type Store interface {
	Fetch(context.Context, Filter) (Account, error)
	FetchMany(context.Context, Filter, func(Account) error) error
	UpdateAccountTotal(context.Context, Filter) error
}

type App interface {
	Store
	StoreTransaction
}
