package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/neotmp/26-02-27/transaction"
)

func Withdrawals(w http.ResponseWriter, r *http.Request) {

	// Server Responds w/ data
	res := new(SRTransaction)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// create transaction
	t := new(transaction.Transaction)
	t.Id = id

	// request data
	if err := t.Get(); err != nil {

		res.Message = "Could not find transaction"

		data, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err, "Can't Marshal")
			// TO DO send some response here
		}

		w.Write(data)

		return

	}

	res.Message = "Requested data has been successfully retrieved."
	res.Ok = true
	res.Data = t

	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "Can't Marshal")
	}

	w.Write(data)

}
