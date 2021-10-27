package usecase

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"github.com/araddon/dateparse"
	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
	"github.com/dinislamdarkhan/simple-wallet/src/pkg/errors"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter/data"
	"github.com/shopspring/decimal"
)

const (
	respMessage        = "Successfully changed wallet amount"
	currencyError      = "Incorrect currency, supports only USD,EUR"
	operationTypeError = "Incorrect type, supports only deposit,withdrawal"
)

type DoTransactionRepository interface {
	GetWalletAmountCassandra(ctx context.Context, currency, userID string) (float64, error)
	UpdateWalletAmountCassandra(ctx context.Context, currency, userID string, amount float64, updatedTime time.Time) error
	CheckAmountExists(ctx context.Context, currency, userID string) (count int, err error)
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

func (f *DoTransactionFacade) CheckAmountExists(ctx context.Context, currency, userID string) (count int, err error) {
	return f.Store.WalletCassandra().CheckAmountExists(ctx, currency, userID)
}

func DoTransaction(ctx context.Context, repo DoTransactionRepository, req *data.PostDoTransactionRequest) (resp *data.PostDoTransactionResponse, err error) {
	logger := logrus.WithContext(ctx)
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		logger.Error(err)
		return nil, errors.BadRequest.SetMessage(err.Error())
	}

	if !data.CurrencyMap[req.Currency] {
		logger.Error(currencyError)
		return nil, errors.BadRequest.SetMessage(currencyError)
	}

	if !data.OperationTypeMap[req.Type] {
		logger.Error(operationTypeError)
		return nil, errors.BadRequest.SetMessage(operationTypeError)
	}

	amount := 0.0
	count, err := repo.CheckAmountExists(ctx, req.Currency, req.UserID)
	if err != nil {
		return nil, err
	}

	if count != 0 {
		amount, err = repo.GetWalletAmountCassandra(ctx, req.Currency, req.UserID)
		if err != nil {
			return nil, err
		}
	}

	convertedAmount := decimal.NewFromFloat(amount)
	reqAmount := decimal.NewFromFloat(req.Amount)

	if req.Type == "deposit" {
		convertedAmount = convertedAmount.Add(reqAmount)
	} else if req.Type == "withdrawal" {
		convertedAmount = convertedAmount.Sub(reqAmount)
	}

	updatedAmount, _ := convertedAmount.Float64()
	formattedTime, err := dateparse.ParseAny(req.TimePlaced)
	if err != nil {
		logger.Error(err)
		return nil, errors.DeserializeError.SetMessage(err.Error())
	}

	err = repo.UpdateWalletAmountCassandra(ctx, req.Currency, req.UserID, updatedAmount, formattedTime)
	if err != nil {
		return nil, err
	}

	resp = &data.PostDoTransactionResponse{Message: respMessage}

	return
}
