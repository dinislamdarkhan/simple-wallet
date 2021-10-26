package http

import (
	"context"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter/data"
	"github.com/go-kit/kit/endpoint"
)

func makePostDoTransactionEndpoint(s presenter.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(data.PostDoTransactionRequest)
		ctx, ctxCancel := context.WithTimeout(ctx, time.Minute)
		defer ctxCancel()

		resp, err := s.PostDoTransaction(ctx, &req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}