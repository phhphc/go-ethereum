package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

var password string

func init() {
	// parse command line argument
	if argc := len(os.Args); argc < 2 {
		func() {
			fmt.Printf("Create new wallet\n\tUsage: %s <pass_phrase>\n", os.Args[0])
		}()
		os.Exit(1)
	}
	password = os.Args[1]
}

func main() {
	// create new account
	ks := keystore.NewKeyStore("./.keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(acc.Address)
}

// func main() {
// 	pvk, err := crypto.GenerateKey()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// export private key
// 	pvData := crypto.FromECDSA(pvk)
// 	fmt.Println(hexutil.Encode(pvData))

// 	// export public key
// 	puData := crypto.FromECDSAPub(&pvk.PublicKey)
// 	fmt.Println(hexutil.Encode(puData))

// 	// get address
// 	addr := crypto.PubkeyToAddress(pvk.PublicKey)
// 	fmt.Println(addr)
// }
