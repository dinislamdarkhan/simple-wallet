package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"
)

type WalletBalanceRepository interface {
	CheckAmountExistsCassandra(ctx context.Context, currency, userID string) (count int, err error)
	GetWalletAmountAndTimeCassandra(ctx context.Context, currency, userID string) (amount float64, time time.Time, err error)
}

type WalletBalanceRepositoryFacade struct {
	Store store.RepositoryStore
}

func (f *WalletBalanceRepositoryFacade) GetWalletAmountAndTimeCassandra(ctx context.Context, currency, userID string) (amount float64, time time.Time, err error) {
	return f.Store.WalletCassandra().GetWalletAmountAndTime(ctx, currency, userID)
}

func (f *WalletBalanceRepositoryFacade) CheckAmountExistsCassandra(ctx context.Context, currency, userID string) (count int, err error) {
	return f.Store.WalletCassandra().CheckAmountExists(ctx, currency, userID)
}

func WalletBalance(ctx context.Context, repo WalletBalanceRepository, req *domain.GetWalletBalanceRequest) (resp *domain.GetWalletBalanceResponse, err error) {
	logger := logrus.WithContext(ctx)
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		logger.Error(err)
		return nil, errors.BadRequest.SetMessage("Empty user id from request")
	}

	var userCurrencies []string
	for currency := range domain.CurrencyMap {
		count, err := repo.CheckAmountExistsCassandra(ctx, currency, req.UserID)
		if err != nil {
			return nil, err
		}

		if count != 0 {
			userCurrencies = append(userCurrencies, currency)
		}
	}

	if len(userCurrencies) == 0 {
		return &domain.GetWalletBalanceResponse{Wallet: nil}, nil
	}

	var account domain.Account
	var respAccounts []domain.Account
	for _, currency := range userCurrencies {
		amount, updatedTime, err := repo.GetWalletAmountAndTimeCassandra(ctx, currency, req.UserID)
		if err != nil {
			return nil, err
		}

		stringAmount := fmt.Sprintf("%.3f", amount)
		formattedAmount, _ := strconv.ParseFloat(stringAmount, 64)

		account.Amount = formattedAmount
		account.Currency = currency
		account.UpdatedTime = updatedTime.String()
		respAccounts = append(respAccounts, account)
	}

	return &domain.GetWalletBalanceResponse{Wallet: respAccounts}, err
}
