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

func findHash(bodyMessage string, prevHash [32]uint8) ([32]uint8, int) {
	n := 0
	newHash := sha256.Sum256([]byte(bodyMessage + strconv.Itoa(n)))

	for !bytes.Equal(newHash[0:2], []byte{0, 0}) {

		if n > 0 {
			if n % 10000 == 0 {
				fmt.Printf("%v: %v\n", n, newHash[0:2])
			}

			newHash = sha256.Sum256([]byte(bodyMessage + strconv.Itoa(n)))
		}

		n++
	}

	fmt.Printf("Nonce found! n = %v\n%v\n\n", n, newHash)

	return newHash, n
}