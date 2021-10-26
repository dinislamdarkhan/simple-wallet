package data

type PostDoTransactionRequest struct {
	UserID     string `json:"user_id" validate:"required"`
	Currency   string `json:"currency" validate:"required,alpha"`
	Amount     int    `json:"amount" validate:"required,gt=0"`
	Type       string `json:"type" validate:"required,alpha"`
	TimePlaced string `json:"time_placed" validate:"required"`
}
