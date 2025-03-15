package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Provide functionalities related to the blockchain interaction here

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
