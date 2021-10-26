package presenter

import (
	"context"

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

func (s *service) PostDoTransaction(ctx context.Context, req *data.PostDoTransactionRequest) (*data.PostDoTransactionResponse, error) {
	return nil, nil
}
