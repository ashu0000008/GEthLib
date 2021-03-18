package lib

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountBalance() {
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount("Creation password")
	fmt.Println(newAcc)

	ethclient.Dial()
}
