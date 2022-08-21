package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Ether = new(big.Float).SetFloat64(math.Pow10(18))
var ethCliUrl = "http://127.0.0.1:8545"

func main() {
	// connect ethereum client
	cl, err := ethclient.DialContext(context.Background(), ethCliUrl)
	if err != nil {
		log.Fatal("Fail to dial ethereum client:", err)
	}
	defer cl.Close()

	// get latest block
	block, err := cl.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("Fail to get latest block:", err)
	}
	fmt.Println(block.Number())

	// get account balance
	addr := common.HexToAddress("0x842d6599d0e173a255c6bad0b3e2a6ca0c72782b")
	wei, err := cl.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		panic(err)
	}
	eth := new(big.Float).Quo(new(big.Float).SetInt(wei), Ether)

	fmt.Printf("Balance of account %s is %s ether\n", addr, eth.String())
}
