package old

import(
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func createWallet(){
	privateKey, err := crypto.GenerateKey()

	if err != nil{
		log.Fatal("Unable to generate key")
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	
	walletAdd := hexutil.Encode(privateKeyBytes)

	fmt.Println(walletAdd)
}