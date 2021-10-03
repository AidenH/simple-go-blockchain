package main

import (
	"fmt"
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
	fmt.Println()
}

func TestFindHash(t *testing.T) {
	fmt.Println()
}