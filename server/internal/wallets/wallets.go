package wallets

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)




func GenerateWallet() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("private key generation failed")
	}
	
	privateKeyBytes := crypto.FromECDSA(privateKey)
	log.Println("hex private key:", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	log.Println("hex public key:", hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	log.Println("wallet address:", address)
}

func main() {
	log.Println("Generating 3 wallets...")

	for i:=1; i<=3; i++ {
		GenerateWallet()
	}
}
