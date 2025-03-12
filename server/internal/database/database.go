package database

// This package holds all the relevant settings for setting a database connection
// The main package is supposed to import this package, create the connection,
// then delegate it to relevant packages to access the database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

	CREATE TABLE IF NOT EXISTS limitOrders (
		id 		INTEGER PRIMARY KEY,
		sourceAddress	TEXT NOT NULL,
		sourceAmount	INT NOT NULL,
		targetAmount	INT NOT NULL,
		fromCurrency	TEXT NOT NULL,
		toCurrency	TEXT NOT NULL,
		FOREIGN KEY(fromCurrency) REFERENCES currencies(name),
		FOREIGN KEY(toCurrency) REFERENCES currencies(name)
	);

	INSERT INTO currencies (name, fullName) VALUES 
		('BTC', 'Bitcoin'),
		('DOGE', 'Dogecoin'),
		('SHIB', 'Shiba Inu'),
		('PEPE', 'Pepe'),
		('BONK', 'Bonk');
	
	INSERT INTO limitOrders (sourceAddress, sourceAmount, targetAmount, fromCurrency, toCurrency) VALUES
		('xsasdaw231', 2599, 1499, 'BTC', 'PEPE');
	`

	_, err := d.Exec(schema)
	return err
}
