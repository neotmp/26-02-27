package transaction

import (
	"database/sql"
	"fmt"

	"github.com/neotmp/26-02-27/database"
)

func (t *Transaction) dbGet() error {

	q := `SELECT id,
	user_id, amount, currency, status
	FROM transactions WHERE
	id = $1`

	if err := database.DB.QueryRow(q,
		t.Id).Scan(&t.Id,
		&t.UserId,
		&t.Amount,
		&t.Currency,
		&t.Status,
	); err != nil {

		if err == sql.ErrNoRows {
			fmt.Println(err, "get transaction 1")
			return err
		}
		fmt.Println(err, "get transaction 2")
		return err
	}

	return nil

}
