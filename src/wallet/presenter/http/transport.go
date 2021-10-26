package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter"
	"github.com/dinislamdarkhan/simple-wallet/src/wallet/presenter/data"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(service presenter.Service, r *mux.Router) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}

	r.Handle(
		"/v1/wallet/do-transaction",
		kithttp.NewServer(
			makePostDoTransactionEndpoint(service),
			decodeDoTransactionRequest,
			kithttp.EncodeJSONResponse,
			opts...,
		),
	).Methods("POST")
}

func decodeDoTransactionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body data.PostDoTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}
