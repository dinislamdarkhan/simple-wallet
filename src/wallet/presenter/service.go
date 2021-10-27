package presenter

import (
	"context"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/logic"

	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
	"github.com/dinislamdarkhan/simple-wallet/src/pkg/errors"
)

type Service interface {
	PostDoTransaction(ctx context.Context, request *domain.PostDoTransactionRequest) (resp *domain.PostDoTransactionResponse, err error)
	GetWalletBalance(ctx context.Context, req *domain.GetWalletBalanceRequest) (resp *domain.GetWalletBalanceResponse, err error)
}

type service struct {
	store store.RepositoryStore
}

func New(store store.RepositoryStore) Service {
	return &service{store: store}
}

func (s *service) PostDoTransaction(ctx context.Context, req *domain.PostDoTransactionRequest) (resp *domain.PostDoTransactionResponse, err error) {
	ch := make(chan error, 1)

	go func() {
		resp, err = logic.DoTransaction(ctx, &logic.DoTransactionFacade{Store: s.store}, req)

		ch <- err
	}()

	select {
	case <-ctx.Done():
		err = errors.NetworkTimeout
	case err = <-ch:
	}

	return
}

func (s *service) GetWalletBalance(ctx context.Context, req *domain.GetWalletBalanceRequest) (resp *domain.GetWalletBalanceResponse, err error) {
	ch := make(chan error, 1)

	go func() {
		resp, err = logic.GetWalletBalance(ctx, &logic.WalletBalanceRepositoryFacade{Store: s.store}, req)

		ch <- err
	}()

	select {
	case <-ctx.Done():
		err = errors.NetworkTimeout
	case err = <-ch:
	}

	return
}
