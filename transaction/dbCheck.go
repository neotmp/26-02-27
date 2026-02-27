package transaction

import (
	"database/sql"
	"fmt"

	"github.com/neotmp/26-02-27/database"
)

func (t *Transaction) dbCheck() (bool, error) {

	fmt.Println(t.Hash, "hash")

	q := `SELECT id,
	user_id, amount, currency, status, hash
	FROM transactions WHERE
	hash = $1`

	if err := database.DB.QueryRow(q,
		t.Hash).Scan(&t.Id,
		&t.UserId,
		&t.Amount,
		&t.Currency,
		&t.Status,
		&t.Hash,
	); err != nil {

		if err == sql.ErrNoRows {
			fmt.Println(err, "check transaction 1")
			return false, nil
		}
		fmt.Println(err, "check transaction 2")
		return false, err
	}

	return true, nil

}
