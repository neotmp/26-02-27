package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/neotmp/26-02-27/withdraw"
)

func Withdraw(w http.ResponseWriter, r *http.Request) {

	// Server Responds w/ data
	res := new(SRCreateTransaction)

	// Read the request body
	payload, err := io.ReadAll(r.Body)
	if err != nil {

		res.Message = "Could not create wallet"

		// marshal response
		data, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err, "Can't Marshal")
			// TO DO send some response here
		}

		w.Write(data)

		return
	}

	req := new(withdraw.WithdrawReq)

	// unmarshall payload into a struct
	err = json.Unmarshal([]byte(payload), req)
	if err != nil {
		fmt.Println(err, "Can't unmarshal data from request")
		// return err here
		w.WriteHeader(500)
		return

	}

	// make pledge
	code, err := req.Withdraw()
	if err != nil {

		w.WriteHeader(code.Code)

		return
	}

	//Data structure in case of success
	res.Message = "Requested data has been successfully created."

	// marshal response
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "Can't Marshal")
		// TO DO send some response here
	}

	w.Write(data)

}
