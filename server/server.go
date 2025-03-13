package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

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
	mux := http.NewServeMux()

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
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/trade/swap", swapHandler)
	mux.HandleFunc("/trade/limit", limitHandler)

	log.Println("Handling connection at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handler (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func swapHandler (w http.ResponseWriter, r *http.Request) {
	// TODO: create a valid SwapRequest from request body
	var swapRequest swap.SwapRequest
	switch r.Method {
		case "POST":
			// Expecting a JSON
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			err := dec.Decode(&swapRequest)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Println("Bad request")
				return
			}

			result, err := swap.Swap(db, swapRequest)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusNotFound)
			}
			json.NewEncoder(w).Encode(result)
			return
	}

	log.Println("You should not have reached here")
	/*
	// test in case mysterious stuff appears
	testSwapReq := swap.SwapRequest {
		SourceCurr: 	"PEPE",
		TargetCurr: 	"BTC",
		SourceAmount: 	1499,
		TargetAmount: 	2599,
		SourceAddress: 	"sdasdaw",
	}

	swap.Swap(db, testSwapReq)
	*/
}

func limitHandler (w http.ResponseWriter, r *http.Request) {

}
