package swap

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"log"
	"server/internal/blockchain"
	"server/internal/database"
	"server/internal/smartContract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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
	Slippage 	decimal.Decimal
}

type SwapResult struct {
	TradedAddress	string	
	TradedAmount	decimal.Decimal
	ReceivedAmount	decimal.Decimal
	FromCurr 	string
	ToCurr		string
}

type RateRequest struct {
	SourceAddress 	string
	SourceCurr 	string
	TargetCurr 	string
}

func GetRate(db *sql.DB, req RateRequest) (decimal.Decimal, error) {
	tx, err := db.Begin()
	if err != nil {
		return decimal.NewFromInt(0), err
	}
	queryMatches := `
		SELECT rate FROM limitOrders WHERE
			sourceAddress <> ? AND
			fromCurrency = ? AND
			toCurrency = ?
		ORDER BY rate ASC;
	`

	rows, err := tx.Query(queryMatches,
				req.SourceAddress,
				req.TargetCurr,
				req.SourceCurr,
		)
	defer rows.Close()
	
	var rates = []decimal.Decimal{}
	var rate decimal.Decimal
	for rows.Next() {
		rows.Scan(&rate)
		rate = rate.Pow(decimal.NewFromInt(-1))
		rates = append(rates, rate)
	}
	if len(rates) == 0 {
		return rate, errors.New("No rates found")
	}
	// take the median
	leng := len(rates) - 1
	leng = leng/2

	return rates[leng], err
}

func Swap(db *sql.DB, sr SwapRequest, instance *smartContract.SmartContract, client *ethclient.Client) (SwapResult, *types.Transaction, error) {
	var (
		returnValues SwapResult
		txc *types.Transaction
	)
	tx, err := db.Begin()
	orderMatch := `
		SELECT * FROM limitOrders WHERE 
			sourceAddress <> ? AND
			fromCurrency = ? AND 
			toCurrency = ? AND
			rate <= ? AND
			rate >= ?
		ORDER BY rate DESC;

	`
	markPending := `
		DELETE FROM limitOrders 
		WHERE id = ?;

		INSERT INTO transactions (sourceAddress, targetAddress, sourceCurr, targetCurr, sourceAmount, targetAmount, status)
		VALUES (?, ?, ?, ?, ?, ?, 'PEND');
	`

	markDone := `
		UPDATE transactions 
		SET status = 'DONE'
		WHERE id = ?;
	`
	var (
		rate, _ = sr.Rate.Pow(decimal.NewFromInt(-1)).Float64()
		slippage, _ = sr.Slippage.Float64()
		
		minrate float64 = rate * (1.0 - slippage)
		maxrate float64 = rate * (1.0 + slippage)
	)
	
	// find if there's any match
	validSwap := tx.QueryRow(orderMatch,
				sr.SourceAddress,
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
		log.Println("swap.go: Something wrong happened when parsing the resulted row")
		return returnValues, txc, err
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
	// filling in the rest of the transaction
	transaction.SourceAddress	= result.SourceAddress
	transaction.TargetAddress 	= sr.SourceAddress
	transaction.SourceCurr		= result.FromCurr
	transaction.TargetCurr 		= result.ToCurr
	

	// calling Smart Contract to trade
	request := blockchain.SMTradeContract {
		SourceAddress: transaction.SourceAddress,
		TargetAddress: transaction.TargetAddress,
		SourceCurr: transaction.SourceCurr,
		TargetCurr: transaction.TargetCurr,
		SourceAmount: transaction.SourceAmount,
		TargetAmount: transaction.TargetAmount,
	}

	txc, err = blockchain.ExecuteTrade(instance, client, request)

	rslt, err := tx.Exec(markPending, 
			result.Id,
			transaction.SourceAddress,
			transaction.TargetAddress,
			transaction.SourceCurr,
			transaction.TargetCurr,
			transaction.SourceAmount,
			transaction.TargetAmount,
		)

	go func(txc *types.Transaction, client *ethclient.Client, rslt driver.Result) {
		// async, wait for transaction to be mined, then update the transaction
		tx2, _ := db.Begin()
		id, err1 := rslt.LastInsertId()
		if err1 != nil {
			tx2.Rollback()
			log.Println(err1)
			return
		}
		bind.WaitMined(context.Background(), client, txc)
		_, err1 = tx2.Exec(markDone, id)
		if err1 != nil {
			tx2.Rollback()
			log.Println(err1)
			return
		}
		tx2.Commit()
	} (txc, client, rslt)

	if err = tx.Commit(); err != nil {
		log.Println("swap.go: database update failed")
		return returnValues, txc, err
	}

	log.Println("swap successfully")


	returnValues.TradedAddress 	= transaction.SourceAddress
	returnValues.TradedAmount	= transaction.TargetAmount
	returnValues.ReceivedAmount 	= transaction.SourceAmount
	returnValues.FromCurr 		= transaction.TargetCurr
	returnValues.ToCurr 		= transaction.SourceCurr

	return returnValues, txc, err
}
