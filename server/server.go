package main

import (
	"fmt"
	//"errors"
	"log"
	"net/http"
	
	// internal server packages
	"server/internal/swap"
)

// TODO: limit wrong url (default to a notfound response)
// TODO: resolves on how to take input for swap and limit order and what to return

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/swap", swapHandler)
	http.HandleFunc("/limit", limitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func swapHandler (w http.ResponseWriter, r *http.Request) {
	swap.Swap()
}

func limitHandler (w http.ResponseWriter, r *http.Request) {

}
