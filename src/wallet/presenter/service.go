package presenter

import (
	"context"

	"github.com/dinislamdarkhan/simple-wallet/src/app/store"
)

type Service interface {
	PostDoTransaction(ctx context.Context, request *req.PostDoTransactionRequest) (*res.PostDoTransactionResponse, error)
}

type service struct {
	store store.RepositoryStore
}

func New(store store.RepositoryStore) Service {
	return &service{store: store}
}

func (s *service) PostDoTransaction(ctx context.Context, req *req.PostDoTransactionRequest) (*res.PostDoTransactionResponse, error) {
	return nil, nil
}
