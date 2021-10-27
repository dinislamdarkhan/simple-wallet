package repository

import (
	"context"
	"time"
)

type CassandraRepository interface {
	UpdateWalletAmount(ctx context.Context, currency, userID string, amount float64, updatedTime time.Time) (err error)
	GetWalletAmount(ctx context.Context, currency, userID string) (amount float64, err error)
	GetWalletAmountAndTime(ctx context.Context, currency, userID string) (amount float64, time time.Time, err error)
	CheckAmountExists(ctx context.Context, currency, userID string) (count int, err error)
}
