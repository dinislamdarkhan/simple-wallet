package presenter

import (
	"context"

	"github.com/dinislamdarkhan/simple-wallet/src/pkg/errors"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/usecase"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter/data"

	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
)

type Service interface {
	PostDoTransaction(ctx context.Context, request *data.PostDoTransactionRequest) (*data.PostDoTransactionResponse, error)
}

type service struct {
	store store.RepositoryStore
}

func New(store store.RepositoryStore) Service {
	return &service{store: store}
}

func (s *service) PostDoTransaction(ctx context.Context, req *data.PostDoTransactionRequest) (response *data.PostDoTransactionResponse, err error) {
	ch := make(chan error, 1)

	go func() {
		response, err = usecase.DoTransaction(ctx, &usecase.DoTransactionFacade{Store: s.store}, req)

		ch <- err
	}()

	select {
	case <-ctx.Done():
		err = errors.NetworkTimeout
	case err = <-ch:
	}

	return
}
