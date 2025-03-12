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
	SourceAmount	int
	TargetAmount	int
	SourceAddress	string
}

// struct for holding results from querying limit orders
type limitRow struct {
	id 		int 
	sourceAddress 	string
	sourceAmount 	int
	targetAmount	int
	fromCurr	string
	toCurr		string
}

func Swap(db *sql.DB, sr SwapRequest) error {
	query := `
		SELECT * FROM limitOrders WHERE 
			fromCurrency = ? AND 
			toCurrency = ? AND
			targetAmount <= ? AND
			sourceAmount/CAST(targetAmount AS REAL) <= ? AND
			sourceAmount/CAST(targetAmount AS REAL) >= ?;

	`
	var (
		rate float64 = float64(sr.TargetAmount)/float64(sr.SourceAmount)
		minrate float64 = rate * (0.95)
		maxrate float64 = rate * (1.05)
	)

	validSwap := db.QueryRow(query, 
				sr.TargetCurr,
				sr.SourceCurr,
				sr.SourceAmount,
				maxrate,
				minrate,
			)
	
	var result limitRow;
	if err := validSwap.Scan(
			&result.id, 
			&result.sourceAddress, 
			&result.sourceAmount, 
			&result.targetAmount, 
			&result.fromCurr, 
			&result.toCurr,
		); err != nil {
		// TODO: needs proper error handling
		log.Println("swap.go: Something wrong happened when parsing the resulted row")
		return err
	}
	log.Println("swap.go: swap query successful")

	// TODO: trigger some sort of smart contract here(?) and delete the row
	log.Println("WARNING: not actually swapped, just simulate the delete for now...")
	_, err := db.Exec("DELETE FROM limitOrders WHERE id = ?", result.id)
	if err != nil {
		log.Println("swap.go: Swap not successful (updating database failed)")
		return err
	}

	log.Println("swap successfully")
	return nil
}
