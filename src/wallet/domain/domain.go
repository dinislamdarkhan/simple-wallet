package domain

var CurrencyMap = map[string]bool{
	"EUR": true,
	"USD": true,
}

var OperationTypeMap = map[string]bool{
	"deposit":    true,
	"withdrawal": true,
}

type Account struct {
	Currency    string  `json:"currency"`
	Amount      float64 `json:"amount"`
	UpdatedTime string  `json:"updated_time"`
}
