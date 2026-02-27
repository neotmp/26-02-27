package api

import "github.com/neotmp/26-02-27/transaction"

type SRCreateTransaction struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type SRTransaction struct {
	Ok      bool                     `json:"ok"`
	Message string                   `json:"message"`
	Data    *transaction.Transaction `json:"data"`
}
