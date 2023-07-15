package main

import(
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common" //install packages with go get 'url'
	"github.com/ethereum/go-ethereum/ethclient"
)

func main(){
	regExpression := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Println("This is a valid address", regExpression.MatchString("0xF17d119eFFA0dCbe24D3fA346860be851150358F"))

	cnn, err := ethclient.Dial("")

	if err != nil{
		log.Fatal("Failed to connect", err)
	}

	address := common.HexToAddress("0xF17d119eFFA0dCbe24D3fA346860be851150358F")

	byteCode, err := cnn.CodeAt(context.Background(), address, nil)

	if err != nil{
		log.Fatal("Unable to check the address", address, err)
	}

	isContract := len(byteCode) > 0

	fmt.Println(address, "Is contract?", isContract)
}