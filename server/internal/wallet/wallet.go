package wallet

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
)

// Wallet represents a cryptocurrency wallet
type Wallet struct {
	Address string
	Name    string
}

// Balance represents a currency balance for a wallet
type Balance struct {
	Currency string
	Amount   int // Amount * 100 to handle 2 decimal places as integers
}

// CreateWallet creates a new wallet and returns its address
func CreateWallet(db *sql.DB, name string) (string, error) {
	// Generate a random wallet address
	addrBytes := make([]byte, 20) // 20 bytes for a standard Ethereum-like address
	if _, err := rand.Read(addrBytes); err != nil {
		return "", err
	}
	address := "0x" + hex.EncodeToString(addrBytes)

	// Insert into database
	_, err := db.Exec("INSERT INTO wallets (address, name) VALUES (?, ?)", address, name)
	if err != nil {
		return "", err
	}

	log.Printf("Created new wallet: %s for %s", address, name)
	return address, nil
}

// GetWallet retrieves a wallet by its address
func GetWallet(db *sql.DB, address string) (*Wallet, error) {
	row := db.QueryRow("SELECT address, name FROM wallets WHERE address = ?", address)

	var wallet Wallet
	if err := row.Scan(&wallet.Address, &wallet.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("wallet not found")
		}
		return nil, err
	}

	return &wallet, nil
}

// AddFunds adds the specified amount of a currency to a wallet
func AddFunds(db *sql.DB, address string, currency string, amount int) error {
	// Check if wallet exists
	_, err := GetWallet(db, address)
	if err != nil {
		return err
	}

	// Check if currency is valid
	var exists bool
	err = db.QueryRow("SELECT 1 FROM currencies WHERE name = ?", currency).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid currency")
		}
		return err
	}

	// Add funds (using UPSERT pattern)
	_, err = db.Exec(`
        INSERT INTO wallet_balances (wallet_address, currency, balance)
        VALUES (?, ?, ?)
        ON CONFLICT(wallet_address, currency) 
        DO UPDATE SET balance = balance + ?
    `, address, currency, amount, amount)

	if err != nil {
		return err
	}

	log.Printf("Added %d of %s to wallet %s", amount, currency, address)
	return nil
}

// GetBalance gets the balance of a specific currency for a wallet
func GetBalance(db *sql.DB, address string, currency string) (int, error) {
	// Check if wallet exists
	_, err := GetWallet(db, address)
	if err != nil {
		return 0, err
	}

	var balance int
	err = db.QueryRow(`
        SELECT balance FROM wallet_balances 
        WHERE wallet_address = ? AND currency = ?
    `, address, currency).Scan(&balance)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // Return 0 if no balance record exists
		}
		return 0, err
	}

	return balance, nil
}

// GetAllBalances gets all currency balances for a wallet
func GetAllBalances(db *sql.DB, address string) ([]Balance, error) {
	// Check if wallet exists
	_, err := GetWallet(db, address)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
        SELECT currency, balance FROM wallet_balances 
        WHERE wallet_address = ?
    `, address)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var balances []Balance
	for rows.Next() {
		var b Balance
		if err := rows.Scan(&b.Currency, &b.Amount); err != nil {
			return nil, err
		}
		balances = append(balances, b)
	}

	return balances, nil
}

// HasSufficientBalance checks if a wallet has enough of a currency
func HasSufficientBalance(db *sql.DB, address string, currency string, amount int) (bool, error) {
	balance, err := GetBalance(db, address, currency)
	if err != nil {
		return false, err
	}

	return balance >= amount, nil
}

// DeductFunds removes funds from a wallet (used during swaps)
func DeductFunds(db *sql.DB, address string, currency string, amount int) error {
	// Check if has sufficient balance
	hasBalance, err := HasSufficientBalance(db, address, currency, amount)
	if err != nil {
		return err
	}

	if !hasBalance {
		return errors.New("insufficient balance")
	}

	// Deduct funds
	_, err = db.Exec(`
        UPDATE wallet_balances 
        SET balance = balance - ? 
        WHERE wallet_address = ? AND currency = ?
    `, amount, address, currency)

	if err != nil {
		return err
	}

	log.Printf("Deducted %d of %s from wallet %s", amount, currency, address)
	return nil
}
