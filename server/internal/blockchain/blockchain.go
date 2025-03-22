// @author Dinh Le Hoang Anh - 105508318
// @author Pham Vu Minh - 105110564
package blockchain

import (
	"context"
	"errors"
	"log"
	"math"
	"math/big"
	"math/rand"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"server/internal/priceFeed"
	"server/internal/smartContract"
)

// Provide functionalities related to the blockchain interaction here

type Wallet struct {
	Address    string
	PrivateHex string
}

type SMTradeContract struct {
	SourceAddress string
	TargetAddress string
	SourceCurr    string
	TargetCurr    string
	SourceAmount  decimal.Decimal
	TargetAmount  decimal.Decimal
}

type BalanceRequest struct {
	Address string
	Curr    string
}

var (
	urlRPC      string         = "https://sepolia.infura.io/v3/5039db0cac5f44b3bd15771581f51116"
	testAccount common.Address = common.HexToAddress(
		"0x25941dC771bB64514Fc8abBce970307Fb9d477e9")
	decimals decimal.Decimal = decimal.NewFromInt(int64(math.Pow10(8)))
	currency = [3]string{
			"BTC",
			"DAI",
			"SNX",
		}
)

// pre-funded accounts for testing
var preFundedAccounts = [3]Wallet{
	{
		Address:    "0x8019F8FA7b20c68F7CF4Bb4d74F13AB54d7f57f3",
		PrivateHex: "a90be14aff65f7a8bb5ef0c3d41f51133ad7dd565eab7f606e09bd96483164d9",
	},
	{
		Address:    "0x9C0802a055D25040C233cCF7aEd8Fe9339c1b8c6",
		PrivateHex: "7231bdd5f1389651a38beeb8ee80679693221ff66ed55bcad47bb5971f4fb1cb",
	},
	{
		Address:    "0xb0d4A8965D85BE95828DEE79B5EC8fe51C9da53e",
		PrivateHex: "712ea5756fb6f2a5477df07bcbe54922b3648e626ff040d3b1079cd04372ba9a",
	},
}

var godWallet Wallet = Wallet {
	Address: "0xfD143FbAe1682c2b2376269B4EDff67b5C0bc9eE",
	PrivateHex: "2e8bd9a6e05accd606c2523795ec798bb5473a37b78458a61afd6bdcdaeab9b7",
}

func Conn() (*ethclient.Client, error) {
	client, err := ethclient.Dial(urlRPC)
	if err != nil {
		log.Println("blockchain.go: something went wrong")
	}
	return client, err
}

func newTransactor(client *ethclient.Client, wallet Wallet) (*bind.TransactOpts, error) {
	var (
		auth *bind.TransactOpts
		err  error
	)

	chainID, _ := client.ChainID(context.Background())
	privateKey, _ := crypto.HexToECDSA(wallet.PrivateHex)

	address := common.HexToAddress(wallet.Address)
	nonce, _ := client.PendingNonceAt(context.Background(), address)
	gasPrices, err := client.SuggestGasPrice(context.Background())

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return auth, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrices

	return auth, err
}

func getWallet(address string) (Wallet, error) {
	var (
		wallet  Wallet
		isFound bool = false
	)

	for _, account := range preFundedAccounts {
		if account.Address == address || account.Address == address[:2] {
			wallet = account
			isFound = true
		}
	}

	if !isFound {
		return wallet, errors.New("no wallet found at this given address")
	}

	return wallet, nil
}

func ConnectWallet() Wallet {
	return preFundedAccounts[rand.Intn(len(preFundedAccounts))]
}

func GetBalance(client *ethclient.Client, wallet Wallet) (float64, *big.Int, error) {
	address := common.HexToAddress(wallet.Address)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	ethBalance, _ := balance.Float64()
	ethBalance = ethBalance / math.Pow(10, 18)
	return ethBalance, balance, err
}

func DeploySmartContract(client *ethclient.Client) (common.Address,
	*smartContract.SmartContract,
	error) {
	var (
		address  common.Address
		instance *smartContract.SmartContract
		tx       *types.Transaction
	)
	auth, err := newTransactor(client, godWallet)
	if err != nil {
		return address, instance, err
	}
	address, tx, instance, err = smartContract.DeploySmartContract(auth, client)
	// Wait for contract to be mined to continue
	log.Println("Waiting for the contract to be mined...")
	bind.WaitMined(context.Background(), client, tx)
	bigIntDecimals, _ := instance.Decimals(nil)
	if err == nil {
		// if no error occurs update the fixed-point decimal
		decimals = decimal.NewFromBigInt(bigIntDecimals, 0)
		log.Println("New decimal point:", decimals)
	}
	log.Print("Done!")

	return address, instance, err
}

