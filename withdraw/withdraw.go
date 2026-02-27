package withdraw

import (
	"errors"
	"fmt"

	"github.com/neotmp/26-02-27/account"
	"github.com/neotmp/26-02-27/transaction"
)

func (w *WithdrawReq) Withdraw() (*ErrorRespCode, error) {

	fmt.Println(w, "here")

	erc := new(ErrorRespCode)
	erc.Code = 0

	// check balance, RACE CONDITIONS!!! FOR TEST PURPOSES ONLY
	a, err := account.GetAccount(w.UserId)
	if err != nil {
		erc.Code = 500
		return erc, err
	}

	// check userId && currency
	if w.UserId == 0 || w.Currency != "USDT" || (w.Destination != 1 && w.Destination != 2) {
		erc.Code = 400
		return erc, errors.New("wrong data")
	}

	// check negative amount
	if w.Amount <= 0 {
		erc.Code = 400
		return erc, errors.New("wrong amount")
	}

	// check balance
	if w.Amount > a.Balance {
		erc.Code = 409
		return erc, errors.New("insufficient funds")
	}

	// check key + payload
	t := new(transaction.Transaction)
	t.Amount = w.Amount
	t.UserId = w.UserId
	t.Currency = w.Currency
	t.Hash = w.IdemKey

	f, err := t.Check()
	if err != nil {
		erc.Code = 500
		return erc, err
	}

	if f && w.UserId == t.UserId && w.Amount == t.Amount && w.Currency == t.Currency {
		return erc, nil
	} else if f && (w.UserId != t.UserId || w.Amount != t.Amount || w.Currency != t.Currency) {
		erc.Code = 422
		return erc, errors.New("error")
	} else {
		t.Status = 1

		// change users balance here via DB transaction

		if err := a.Update(t); err != nil {
			erc.Code = 500
			return erc, err
		}

		return erc, nil
	}

}
