package main

import (
	"bytes"
	"fmt"
	"simple-go-blockchain/blockchain"
	"testing"
	"time"
)

func TestNewChain(t *testing.T) {
	testChain := blockchain.NewChain()
	varType := fmt.Sprintf("%T", testChain)

	// Check whether testChain instance is Chain struct
	if varType != "blockchain.Chain" {
		t.Errorf("error initializing blockchain.Chain"+
			"Result = %T", varType)
	}
}

func TestNewBlock(t *testing.T) {
	testChain := blockchain.NewChain()
	testChain.AddToPool(blockchain.Transaction{
		Sender:   "Jocelyn",
		Receiver: "Hal",
		Amount:   230,
	})
	testChain.NewBlock()

	newNonce := testChain.ChainSlice[0].Nonce
	newHash := bytes.Equal(testChain.ChainSlice[0].HashVal[0:2], []byte{0, 0})

	// Test for ChainSlice for no blocks
	if len(testChain.ChainSlice) == 0 {
		t.Errorf("error creating new block"+
			"\nChain: %v", testChain.ChainSlice)
	}

	// Test that private findHash() is producing nonce and hash
	if newNonce != 7099 || !newHash {
		t.Errorf("nonce or HashVal mismatched/absent in block"+
			"\nNonce: %v, Hash: %x", newNonce, testChain.ChainSlice[0].HashVal)
	}
}

func TestValidateChain(t *testing.T) {
	testChain := blockchain.NewChain()

	// Create ten test blocks on chain
	for i := 0; i < 10; i++ {
		testChain.AddToPool(blockchain.Transaction{
			Sender:    "Richard",
			Receiver:  "Linda",
			Amount:    12,
			Timestamp: time.Now(),
		})
		testChain.NewBlock()
	}

	blockNum, err := testChain.ValidateChain()
	if err != nil {
		t.Errorf("%v, Block: %v", err, blockNum)
	}
}

func TestAddToPool(t *testing.T) {
	testChain := blockchain.NewChain()

	transact := blockchain.Transaction{
		Sender:    "Ralph",
		Receiver:  "Sarah",
		Amount:    32,
		Timestamp: time.Now(),
	}

	err := testChain.AddToPool(transact)
	if err != nil {
		t.Errorf("%v", err)
	}

	// Read testChain.Pool afterwards and check the struct passe correctly
	if testChain.Pool[0] != transact {
		t.Errorf("error appending to pool - Transaction mismatch")
	}
}