func GetCurrBalance(smAddress common.Address,
	client *ethclient.Client,
	request BalanceRequest) (decimal.Decimal, error) {

	instance, err := smartContract.NewSmartContract(smAddress, client)
	if err != nil {
		log.Println("failed at getting an instance")
		return decimal.NewFromInt(1), err
	}

	address := common.HexToAddress(request.Address)
	balance, err := instance.GetBalance(nil, address, request.Curr)
	if err != nil {
		return decimal.NewFromInt(3), err
	}

	decimalBalance := decimal.NewFromBigInt(balance, -int32(decimals.IntPart()))

	return decimalBalance, err
}

func Mint(instance *smartContract.SmartContract, client *ethclient.Client) error {
	// minting some tokens for accounts to trade
	var (
		
		err error
		wg  sync.WaitGroup
	)
	amount := decimal.NewFromInt(1000)
	mult := decimal.NewFromInt(10).Pow(decimals) // multiplier to convert into raw int form with preferred precision
	smAmount := amount.Mul(mult).BigInt()

	// we want to loop all of our minting transactions quickly
	// but we also want to wait for all transactions to be mined before going any
	// further
	// wg is here as a counter for pending transaction, adding one when a new
	// transaction is created
	// each WaitMined + wg counter decrease is wrapped in a goroutine, allowing us
	// to do this asynchronously
	for _, wallet := range preFundedAccounts {
		for _, curr := range currency {
			wg.Add(1)
			auth, err := newTransactor(client, godWallet)
			if err != nil {
				return err
			}
			address := common.HexToAddress(wallet.Address)
			tx, err := instance.MintTokens(auth, address, curr, smAmount)
			if err != nil {
				log.Println(err)
			}
			go func(client *ethclient.Client, tx *types.Transaction) {
				defer wg.Done()
				bind.WaitMined(context.Background(), client, tx)
				log.Println("Minted", amount, curr, "at", address)
			}(client, tx)
		}
	}

	wg.Wait()
	return err

}

func ExecuteTrade(instance *smartContract.SmartContract, client *ethclient.Client, request SMTradeContract) (*types.Transaction, error) {
	var (
		tx  *types.Transaction
		err error
	)

	sourceWallet, _ := getWallet(request.SourceAddress)
	targetWallet, _ := getWallet(request.TargetAddress)

	auth, err := newTransactor(client, sourceWallet)
	if err != nil {
		return tx, err
	}

	mult := decimal.NewFromInt(10).Pow(decimals) // multiplier to convert into raw int form with preferred precision
	sourceAmount := request.SourceAmount.Mul(mult).BigInt()
	targetAmount := request.TargetAmount.Mul(mult).BigInt()

	//sourceAddress := common.HexToAddress(request.SourceWallet.Address)
	targetAddress := common.HexToAddress(targetWallet.Address)
	tx, err = instance.ExecuteTrade(auth, targetAddress, request.SourceCurr, request.TargetCurr, sourceAmount, targetAmount)

	return tx, err
}

// =========== Price feed section ==============
func LoadPriceFeeds(addresses []string, 
			client *ethclient.Client) ([]*priceFeed.PriceFeed, error) {
	
	// find price feed contract addresses on Chainlink website
	var (
		instances []*priceFeed.PriceFeed
		err error
	)
	
	for _, add := range addresses {
		address := common.HexToAddress(add)
		instance, err := priceFeed.NewPriceFeed(address, client)
		if err != nil {
			return instances, err
		}

		instances = append(instances, instance)
	}

	return instances, err
}

func GetRate(curr1 *priceFeed.PriceFeed, 
		curr2 *priceFeed.PriceFeed, 
		client *ethclient.Client) (decimal.Decimal, error){

		var rate decimal.Decimal
		// get the raw data
		curr1Answer, err := curr1.LatestRoundData(nil)
		if err != nil {
			log.Println("blockchain.go: can't get curr1 data")
			return rate, err
		}

		curr2Answer, err := curr2.LatestRoundData(nil)
		if err != nil {
			log.Println("blockchain.go: can't get curr2 data")
			return rate, err
		}

		// get the decimals
		curr1Decimals, err := curr1.Decimals(nil)
		if err != nil {
			log.Println("blockchain.go: can't get curr1 decimals")
			return rate, err
		}

		curr2Decimals, err := curr2.Decimals(nil)
		if err != nil {
			log.Println("blockchain.go: can't get curr2 decimals")
			return rate, err
		}

		// convert the answers to decimals
		curr1Rate := decimal.NewFromBigInt(curr1Answer.Answer, int32(curr1Decimals))
		curr2Rate := decimal.NewFromBigInt(curr2Answer.Answer, int32(curr2Decimals))

		return curr1Rate.Div(curr2Rate), err
}
