package lib

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func GetAccountBalance() {
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount("Creation password")
	fmt.Println(newAcc)

	client, _ := ethclient.Dial("https://rinkeby.infura.io/v3/bebe1ba6263e44279370df5b205e8b9c")

	account := common.HexToAddress("0xCC5d22bB804D7BFFc48BDAe7563Ba7843FE4928a")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034
}
