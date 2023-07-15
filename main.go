package main

import(
	// "context"
	"fmt"
	"log"

	// "github.com/ethereum/go-ethereum/common" //install packages with go get 'url'
	// "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main(){
	privateKey, err := crypto.GenerateKey()

	if err != nil{
		log.Fatal("Unable to generate key")
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	
	walletAdd := hexutil.Encode(privateKeyBytes)

	fmt.Println(walletAdd)
}