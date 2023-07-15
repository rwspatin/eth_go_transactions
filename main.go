package main

import(
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common" //install packages with go get 'url'
	"github.com/ethereum/go-ethereum/ethclient"
)

func main(){
	cnn, err := ethclient.Dial("")
	
	if err != nil{
		log.Fatal("Failed to connect", err)
	}

	account := common.HexToAddress("")
	
	balance, err := cnn.BalanceAt(context.Background(), account, nil)

	if err != nil{
		log.Fatal("unable to get balances")
	} else {
		fmt.Println(balance)
	}
}