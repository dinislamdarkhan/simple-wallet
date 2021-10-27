package domain

type PostDoTransactionResponse struct {
	Message string `json:"message"`
}

type GetWalletBalanceResponse struct {
	Wallet []Account `json:"wallet"`
}
