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
	dbOptsURL string 			= "file:" + dbFile + "?_journal=WAL"
	maxOpenConn int 			= 10
	maxIdleConn int 			= 5
	maxConnLifetime time.Duration 		= 5 * time.Minute
	maxIdleConnLifetime time.Duration 	= 5 * time.Minute
)

var needBootstrap = false

type TransactionRow struct {
	Id 		int
	SourceAddress	string
	TargetAddress 	string
	SourceCurr 	string
	TargetCurr 	string
	SourceAmount 	decimal.Decimal
	TargetAmount 	decimal.Decimal
}

type LimitRow struct {
	Id 		int 
	SourceAddress 	string
	SourceAmount 	decimal.Decimal
	Rate 		decimal.Decimal
	FromCurr	string
	ToCurr		string
}

func Open() (*sql.DB, error) {
	if _, err := os.Stat(dbFile); err != nil {
		needBootstrap = true
	}

	db, err := sql.Open(dbType, dbOptsURL)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)
	db.SetConnMaxIdleTime(maxIdleConnLifetime)
	db.SetConnMaxLifetime(maxConnLifetime)

	return db,err
}

func Bootstrap(d *sql.DB) error {
	if !needBootstrap {
		log.Println("dbFile already exists. Skipping the bootstrapping...")
		return nil
	}
	needBootstrap = false
	log.Println("dbFile not found. Bootstrapping...")

	// creating appropriate tables if not exists

	// TODO: needs more mock data
	schema := `
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
	insertMock := `
		INSERT INTO limitOrders (sourceAddress, 
					sourceAmount, 
					rate, 
					fromCurrency, 
					toCurrency) VALUES 
					(?, ?, ?, ?, ?);
	`

	tx, err := d.Begin()
	if err != nil {
		return err
	}
	
	var currencies = [5]string {
			"DOGE",
			"BTC",
			"SHIB",
			"BONK",
			"PEPE",
	}

	var addresses = [21]string {
		"0x8943545177806ED17B9F23F0a21ee5948eCaa776",
		"0xE25583099BA105D9ec0A67f5Ae86D90e50036425",
		"0x614561D2d143621E126e87831AEF287678B442b8",
		"0xf93Ee4Cf8c6c40b329b0c0626F28333c132CF241",
		"0x802dCbE1B1A97554B4F50DB5119E37E8e7336417",
		"0xAe95d8DA9244C37CaC0a3e16BA966a8e852Bb6D6",
		"0x2c57d1CFC6d5f8E4182a56b4cf75421472eBAEa4",
		"0x741bFE4802cE1C4b5b00F9Df2F5f179A1C89171A",
		"0xc3913d4D8bAb4914328651C2EAE817C8b78E1f4c",
		"0x65D08a056c17Ae13370565B04cF77D2AfA1cB9FA",
		"0x3e95dFbBaF6B348396E6674C7871546dCC568e56",
		"0x5918b2e647464d4743601a865753e64C8059Dc4F",
		"0x589A698b7b7dA0Bec545177D3963A2741105C7C9",
		"0x4d1CB4eB7969f8806E2CaAc0cbbB71f88C8ec413",
		"0xF5504cE2BcC52614F121aff9b93b2001d92715CA",
		"0xF61E98E7D47aB884C244E39E031978E33162ff4b",
		"0xf1424826861ffbbD25405F5145B5E50d0F1bFc90",
		"0xfDCe42116f541fc8f7b0776e2B30832bD5621C85",
		"0xD9211042f35968820A3407ac3d80C725f8F75c14",
		"0xD8F3183DEF51A987222D845be228e0Bbb932C222",
		"0xafF0CA253b97e54440965855cec0A8a2E2399896",
	}

	_, err = tx.Exec(schema)
	if err != nil {
		tx.Rollback()
		return err
	}


	
	for i:=1; i<=1000; i++ {
		addInd := rand.Intn(len(addresses))
		curr1Ind := rand.Intn(len(currencies))
		curr2Ind := curr1Ind
		for { 
			curr2Ind = rand.Intn(len(currencies))
			if curr2Ind != curr1Ind {
				break
			}
		}
		rate := (rand.Float64() + 0.001) * 100
		amount := (rand.Float64() + 0.001) * 1000

		_, err := tx.Exec(insertMock, addresses[addInd], amount, rate, currencies[curr1Ind], currencies[curr2Ind])
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	

		tx.Commit()
	return err
}
