package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/phhphc/go-ethereum/accounts"
	todo "github.com/phhphc/go-ethereum/gen"
)

var cAddr = common.HexToAddress("0x9D2C1BDE2F22fe07224eb3Fce429fdef3332A4de")

func main() {
	cl, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer cl.Close()

	t, err := todo.NewTodo(cAddr, cl)
	if err != nil {
		log.Fatal(err)
	}

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

	// add new task to contract
	tx, err := t.Add(auth, "hi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction hash:", tx.Hash().Hex())

	// list all tasks in contract
	tasks, err := t.List(&bind.CallOpts{From: account.Address})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tasks: %+v\n", tasks)
}
