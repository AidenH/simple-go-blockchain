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

var verbose bool = false

// Init a new chain - a slice of Block structs
func NewChain() Chain {

	c := Chain{ChainSlice: []Block{}}

	return c
}

// Validate each block on the chain for signature
func (c *Chain) ValidateChain() (int, error) {
	
	for blockNum := range c.ChainSlice {
		thisHash := c.ChainSlice[blockNum].HashVal
		thisprevHash := c.ChainSlice[blockNum].PrevHashVal

		// Check first two bytes of Hash in block are signed
		if !bytes.Equal(c.ChainSlice[blockNum].HashVal[0:2], []byte{0, 0}) {

			err := fmt.Errorf("error: block's hash is not signed" +
				"%x", thisHash)
			
			return blockNum, err
		}

		// Check prevHash value is correct with prev block
		if blockNum > 0 {
			realprevHash := c.ChainSlice[blockNum-1].HashVal

			if !bytes.Equal(realprevHash[:], thisprevHash[:]) {
				err := fmt.Errorf("error: Previous hash mismatch." +
					"This: %x" +
					"Last: %x", thisprevHash, realprevHash)

				return blockNum, err
			}
		}
	}
	
	return 0, nil
}

// Create individual block
func (c *Chain) NewBlock(bodyMessage string) {

	chainLength := len(c.ChainSlice)

	b := Block{}

	// If new block is first block in chain
	if len(c.ChainSlice) == 0 {

		//firstNonce := 42
		//firstHash := sha256.Sum256([]byte(bodyMessage+strconv.Itoa(firstNonce)))
		newHash, newNonce := findHash(bodyMessage, [32]uint8{})

		// Assemble first block with duplicate hash and prev hash
		b = Block{
			BlockNumber: chainLength,
			Nonce:       newNonce,
			Body:        bodyMessage,
			HashVal:     newHash,
			PrevHashVal: newHash,
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

				if verbose {
					fmt.Printf("%v: %v\n", n, newHash[0:2])
				}
			}

			newHash = sha256.Sum256([]byte(bodyMessage + strconv.Itoa(n)))
		}

		n++
	}

	if verbose {
		fmt.Printf("Nonce found! n = %v\n%x\n\n", n, newHash)
	}

	return newHash, n
}