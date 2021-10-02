package main

import (
	"fmt"
	"strconv"
	"very-simple-go-blockchain/blockchain"
)

func main() {

	b1 := blockchain.NewChain()

	// Create list of blocks for testing
	for i := 0; i < 5; i++ {
		b1.NewBlock("body" + strconv.Itoa(i))
	}

	// List blocks
	fmt.Printf("BLOCKS:\n")

	for b := range b1.ChainSlice {
		fmt.Printf("%+v\n\n", b1.ChainSlice[b])
	}
}