package main

import(
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient" //install packages with go get 'url'
	
)

func main(){
	cnn, err := ethclient.Dial("")

	if err != nil{
		log.Fatal("Failed to connect", err)
	}

	header, err := cnn.HeaderByNumber(context.Background(), nil)

	if err != nil{
		log.Fatal("Unable to get header", err)
	}

	fmt.Println(header.Number.String())

	blockNumber := big.NewInt(5231314)
	block, err := cnn.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal("Unable to get the block by number", blockNumber)
	}

	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Time())
	fmt.Println(block.Difficulty())
	fmt.Println(block.Hash().Hex())
	fmt.Println(len(block.Transactions()))

	for _, tx := range block.Transactions() {
		if tx.Value().String() != "0" {
			fmt.Println(tx.Hash().Hex())
			fmt.Println(tx.Value().String())
			fmt.Println(tx.Nonce())
			fmt.Println(tx.To().Hex())
		}
	}
}