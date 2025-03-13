package swap

import (
	"database/sql"
	"log"
	"github.com/shopspring/decimal"
)

// TODO: implement Swap:
// input: takes in a swapRequest, which should contains (sourceCurr, targetCurr) plus their amount, plus the slippage range
// search the database in limit order, find the closest match within slippage range (targetCurr, sourceCurr)
// if not found throws an error (server handles the error)
// if found, execute the trade, returns status 200 + recipient address

type SwapRequest struct {
	SourceCurr 	string
	TargetCurr 	string
	SourceAmount	decimal.Decimal
	TargetAmount	decimal.Decimal
	SourceAddress	string
}

type SwapResult struct {
	Address 	string
	TradedAmount	decimal.Decimal 
	ReceivedAmount	decimal.Decimal
}

// struct for holding results from querying limit orders
type limitRow struct {
	id 		int 
	sourceAddress 	string
	sourceAmount 	decimal.Decimal
	targetAmount	decimal.Decimal
	fromCurr	string
	toCurr		string
	status 		string
}

func Swap(db *sql.DB, sr SwapRequest) (SwapResult, error) {
	tx, err := db.Begin()

	var swapResult SwapResult
	orderMatch := `
		SELECT * FROM limitOrders WHERE 
			fromCurrency = ? AND 
			toCurrency = ? AND
			targetAmount = ? AND
			sourceAmount/targetAmount <= ? AND
			sourceAmount/targetAmount >= ? AND
			status = 'AVAIL';

	`
	markDone := `
		UPDATE limitOrders 
		SET status = 'DONE'
		WHERE id = ?;
	`
	var (
		rate, _ = sr.TargetAmount.Div(sr.SourceAmount).Float64()
		
		minrate float64 = rate * (0.95)
		maxrate float64 = rate * (1.05)
	)

	validSwap := tx.QueryRow(orderMatch, 
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
			&result.status,
		); err != nil {
		// TODO: needs proper error handling
		log.Println("swap.go: Something wrong happened when parsing the resulted row")
		return swapResult, err
	}
	log.Println("swap.go: swap query successful")

	// TODO: trigger some sort of smart contract here(?) and update the row status

	swapResult.Address = result.sourceAddress
	swapResult.TradedAmount = sr.SourceAmount 
	swapResult.ReceivedAmount = result.sourceAmount

	log.Println("WARNING: not actually swapped, just simulate the udating status for now...")

	_, err = tx.Exec(markDone, result.id)
	if err != nil {
		log.Println("swap.go: Swap not successful (updating database failed)")
		tx.Rollback()
		return swapResult, err
	}
	
	if err = tx.Commit(); err != nil {
		log.Println("swap.go: database update failed")
		return swapResult, err
	}

	log.Println("swap successfully")
	return swapResult, nil
}
