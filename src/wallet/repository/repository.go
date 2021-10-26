package repository

import (
	"context"
	"time"
)

type CassandraRepository interface {
	UpdateWalletAmount(ctx context.Context, currency, userID string, amount float64, updatedTime time.Time) error
	GetWalletAmount(ctx context.Context, currency, userID string) (float64, error)
}
