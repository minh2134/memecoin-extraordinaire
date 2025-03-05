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

// TODO: define a global pointer variable here for the DB
var db *sql.DB

func main() {
	var err error
	db, err = database.Open()

	if (err != nil) { log.Fatal("Opening database failed.") }
	log.Println("Connection to database successful.")

	defer db.Close()

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
	swap.Swap(db)
}

func limitHandler (w http.ResponseWriter, r *http.Request) {

}
