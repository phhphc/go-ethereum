package accounts

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

var Accounts []*keystore.Key
var password string

func init() {
	fmt.Print("Passphrase: ")
	fmt.Scan(&password)
	getAccounts()
	if len(Accounts) < 1 {
		panic("Need to create at least one account")
	}
}

func getAccounts() {
	ks := keystore.NewKeyStore("./.keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	accounts := ks.Accounts()
	Accounts = make([]*keystore.Key, len(accounts))
	for i, a := range accounts {
		// load keyjson
		bs, err := os.ReadFile(a.URL.Path)
		if err != nil {
			log.Fatal(err)
		}

		// descript key json
		key, err := keystore.DecryptKey(bs, password)
		if err != nil {
			log.Fatal(err)
		}

		// add to keys slice
		Accounts[i] = key
	}
}
