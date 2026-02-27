package account

import (
	"database/sql"
	"fmt"

	"github.com/neotmp/26-02-27/database"
)

func dbGetAccount(userId int) (*Account, error) {

	a := new(Account)

	q := `SELECT id,
	user_id, balance
	FROM accounts WHERE
	user_id = $1`

	if err := database.DB.QueryRow(q, userId).Scan(&a.Id,
		&a.UserId,
		&a.Balance,
	); err != nil {

		if err == sql.ErrNoRows {
			fmt.Println(err, "get account 1cd")
			return nil, err
		}
		fmt.Println(err, "get account 2al")
		return nil, err
	}

	return a, nil

}
