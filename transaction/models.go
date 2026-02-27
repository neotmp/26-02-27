package transaction

type Transaction struct {
	Id          int     `json:"id"`
	UserId      int     `json:"userId" db:"userId"`
	Amount      float64 `json:"amount"`
	Destination int     `json:"destination"` // 1 - deposit, 2 - withdrawal
	Currency    string  `json:"currency"`
	Status      int     `json:"status"` // 1 - pending, 2 - success, 0 - failure
	Hash        string  `json:"hash"`
}
