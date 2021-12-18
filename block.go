package blockchaingo

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//NewBlock creates a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()

	return block
}
