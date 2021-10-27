package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"

	"github.com/dinislamdarkhan/simple-wallet/src/pkg/errors"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(service presenter.Service) http.Handler {
	router := mux.NewRouter()
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(errors.EncodeErrorJSON),
	}

	router.Handle(
		"/v1/wallet/do-transaction",
		kithttp.NewServer(
			makePostDoTransactionEndpoint(service),
			decodeDoTransactionRequest,
			kithttp.EncodeJSONResponse,
			opts...,
		),
	).Methods("POST")

	return router
}

func decodeDoTransactionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body domain.PostDoTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, errors.DeserializeError.SetMessage(err.Error())
	}

	return body, nil
}
