package main

import (
	"fmt"
	"very-simple-go-blockchain/blockchain"
)

func main() {

	b1 := blockchain.NewChain()

	b1.NewBlock("body")

	fmt.Println(b1)
}