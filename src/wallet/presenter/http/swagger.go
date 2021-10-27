package http

import "github.com/dinislamdarkhan/simple-wallet/src/wallet/domain"

// swagger:parameters PostDoTransactionRequest
type PostDoTransactionRequest struct {
	// in:body
	Body domain.PostDoTransactionRequest
}

// swagger:parameters PostDoTransactionResponse
type PostDoTransactionResponse struct {
	// in:body
	Body domain.PostDoTransactionResponse
}
