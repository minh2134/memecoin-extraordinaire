package blockchain

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"server/internal/smartContract"
)

// Provide functionalities related to the blockchain interaction here

type Wallet struct {
	address 	string
	privateHex 	string
}

var (
	urlRPC string 	= "http://127.0.0.1:32798"
	testAccount common.Address = common.HexToAddress(
					"0x25941dC771bB64514Fc8abBce970307Fb9d477e9")
)

func Conn() (*ethclient.Client, error) {
	client, err := ethclient.Dial(urlRPC)
	if err != nil {
		log.Println("blockchain.go: something went wrong")
	}
	return client, err
}

func WalletConnect(client *ethclient.Client, wallet Wallet) *bind.TransactOpts {
	networkID, _ := client.NetworkID(context.Background())
	privateKey, _ := crypto.HexToECDSA(wallet.privateHex)
	
	address := common.HexToAddress(wallet.address)
	nonce, _ := client.PendingNonceAt(context.Background(), address)
	gasPrices, _ := client.SuggestGasPrice(context.Background())

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrices
	
	return auth
}

func DeploySmartContract(auth *bind.TransactOpts, client *ethclient.Client) {
	address, tx, instance, err := smartContract.DeploySmartContract(auth, client)
	log.Println(address, tx, instance, err)
}
