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

	cntx := context.Background()

	txn, pending, _ := cnn.TransactionByHash(cntx, common.HexToHash("0xab313b0172492446df5704fa97d008478a31e6088d970bb81ae75bee5fc20683"))

	if pending {
		fmt.Println("Txn is pending", txn)
	} else{
		fmt.Println("Txn is not pending", txn)
	}
}