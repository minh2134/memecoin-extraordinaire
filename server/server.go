package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// internal server packages
	"server/internal/database"
	"server/internal/swap"
)

// TODO: limit wrong url (default to a notfound response)
// TODO: resolves on how to take input for swap and limit order and what to return

// global database pointer to pass around in the handler
var db *sql.DB

func main() {
	var err error

	// establishing DB availablity
	db, err = database.Open()

	if (err != nil) { 
		log.Fatal("Opening database failed.") 
	}
	log.Println("Opened the database successfully.")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Establishing a connection failed.")
	}
	log.Println("Connected to the database successfully.")
	
	err = database.Bootstrap(db)

	// Router
	http.HandleFunc("/", handler)
	http.HandleFunc("/trade/swap", swapHandler)
	http.HandleFunc("/trade/limit", limitHandler)

	log.Println("Handling connection at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func swapHandler (w http.ResponseWriter, r *http.Request) {
	// TODO: create a valid SwapRequest from request body
	testSwapReq := swap.SwapRequest {
		SourceCurr: 	"PEPE",
		TargetCurr: 	"BTC",
		Amount: 	14.99,
		SourceAddress: 	"sdasdaw",
	}

	swap.Swap(db, testSwapReq)
}

func limitHandler (w http.ResponseWriter, r *http.Request) {

}
