package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Chain struct {
	ChainSlice []Block
}

type Block struct {
	BlockNumber int
	Nonce       int
	Body        string
	HashVal     [32]uint8
	PrevHashVal [32]uint8
}

// Init a new chain - a slice of Block structs
func NewChain() Chain {

	c := Chain{ChainSlice: []Block{}}

	return c
}

// Init individual block struct
func (c *Chain) NewBlock(bodyMessage string) {

	chainLength := len(c.ChainSlice)

	b := Block{}

	// If new block is first block in chain
	if len(c.ChainSlice) == 0 {

		firstNonce := 42
		firstHash := sha256.Sum256([]byte(bodyMessage+strconv.Itoa(firstNonce)))

		b = Block{
			BlockNumber: chainLength,
			Nonce:       42,
			Body:        bodyMessage,
			HashVal:     firstHash,
			PrevHashVal: firstHash,
		}
	} else {

		// Get hash from previous block
		prevHash := c.ChainSlice[chainLength-1].HashVal

		// Find new hash and nonce
		newHash, newNonce := findHash(bodyMessage, prevHash)

		// Assemble new block with new hash and nonce
		b = Block{
			BlockNumber: chainLength,
			Nonce:       newNonce,
			Body:        bodyMessage,
			HashVal:     newHash,
			PrevHashVal: prevHash,
		}
	}

	c.ChainSlice = append(c.ChainSlice, b)
}

func findHash(bodyMessage string, prevHash [32]uint8) ([32]uint8, int) {
	n := 0

	// Init newHash so we don't find a signed hash at n = 0
	newHash := sha256.Sum256([]byte(bodyMessage + strconv.Itoa(n)))

	// Find hash of input string bodyMessage and n nonce
	for !bytes.Equal(newHash[0:2], []byte{0, 0}) {

		if n > 0 {
			if n % 10000 == 0 {
				fmt.Printf("%v: %v\n", n, newHash[0:2])
			}

			newHash = sha256.Sum256([]byte(bodyMessage + strconv.Itoa(n)))
		}

		n++
	}

	fmt.Printf("Nonce found! n = %v\n%x\n\n", n, newHash)

	return newHash, n
}