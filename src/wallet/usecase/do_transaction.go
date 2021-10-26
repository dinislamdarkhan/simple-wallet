package usecase

import (
	"context"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter/data"
	"github.com/shopspring/decimal"
)

type DoTransactionRepository interface {
	GetWalletAmountCassandra(ctx context.Context, currency, userID string) (float64, error)
	UpdateWalletAmountCassandra(ctx context.Context, currency, userID string, amount decimal.Decimal) error
}

type DoTransactionFacade struct {
	Store store.RepositoryStore
}

func (f *DoTransactionFacade) GetWalletAmountCassandra(ctx context.Context, currency, userID string) (float64, error) {
	return f.Store.WalletCassandra().GetWalletAmount(ctx, currency, userID)
}

func (f *DoTransactionFacade) UpdateWalletAmountCassandra(ctx context.Context, currency, userID string, amount float64, updatedTime time.Time) error {
	return f.Store.WalletCassandra().UpdateWalletAmount(ctx, currency, userID, amount, updatedTime)
}

func DoTransaction(ctx context.Context, repo DoTransactionRepository, req *data.PostDoTransactionRequest) (*data.PostDoTransactionResponse, error) {
	return nil, nil
}
