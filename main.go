package main

import (
	"fmt"
	"strconv"
	"very-simple-go-blockchain/blockchain"
)

func main() {

	b1 := blockchain.NewChain()

	for i := 0; i < 5; i++ {
		b1.NewBlock(strconv.Itoa(i))
	}

	fmt.Println(b1)
}