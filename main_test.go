package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"simple-go-blockchain/blockchain"
	"testing"
)

func TestNewChain(t *testing.T) {
	testChain := blockchain.NewChain()
	varType := fmt.Sprintf("%T", testChain)

	// Check whether testChain instance is Chain struct
	if varType != "blockchain.Chain" {
		t.Errorf("Error initializing blockchain.Chain. Result = %T", varType)
	}
}

func TestNewBlock(t *testing.T) {
	testChain := blockchain.NewChain()
	testChain.NewBlock("TestNewBlock")

	newNonce := testChain.ChainSlice[0].Nonce
	newHash := bytes.Equal(testChain.ChainSlice[0].HashVal[0:2], []byte{0, 0})

	// Test for ChainSlice for no blocks
	if len(testChain.ChainSlice) == 0 {
		t.Errorf("Error creating new block." +
			"\nChain: %v", testChain.ChainSlice)
	}

	// Test that private findHash() is producing nonce and hash
	if (newNonce != 122361 || !newHash) {
		t.Errorf("Nonce or HashVal mismatched/absent in block." +
			"\nNonce: %v, Hash: %x", newNonce, testChain.ChainSlice[0].HashVal)
	}
}

func TestValidateChain(t *testing.T) {
	testChain := blockchain.NewChain()
	
	// Create ten test blocks on chain
	for i := 0; i < 10; i++ {
		testChain.NewBlock(fmt.Sprintf("%v", rand.Intn(99999)))
	}
	
	blockNum, err := testChain.ValidateChain()
	if err != nil {
		t.Errorf("%v, Block: %v", err, blockNum)
	}
}