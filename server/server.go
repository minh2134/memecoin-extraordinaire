package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// internal server packages
	"server/internal/blockchain"
	"server/internal/database"
	"server/internal/limit"
	"server/internal/smartContract"
	"server/internal/swap"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// global pointers to pass around in the handler
var (
	db *sql.DB
	client *ethclient.Client
	wallet blockchain.Wallet
	smAddress common.Address
	instance *smartContract.SmartContract
)

func enableCORS(w *http.ResponseWriter) {
	// signal to clients this response is allowed for Cross-Origin Resource Sharing
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	var err error
	mux := http.NewServeMux()

	// establishing DB availablity
	db, err = database.Open()

	if (err != nil) { 
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
	smAddress, instance, err = blockchain.DeploySmartContract(client)

	if err != nil {
		log.Fatal(err)
	}

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
	mux.HandleFunc("GET /account/balance", balanceHandler)
	mux.HandleFunc("GET /account/curr/{curr}", currHandler)
	
	log.Println("Handling connection at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handler (w http.ResponseWriter, r *http.Request) {
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
	ret := map[string] any { "address": wallet.Address, "balance": balance, "weiBalance": weiBalance }
	json.NewEncoder(w).Encode(ret)
}



func swapHandler (w http.ResponseWriter, r *http.Request) {

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
	swapRequest.SourceAddress = wallet.Address
	// Expecting a valid JSON
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&swapRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, tx, err := swap.Swap(db, swapRequest, instance, client)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return 
	}

	log.Println(tx.Hash())
	
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
	/*
	// test in case mysterious stuff appears
	testSwapReq := swap.SwapRequest {
		SourceCurr: 	"PEPE",
		TargetCurr: 	"BTC",
		SourceAmount: 	1499,
		Rate:		0.58823529411
		SourceAddress: 	"sdasdaw",
	}

	swap.Swap(db, testSwapReq)
	*/
}

func limitHandler (w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	var limitRequest limit.LimitRequest
	// Expecting a valid JSON
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&limitRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := limit.Limit(db, limitRequest)
	if err != nil {
		log.Println(err)
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
	ret := map[string] any {"address": wallet.Address, "balance": balance, "weiBalance": weiBalance }
	json.NewEncoder(w).Encode(ret)
}

func currHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if wallet.Address == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := blockchain.BalanceRequest {
		Address: wallet.Address,
		Curr: r.PathValue("curr"),
	}

	balance, err := blockchain.GetCurrBalance(smAddress, client, request)
	if err != nil {
		http.NotFound(w, r)
		log.Println(request.Curr, ":", err)
		return
	}

	ret := map[string] any {"currency": request.Curr, "amount": balance}
	json.NewEncoder(w).Encode(ret)
}
