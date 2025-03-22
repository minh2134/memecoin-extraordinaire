// @author Dinh Le Hoang Anh - 105508318
// @author Pham Vu Minh - 105110564
package database

// This package holds all the relevant settings for setting a database connection
// The main package is supposed to import this package, create the connection,
// then delegate it to relevant packages to access the database

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
)

var dbType = "sqlite3"
var dbFile string = "./database/memeextra.db"

// sets sqlite3 opts as per go-sqlite3 doc
// set journal to WAL for concurrent reads and writes
var (
	dbOptsURL           string        = "file:" + dbFile + "?_journal=WAL"
	maxOpenConn         int           = 10
	maxIdleConn         int           = 5
	maxConnLifetime     time.Duration = 5 * time.Minute
	maxIdleConnLifetime time.Duration = 5 * time.Minute
)

var needBootstrap = false

type TransactionRow struct {
	Id            int
	SourceAddress string
	TargetAddress string
	SourceCurr    string
	TargetCurr    string
	SourceAmount  decimal.Decimal
	TargetAmount  decimal.Decimal
}

type LimitRow struct {
	Id            int
	SourceAddress string
	SourceAmount  decimal.Decimal
	Rate          decimal.Decimal
	FromCurr      string
	ToCurr        string
}

// define schema and sql queries
var schema string = `
	CREATE TABLE IF NOT EXISTS currencies (
		name		TEXT PRIMARY KEY,
		fullName	TEXT
	);

	CREATE TABLE IF NOT EXISTS status (
		code 		TEXT PRIMARY KEY,
		desc 		TEXT
	);

	CREATE TABLE IF NOT EXISTS limitOrders (
		id 		INTEGER PRIMARY KEY,
		sourceAddress	TEXT NOT NULL,
		sourceAmount	REAL NOT NULL,
		rate		REAL NOT NULL,
		fromCurrency	TEXT NOT NULL,
		toCurrency	TEXT NOT NULL,
		FOREIGN KEY(fromCurrency) REFERENCES currencies(name),
		FOREIGN KEY(toCurrency) REFERENCES currencies(name)
	);

	CREATE TABLE IF NOT EXISTS transactions (
		id 		INTEGER PRIMARY KEY,
		sourceAddress 	TEXT NOT NULL,
		targetAddress	TEXT NOT NULL,
		sourceCurr	TEXT NOT NULL,
		targetCurr	TEXT NOT NULL,
		sourceAmount 	REAL NOT NULL,
		targetAmount	REAL NOT NULL,
		status 		TEXT NOT NULL,
		FOREIGN KEY(sourceCurr) REFERENCES currencies(name),
		FOREIGN KEY(targetCurr) REFERENCES currencies(name),
		FOREIGN KEY(status) REFERENCES status(code)
	);

	INSERT INTO currencies (name, fullName) VALUES 
		('BTC', 'Bitcoin'),
		('DOGE', 'Dogecoin'),
		('SHIB', 'Shiba Inu'),
		('PEPE', 'Pepe'),
		('BONK', 'Bonk');

	INSERT INTO status (code, desc) VALUES
		('PEND', 'pending transaction'),
		('DONE', 'transaction already mined and confirmed');

	`

var insertMock string = `
	INSERT INTO limitOrders (sourceAddress, 
				sourceAmount, 
				rate, 
				fromCurrency, 
				toCurrency) VALUES 
				(?, ?, ?, ?, ?);
`

func Open() (*sql.DB, error) {
	if _, err := os.Stat(dbFile); err != nil {
		needBootstrap = true
	}

	db, err := sql.Open(dbType, dbOptsURL)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)
	db.SetConnMaxIdleTime(maxIdleConnLifetime)
	db.SetConnMaxLifetime(maxConnLifetime)

	return db, err
}

func Bootstrap(d *sql.DB) error {
	// if database already exists, skip...
	if !needBootstrap {
		log.Println("dbFile already exists. Skipping the bootstrapping...")
		return nil
	}
	needBootstrap = false
	log.Println("dbFile not found. Bootstrapping...")

	tx, err := d.Begin()
	if err != nil {
		return err
	}

	var currencies = [5]string{
		"BTC",
		"DAI",
		"SNX",
	}

	var addresses = [3]string{
		"0x8019F8FA7b20c68F7CF4Bb4d74F13AB54d7f57f3",
		"0x9C0802a055D25040C233cCF7aEd8Fe9339c1b8c6",
		"0xb0d4A8965D85BE95828DEE79B5EC8fe51C9da53e",
	}

	// creating appropriate tables if not exists
	_, err = tx.Exec(schema)
	if err != nil {
		tx.Rollback()
		return err
	}

	for i := 1; i <= 100; i++ {
		addInd := rand.Intn(len(addresses))
		curr1Ind := rand.Intn(len(currencies))
		curr2Ind := curr1Ind
		for {
			curr2Ind = rand.Intn(len(currencies))
			if curr2Ind != curr1Ind {
				break
			}
		}
		rate := (rand.Float64() + 0.001) * 50
		amount := (rand.Float64() + 0.001) * 10

		_, err := tx.Exec(insertMock, addresses[addInd], amount, rate, currencies[curr1Ind], currencies[curr2Ind])
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return err
}
