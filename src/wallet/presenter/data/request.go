package data

type PostDoTransactionRequest struct {
	UserID     string `json:"user_id"`
	Currency   string `json:"currency"`
	Amount     int    `json:"amount"`
	Type       string `json:"type"`
	TimePlaced string `json:"time_placed"`
}
