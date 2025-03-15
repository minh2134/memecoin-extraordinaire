package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// internal server packages
	"server/internal/database"
	"server/internal/limit"
	"server/internal/swap"
)

// TODO: resolves on how to take input for swap and limit order and what to return

// global database pointer to pass around in the handler
var db *sql.DB

func enableCORS(w *http.ResponseWriter) {
	// signal to clients this response is allowed for Cross-Origin Resource Sharing
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

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
	mux.HandleFunc("POST /trade/swap", swapHandler)
	mux.HandleFunc("POST /trade/limit", limitHandler)

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
	if r.URL.Path != "/trade/swap" {
		http.NotFound(w, r)
		return
	}

	enableCORS(&w)
	var swapRequest swap.SwapRequest
	// Expecting a valid JSON
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&swapRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := swap.Swap(db, swapRequest)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return 
	}
	
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
	/*
	// test in case mysterious stuff appears
	testSwapReq := swap.SwapRequest {
		SourceCurr: 	"PEPE",
		TargetCurr: 	"BTC",
		SourceAmount: 	1499,
		Rate:		0.58823529411
		SourceAddress: 	"sdasdaw",
	}

	swap.Swap(db, testSwapReq)
	*/
}

func limitHandler (w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	var limitRequest limit.LimitRequest
	// Expecting a valid JSON
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&limitRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := limit.Limit(db, limitRequest)
	if err != nil {
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
}
