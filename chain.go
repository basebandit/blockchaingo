package blockchaingo

//Blockchain is just a database with an ordered, back-linked list.Blocks are stored in the insertion order and that each
//block is linked to the previous one.
type Blockchain struct {
	//keep ordered hashes (arrays are ordered in go)
	hashes [][]byte
	//hash -> block pair for retrieving a block by its hash.(maps are unordered in go)
	blocks map[string]*Block
}

//AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(data string) {
	prevBlockHash := bc.hashes[len(bc.hashes)-1]
	prevBlock := bc.blocks[string(prevBlockHash)]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.hashes = append(bc.hashes, newBlock.Hash)
	bc.blocks[string(newBlock.Hash)] = newBlock
}

//GetBlocks retrieves all blocks in the blockchain
func (bc *Blockchain) GetBlocks() []*Block {
	var blocks []*Block

	for _, block := range bc.blocks {
		blocks = append(blocks, block)
	}

	return blocks
}

//NewGenesisBlock creates the very first block in the blockchain.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//NewBlockchain creates a new blockchain with the genesis block.
func NewBlockchain() *Blockchain {
	genesisBlock := NewGenesisBlock()
	return &Blockchain{[][]byte{genesisBlock.Hash}, map[string]*Block{string(genesisBlock.Hash): genesisBlock}}
}
