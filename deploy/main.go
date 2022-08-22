package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/phhphc/go-ethereum/accounts"
	todo "github.com/phhphc/go-ethereum/gen"
)

func main() {
	cl, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer cl.Close()

	account := accounts.Accounts[0]
	ctx := context.Background()

	gasPrice, err := cl.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	chanID, err := cl.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, chanID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = 3000000

	addr, tx, _, err := todo.DeployTodo(auth, cl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract address:", addr.Hex())
	fmt.Println("Contract transaction hash:", tx.Hash().Hex())
}
