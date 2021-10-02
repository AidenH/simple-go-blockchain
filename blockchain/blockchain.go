package blockchain

type Chain struct {
	ChainSlice []Block
}

type Block struct {
	BlockNumber int
	Nonce       int
	Body        string
	HashVal     string
	PrevHashVal string
}

// Init a new chain - a slice of Block structs
func NewChain() Chain {

	c := Chain{ChainSlice: []Block{}}

	return c
}

// Init individual block struct
func (c *Chain) NewBlock(bodyMessage string) {

	b := Block{
		BlockNumber: len(c.ChainSlice),
		Nonce:       42,
		Body:        bodyMessage,
		HashVal:     "6a",
		PrevHashVal: "5b",
	}

	c.ChainSlice = append(c.ChainSlice, b)
}