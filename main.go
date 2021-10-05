package main

import (
	"fmt"
	"simple-go-blockchain/blockchain"
	"strconv"
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
		//fmt.Printf("%+v\n\n", b1.ChainSlice[b])
		fmt.Printf("#%d\n", b1.ChainSlice[b].BlockNumber)
		fmt.Printf("Nonce: %d\n", b1.ChainSlice[b].Nonce)
		fmt.Printf("Body: %s\n", b1.ChainSlice[b].Body)
		fmt.Printf("Prev: %x\n", b1.ChainSlice[b].PrevHashVal)
		fmt.Printf("Hash: %x\n", b1.ChainSlice[b].HashVal)
	}
}