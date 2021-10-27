package http

import (
	"context"
	"time"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
	"github.com/go-kit/kit/endpoint"
)

// swagger:route POST /v1/wallet/do-transaction Wallet PostDoTransactionRequest
// Сделать транзакцию
// responses:
// 200: PostDoTransactionResponse
func makePostDoTransactionEndpoint(s presenter.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.PostDoTransactionRequest)
		ctx, ctxCancel := context.WithTimeout(ctx, time.Minute)
		defer ctxCancel()

		resp, err := s.PostDoTransaction(ctx, &req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
