// @author Dinh Le Hoang Anh - 105508318
// @author Pham Vu Minh - 105110564
package main

import (
	// internal server packages
	"os"
	"server/internal/blockchain"
	"server/internal/database"
	"server/internal/limit"
	"server/internal/priceFeed"
	"server/internal/smartContract"
	"server/internal/swap"

	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// global pointers to pass around in the handler
var (
	db        *sql.DB
	client    *ethclient.Client
	wallet    blockchain.Wallet
	smAddress common.Address
	instance  *smartContract.SmartContract
	btcPrice  *priceFeed.PriceFeed
	daiPrice  *priceFeed.PriceFeed
	snxPrice  *priceFeed.PriceFeed
)

func enableCORS(w *http.ResponseWriter) {
	// signal to clients this response is allowed for Cross-Origin Resource Sharing
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	var err error
	mux := http.NewServeMux()

	// loading env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// establishing DB availablity
	db, err = database.Open(os.Getenv("ME_DATABASE_PATH"))

	if err != nil {
		log.Fatal("Opening database failed.")
	}
	log.Println("Opened the database successfully.")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Establishing a connection failed.")
	}
	log.Println("Connected to the database successfully.")

	err = database.Bootstrap(db)

	// connecting to a node in local blockchain testnet
	client, err = blockchain.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	
	
	// Deploy the transaction on chain
	log.Println("Deploying the contract")
	smAddress, instance, err = blockchain.DeploySmartContract(os.Getenv("ME_SMACON_ADDRESS"), client)
	if err != nil {
		log.Fatal(err)
	} 
	log.Println("Consider to add this line to the .env file to skip the deployment on next run:")
	log.Println("ME_SMACON_ADDRESS=", smAddress.Hex())

	// Deploy the price feed contracts on chain
	priceFeedContracts := []string {
		"0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43", // BTC/USD
		"0x14866185B1962B63C3Ea9E03Bc1da838bab34C19", // DAI/USD
		"0xc0F82A46033b8BdBA4Bb0B0e28Bc2006F64355bC", // SNX/USD
	}
	priceFeeds, err := blockchain.LoadPriceFeeds(priceFeedContracts, client)
	
	// assign the results to global instances
	btcPrice = priceFeeds[0]
	daiPrice = priceFeeds[1]
	snxPrice = priceFeeds[2]
	
	// Mint some tokens for trading
	log.Println("Minting tokens to accounts for trading...")
	err = blockchain.Mint(instance, client)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Minting done!")

	// Router
	mux.HandleFunc("/", handler)
	mux.HandleFunc("GET /auth", authHandler)
	mux.HandleFunc("POST /trade/swap", swapHandler)
	mux.HandleFunc("POST /trade/limit", limitHandler)
	mux.HandleFunc("GET /trade/swap/{sourceCurr}/{targetCurr}", rateSuggestorHandler)
	mux.HandleFunc("GET /account/balance", balanceHandler)
	mux.HandleFunc("GET /account/curr/{curr}", currHandler)

	log.Println("Handling connection at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth" {
		http.NotFound(w, r)
		return
	}
	enableCORS(&w)

	wallet = blockchain.ConnectWallet()
	balance, weiBalance, _ := blockchain.GetBalance(client, wallet)
	ret := map[string]any{"address": wallet.Address, "balance": balance, "weiBalance": weiBalance}
	json.NewEncoder(w).Encode(ret)
}

func swapHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/trade/swap" {
		http.NotFound(w, r)
		return
	}

	enableCORS(&w)
	if wallet.Address == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var swapRequest swap.SwapRequest
	// Expecting a valid JSON
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&swapRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	swapRequest.SourceAddress = wallet.Address // source address always is wallet address

	result, tx, err := swap.Swap(db, swapRequest, instance, client)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Println("swapTxHash:", tx.Hash())

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
}

func limitHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if wallet.Address == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var limitRequest limit.LimitRequest
	// Expecting a valid JSON
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&limitRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("not good")
		return
	}

	limitRequest.SourceAddress = wallet.Address // source address always is wallet address

	result, tx, err := limit.Limit(db, limitRequest, instance, client)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	if result.IsMatched {
		log.Println("limitTxHash:", tx.Hash())
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.URL.Path != "/account/balance" {
		http.NotFound(w, r)
		return
	}

	if wallet.Address == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	balance, weiBalance, _ := blockchain.GetBalance(client, wallet)
	ret := map[string]any{"address": wallet.Address, "balance": balance, "weiBalance": weiBalance}
	json.NewEncoder(w).Encode(ret)
}

func currHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if wallet.Address == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := blockchain.BalanceRequest{
		Address: wallet.Address,
		Curr:    r.PathValue("curr"),
	}

	balance, err := blockchain.GetCurrBalance(smAddress, client, request)
	if err != nil {
		http.NotFound(w, r)
		log.Println(request.Curr, ":", err)
		return
	}

	ret := map[string]any{"currency": request.Curr, "amount": balance}
	json.NewEncoder(w).Encode(ret)
}

// helper function to find pricefeed
func getPriceFeed(curr string) (*priceFeed.PriceFeed, error) {
	if curr == "BTC" {
		return btcPrice, nil
	}
	
	if curr == "DAI" {
		return daiPrice, nil
	}

	if curr == "SNX" {
		return snxPrice, nil
	}

	return nil, errors.New("not found") // nothing else
}

func rateSuggestorHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	
	sourceCurr, err := getPriceFeed(r.PathValue("sourceCurr"))
	if err != nil {
		http.NotFound(w, r)
		return 
	}

	targetCurr, err := getPriceFeed(r.PathValue("targetCurr"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rate, err := blockchain.GetRate(sourceCurr, targetCurr, client)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	rateReturn := map[string]any{"rate": rate}
	json.NewEncoder(w).Encode(rateReturn)
}
