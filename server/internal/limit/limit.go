// @author Dinh Le Hoang Anh - 105508318
// @author Pham Vu Minh - 105110564
package limit

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"server/internal/blockchain"
	"server/internal/database"
	"server/internal/smartContract"
	"server/internal/swap"
)

type LimitRequest struct {
	SourceAddress string
	SourceCurr    string
	TargetCurr    string
	SourceAmount  decimal.Decimal
	Rate          decimal.Decimal
}

type LimitResult struct {
	IsMatched   bool
	SwapDetails swap.SwapResult
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
	err := validLimit.Scan(
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

func Limit(db *sql.DB, lr LimitRequest, instance *smartContract.SmartContract, client *ethclient.Client) (LimitResult, *types.Transaction, error) {
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

	insertToBook := `
		INSERT INTO limitOrders(sourceAddress, sourceAmount, rate, fromCurrency, toCurrency)
		VALUES (?, ?, ?, ?, ?);
	`
	tx, _ := db.Begin()
	var (
		err          error
		returnResult LimitResult
		result       database.LimitRow
		txc          *types.Transaction
	)
	result, returnResult.IsMatched = match(tx, lr)

	if returnResult.IsMatched {
		var transaction database.TransactionRow
		transaction.SourceAddress = result.SourceAddress
		transaction.TargetAddress = lr.SourceAddress
		transaction.SourceCurr = result.FromCurr
		transaction.TargetCurr = result.ToCurr
		if lr.SourceAmount.Mul(result.Rate).LessThan(lr.SourceAmount) { // limit order lower then using it as amount
			transaction.SourceAmount = result.SourceAmount
			transaction.TargetAmount = result.SourceAmount.Mul(result.Rate.Pow(decimal.NewFromInt(-1)))
		} else { // use limit request
			transaction.SourceAmount = lr.SourceAmount.Mul(result.Rate)
			transaction.TargetAmount = lr.SourceAmount
		}

		returnResult.SwapDetails.TradedAddress = transaction.SourceAddress
		returnResult.SwapDetails.TradedAmount = transaction.TargetAmount
		returnResult.SwapDetails.ReceivedAmount = transaction.SourceAmount
		returnResult.SwapDetails.FromCurr = transaction.TargetCurr
		returnResult.SwapDetails.ToCurr = transaction.SourceCurr

		rslt, err := tx.Exec(markPending,
			result.Id,
			transaction.SourceAddress,
			transaction.TargetAddress,
			transaction.SourceCurr,
			transaction.TargetCurr,
			transaction.SourceAmount,
			transaction.TargetAmount,
		)

		// calling Smart Contract to trade
		request := blockchain.SMTradeContract{
			SourceAddress: transaction.SourceAddress,
			TargetAddress: transaction.TargetAddress,
			SourceCurr:    transaction.SourceCurr,
			TargetCurr:    transaction.TargetCurr,
			SourceAmount:  transaction.SourceAmount,
			TargetAmount:  transaction.TargetAmount,
		}

		txc, err = blockchain.ExecuteTrade(instance, client, request)

		if err != nil {
			log.Println("limit.go: error updating the transaction")
			tx.Rollback()
			return returnResult, txc, err
		}

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
		}(txc, client, rslt)

		returnResult.SwapDetails.TradedAddress = transaction.SourceAddress
		returnResult.SwapDetails.TradedAmount = transaction.TargetAmount
		returnResult.SwapDetails.ReceivedAmount = transaction.SourceAmount
		returnResult.SwapDetails.FromCurr = transaction.TargetCurr
		returnResult.SwapDetails.ToCurr = transaction.SourceCurr

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
			return returnResult, txc, err
		}

	}

	tx.Commit()
	return returnResult, txc, err
}
