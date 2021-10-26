package usecase

import (
	"context"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter/data"
)

const (
	respMessage = "Successfully changed wallet amount"
)

type DoTransactionRepository interface {
	GetWalletAmountCassandra(ctx context.Context, currency, userID string) (float64, error)
	UpdateWalletAmountCassandra(ctx context.Context, currency, userID string, amount float64, updatedTime time.Time) error
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

func DoTransaction(ctx context.Context, repo DoTransactionRepository, req *data.PostDoTransactionRequest) (resp *data.PostDoTransactionResponse, err error) {
	amount, err := repo.GetWalletAmountCassandra(ctx, req.Currency, req.UserID)
	if err != nil {
		return nil, errors.CassandraReadError.SetMessage(err.Error())
	}
	currentAmount := decimal.NewFromFloat(amount)
	reqAmount := decimal.NewFromFloat(req.Amount)

	if req.Type == "deposit" {
		currentAmount = currentAmount.Add(reqAmount)
	} else if req.Type == "withdrawal" {
		currentAmount = currentAmount.Sub(reqAmount)
	}

	utc := time.Now().UTC()
	amount, _ = currentAmount.Float64()
	err = repo.UpdateWalletAmountCassandra(ctx, req.Currency, req.UserID, amount, utc)
	if err != nil {
		return nil, errors.CassandraSaveError.SetMessage(err.Error())
	}

	resp.Message = respMessage
	return resp, nil
}
