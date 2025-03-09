package swap

import (
	"database/sql"
	"log"
)

// TODO: implement Swap:
// input: takes in a swapRequest, which should contains (sourceCurr, targetCurr) plus their amount, plus the slippage range
// search the database in limit order, find the closest match within slippage range (targetCurr, sourceCurr)
// if not found throws an error (server handles the error)
// if found, execute the trade, returns status 200 + recipient address

type SwapRequest struct {
	SourceCurr 	string
	TargetCurr 	string
	Amount		float64
	SourceAddress	string
}

// struct for holding results from querying limit orders
type limitRow struct {
	id 		int 
	sourceAddress 	string
	amount		float64
	fromCurr	string
	toCurr		string
}

func Swap(db *sql.DB, sr SwapRequest) {
	validSwap := db.QueryRow("SELECT * FROM limitOrders WHERE toCurrency = ?", sr.SourceCurr)
	log.Println("swap query successful")
	
	var result limitRow;
	if err := validSwap.Scan(&result.id, &result.sourceAddress, &result.amount, &result.fromCurr, &result.toCurr); err != nil {
		// TODO: needs proper error handling
		log.Println("not good")
		log.Println(err)
		return
	}
	
	// TODO: trigger some sort of smart contract here(?) and delete the row
	log.Println("WARNING: not actually swapped, just simulate the delete for now...")
	_, err := db.Exec("DELETE FROM limitOrders WHERE id = ?", result.id)
	if err != nil {
		log.Println("Swap not successful (updating database failed)")
		log.Println(err)
	}

	log.Println("swap successfully")
	log.Println(result.sourceAddress)
}
