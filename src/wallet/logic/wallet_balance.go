package logic

import (
	"context"

	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"
)

type WalletBalanceRepository interface{}

type WalletBalanceRepositoryFacade struct {
	Store store.RepositoryStore
}

func GetWalletBalance(ctx context.Context, repo WalletBalanceRepository, req *domain.GetWalletBalanceRequest) (resp *domain.GetWalletBalanceResponse, err error) {
	return nil, err
}
