package main

import (
	"fmt"
	"simple-go-blockchain/blockchain"
	"time"
)

func main() {

	b1 := blockchain.NewChain()

	// Create list of blocks for testing
	for i := 0; i < 5; i++ {
		b1.AddToPool(blockchain.Transaction{
			Sender:    "Jaquelyn",
			Receiver:  "Hiro",
			Amount:    230,
			Timestamp: time.Now(),
		})
		b1.NewBlock()
	}

	// List blocks
	fmt.Printf("BLOCKS:\n")

	for b := range b1.ChainSlice {
		//fmt.Printf("%+v\n\n", b1.ChainSlice[b])
		fmt.Printf("#%d\n", b1.ChainSlice[b].BlockNumber)
		fmt.Printf("Nonce: %d\n", b1.ChainSlice[b].Nonce)
		fmt.Printf("Body: %v\n", b1.ChainSlice[b].Body)
		fmt.Printf("Prev: %x\n", b1.ChainSlice[b].PrevHashVal)
		fmt.Printf("Hash: %x\n", b1.ChainSlice[b].HashVal)
	}
}
