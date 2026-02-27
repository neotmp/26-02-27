package main

import (
	"log"
	"net/http"

	"github.com/neotmp/26-02-27/api"
	"github.com/neotmp/26-02-27/database"
)

func main() {

	database.Connect()
	//
	mux := http.NewServeMux()
	// Routes

	mux.HandleFunc("/v1/withdrawals", api.Withdraw)
	mux.HandleFunc("GET /v1/withdrawals/{id}", api.Withdrawals)

	log.Print("starting server on :4444")
	err := http.ListenAndServe(":4444", mux)
	log.Fatal(err)
}
