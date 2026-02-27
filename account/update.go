package account

import "github.com/neotmp/26-02-27/transaction"

func (a *Account) Update(t *transaction.Transaction) error {

	a.Balance -= t.Amount

	if err := dbUpdate(t, a); err != nil {
		return err
	}

	return nil
}
