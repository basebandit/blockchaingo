package main

import (
	"blockchaingo"
	"fmt"
	"strconv"
)

func main() {
	bc := blockchaingo.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Parish")
	bc.AddBlock("Send 2 more BTC to Parish")

	for _, block := range bc.GetBlocks() {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchaingo.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
