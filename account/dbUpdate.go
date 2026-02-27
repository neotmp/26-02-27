package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/neotmp/26-02-27/database"
	"github.com/neotmp/26-02-27/transaction"
)

// create transaction and update user's balance in one transaction
func dbUpdate(t *transaction.Transaction, a *Account) error {
	//First You begin a transaction with a call to database.DB.Begin()
	ctx := context.Background()
	tx, err := database.DB.BeginTx(ctx, nil)
	if err != nil {
		return errors.New(err.Error())
	}

	q := `INSERT INTO transactions (user_id, amount, currency, status, hash)
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	//	First part of transaction
	if err = tx.QueryRowContext(ctx, q, t.UserId,
		t.Amount,
		t.Currency,
		t.Status,
		t.Hash).Scan(&t.Id); err != nil {
		tx.Rollback()
		fmt.Println(err, "roll-back 1")
		return err
	}

	q2 := "UPDATE accounts SET balance = $2 WHERE id = $1 RETURNING id, balance"

	//	Second part of transaction
	if err := tx.QueryRowContext(ctx, q2, a.Id, a.Balance).Scan(&a.Id, &a.Balance); err != nil {
		tx.Rollback()
		fmt.Println("roll-back 2")
		return errors.New(err.Error())
	}

	//	close the transaction with a Commit() or Rollback() method on the resulting Tx variable.
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
