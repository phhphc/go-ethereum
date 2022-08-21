package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const etheURL = "http://localhost:8545"

var password string
var accounts []*keystore.Key

func init() {
	// parse command line argument
	if argc := len(os.Args); argc < 2 {
		func() {
			fmt.Printf("Usage: %s <pass_phrase>\n", os.Args[0])
		}()
		os.Exit(1)
	}
	password = os.Args[1]

	// parse key list
	accounts = getAccounts()
}

func main() {
	if len(accounts) < 2 {
		fmt.Println("Please create at least 2 wallet")
		os.Exit(1)
	}

	cl, err := ethclient.DialContext(context.Background(), etheURL)
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := cl.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signer := types.NewEIP155Signer(chainID)

	// transfer amount wei from wallet 1 to wallet 2
	account1 := accounts[0]
	account2 := accounts[1]
	amount := big.NewInt(1e18)
	ctx := context.Background()

	nonce, err := cl.NonceAt(ctx, account1.Address, nil)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := cl.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// create new transaction
	tx := types.NewTransaction(nonce, account2.Address, amount, 21000, gasPrice, nil)
	tx, err = types.SignTx(tx, signer, account1.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// apply transaction
	err = cl.SendTransaction(ctx, tx)
	if err != nil {
		log.Fatal(err)
	}

	// print transaction result
	fmt.Printf("tx sent: %v", tx.Hash().Hex())
}

func getAccounts() []*keystore.Key {
	ks := keystore.NewKeyStore("./.keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	accounts := ks.Accounts()

	var keys []*keystore.Key
	for _, a := range accounts {
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
		keys = append(keys, key)
	}

	return keys
}
