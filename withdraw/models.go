package withdraw

type WithdrawReq struct {
	UserId      int     `json:"userId" db:"user_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Destination int     `json:"destination"`
	IdemKey     string  `json:"idemKey" db:"idem_key"`
}

type WithdrawResp struct{}

type ErrorRespCode struct {
	Code int
}
