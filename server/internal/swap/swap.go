package swap

import (
	"database/sql"
	"log"
	"server/internal/database"

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
	Rate		decimal.Decimal
	SourceAddress	string
}

type SwapResult struct {
	TradedAddress	string	
	TradedAmount	decimal.Decimal
	ReceivedAmount	decimal.Decimal
	FromCurr 	string
	ToCurr		string
}

func Swap(db *sql.DB, sr SwapRequest) (SwapResult, error) {
	var returnValues SwapResult
	tx, err := db.Begin()
	orderMatch := `
		SELECT * FROM limitOrders WHERE 
			fromCurrency = ? AND 
			toCurrency = ? AND
			rate <= ? AND
			rate >= ?;

	`
	markDone := `
		DELETE FROM limitOrders 
		WHERE id = ?;

		INSERT INTO transactions (sourceAddress, targetAddress, sourceCurr, targetCurr, sourceAmount, targetAmount)
		VALUES (?, ?, ?, ?, ?, ?);
	`
	var (
		rate, _ = sr.Rate.Pow(decimal.NewFromInt(-1)).Float64()
		
		minrate float64 = rate * (0.95)
		maxrate float64 = rate * (1.05)
	)

	validSwap := tx.QueryRow(orderMatch, 
				sr.TargetCurr,
				sr.SourceCurr,
				maxrate,
				minrate,
			)
	
	var result database.LimitRow
	if err := validSwap.Scan(
			&result.Id, 
			&result.SourceAddress, 
			&result.SourceAmount, 
			&result.Rate, 
			&result.FromCurr, 
			&result.ToCurr,
		); err != nil {
		// TODO: needs proper error handling
		log.Println("swap.go: Something wrong happened when parsing the resulted row")
		return returnValues, err
	}
	log.Println("swap.go: swap query successful")


	
	// determining the amount to be traded:
	// find the maximum targetAmount of the limitRow, compare that to the SourceAmount of the swap, whichever lower will be the final value
	var (
		transaction database.TransactionRow
		limitTargetAmount decimal.Decimal = result.SourceAmount.Mul(result.Rate)
		swapTargetAmount decimal.Decimal = sr.SourceAmount.Mul(result.Rate.Pow(decimal.NewFromInt(-1)))
	)

	if limitTargetAmount.LessThan(sr.SourceAmount) { // limit order lower
		transaction.SourceAmount = result.SourceAmount
		transaction.TargetAmount = limitTargetAmount
	} else { // swap request lower
		transaction.SourceAmount = swapTargetAmount
		transaction.TargetAmount = sr.SourceAmount
	}


	// TODO: trigger some sort of smart contract here(?) 

	// filling in the rest of the transaction
	transaction.SourceAddress	= result.SourceAddress
	transaction.TargetAddress 	= sr.SourceAddress
	transaction.SourceCurr		= result.FromCurr
	transaction.TargetCurr 		= result.ToCurr

	log.Println("WARNING: not actually swapped, just simulate the udating database for now...")
	
	
	_, err = tx.Exec(markDone, 
			result.Id,
			transaction.SourceAddress,
			transaction.TargetAddress,
			transaction.SourceCurr,
			transaction.TargetCurr,
			transaction.SourceAmount,
			transaction.TargetAmount,
		)
	if err != nil {
		log.Println("swap.go: Swap not successful (updating database failed)")
		tx.Rollback()
		return returnValues, err
	}
	
	if err = tx.Commit(); err != nil {
		log.Println("swap.go: database update failed")
		return returnValues, err
	}

	log.Println("swap successfully")


	returnValues.TradedAddress 	= transaction.SourceAddress
	returnValues.TradedAmount	= transaction.TargetAmount
	returnValues.ReceivedAmount 	= transaction.SourceAmount
	returnValues.FromCurr 		= transaction.TargetCurr
	returnValues.ToCurr 		= transaction.SourceCurr
	return returnValues, nil
}
