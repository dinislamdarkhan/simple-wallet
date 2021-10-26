package repository

import (
	"context"
)

type CassandraRepository interface {
	UpdateWalletAmount(ctx context.Context, currency, userID string, amount float64) error
	GetWalletAmount(ctx context.Context, currency, userID string) (float64, error)
}
