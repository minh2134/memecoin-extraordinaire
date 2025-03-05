package swap

import (
	"database/sql"
	"fmt"
	"log"
)

// TODO: implement Swap:
// input: takes in a swapRequest, which should contains (sourceCurr, targetCurr) plus their amount, plus the slippage range
// search the database in limit order, find the closest match within slippage range (targetCurr, sourceCurr)
// if not found throws an error (server handles the error)
// if found, execute the trade, returns status 200 + recipient address

func Swap(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM test")
	log.Println("query successfull")
	rows.Scan()
	fmt.Print("not implemented")
}
