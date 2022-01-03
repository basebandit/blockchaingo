package blockchaingo

import (
	"time"
)

//Block stores transactions and its associated metadata.
type Block struct {
	//Timestamp is the current timestamp(when the block was created).
	Timestamp int64
	//Data is the actual valuable information contained in the block.
	Data []byte
	//PrevBlockHash stores the hash of the previous block.
	PrevBlockHash []byte
	//Hash is the hash of the block.
	Hash []byte
	//Nonce arbitrary number to be used once.
	Nonce int
}

//NewBlock creates a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
