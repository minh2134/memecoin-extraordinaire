package limit

import (
	"database/sql"
	"log"

	"github.com/shopspring/decimal"

	"server/internal/database"
	"server/internal/swap"
)

// TODO: should insert new limit order into database (maybe after checking if the order is swappable?)
// maybe limit order is essentially swap with 0% slippage range?
// how to match orders? Can a partial match acceptable?
type LimitRequest struct {
	SourceAddress	string
	SourceCurr 	string
	TargetCurr	string
	SourceAmount	decimal.Decimal
	Rate 		decimal.Decimal
}

type LimitResult struct {
	IsMatched	bool
	SwapDetails	swap.SwapResult
}

func match(tx *sql.Tx, lr LimitRequest) (database.LimitRow, bool) {
	query := `
		SELECT * FROM limitOrders WHERE
			fromCurrency = ? AND
			toCurrency = ? AND
			sourceAmount >= ? AND
			rate >= ?;
	`
	reversedRate := lr.Rate.Pow(decimal.NewFromInt(-1))

	validLimit := tx.QueryRow(query,
				lr.TargetCurr,
				lr.SourceCurr,
				lr.SourceAmount.Mul(reversedRate),
				reversedRate,
			)

	var result database.LimitRow
	err:= validLimit.Scan(
			&result.Id,
			&result.SourceAddress, 
			&result.SourceAmount, 
			&result.Rate, 
			&result.FromCurr, 
			&result.ToCurr,
		) 
	if err != nil {
			return result, false
	}
	return result, true
}

func Limit(db *sql.DB, lr LimitRequest) (LimitResult, error) {
	markDone := `
		DELETE FROM limitOrders
		WHERE id = ?;

		INSERT INTO transactions (sourceAddress, targetAddress, sourceCurr, targetCurr, sourceAmount, targetAmount)
		VALUES (?, ?, ?, ?, ?, ?);
	`
	
	insertToBook := `
		INSERT INTO limitOrders(sourceAddress, sourceAmount, rate, fromCurrency, toCurrency)
		VALUES (?, ?, ?, ?, ?);
	`
	tx,_ := db.Begin()
	var (
		err error
		returnResult LimitResult
		result database.LimitRow
	)
	result, returnResult.IsMatched = match(tx, lr)
	
	
	if returnResult.IsMatched {
		var transaction database.TransactionRow
		transaction.SourceAddress 	= result.SourceAddress
		transaction.TargetAddress 	= lr.SourceAddress
		transaction.SourceCurr		= result.FromCurr
		transaction.TargetCurr 		= result.ToCurr
		if lr.SourceAmount.Mul(result.Rate).LessThan(lr.SourceAmount) { // limit order lower
			transaction.SourceAmount = result.SourceAmount
			transaction.TargetAmount = result.SourceAmount.Mul(result.Rate.Pow(decimal.NewFromInt(-1)))
		} else { // limit request lower
			transaction.SourceAmount = lr.SourceAmount.Mul(result.Rate)
			transaction.TargetAmount = lr.SourceAmount
		}
			
		// TODO: trigger some sort of smart contract here

		_,err := tx.Exec(markDone,
				result.Id,
				transaction.SourceAddress,
				transaction.TargetAddress,
				transaction.SourceCurr,
				transaction.TargetCurr,
				transaction.SourceAmount,
				transaction.TargetAmount,
			)
		if err != nil {
			log.Println("limit.go: error updating the transaction")
			tx.Rollback()
			return returnResult, err
		}

		returnResult.SwapDetails.TradedAddress 	= transaction.SourceAddress
		returnResult.SwapDetails.TradedAmount 	= transaction.TargetAmount
		returnResult.SwapDetails.ReceivedAmount = transaction.SourceAmount
		returnResult.SwapDetails.FromCurr 	= transaction.TargetCurr
		returnResult.SwapDetails.ToCurr 	= transaction.SourceCurr

	} else {
		// if no match, insert into order book for later matches
		_, err := tx.Exec(insertToBook,
				lr.SourceAddress,
				lr.SourceAmount,
				lr.Rate,
				lr.SourceCurr,
				lr.TargetCurr,
		)
		if err != nil {
			log.Println("limit.go: failed to insert to limit order book")
			tx.Rollback()
			return returnResult, err
		}
	
	tx.Commit()

	}
	return returnResult, err
}
