package main

import (
	"fmt"
	//"errors"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/swap", swap)
	http.HandleFunc("/limit", limit)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func swap (w http.ResponseWriter, r *http.Request) {

}

func limit (w http.ResponseWriter, r *http.Request) {

}
