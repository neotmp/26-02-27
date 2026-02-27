package account

type Account struct {
	Id      int     `json:"id"`
	Balance float64 `json:"balance"`
	UserId  int     `json:"userId" db:"user_id"`
}
