package main

import (
	"fmt"
	"strconv"

	"maxkam/golang_blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for i, block := range chain.Blocks {
		fmt.Printf("Block #: %d\n", i)
		fmt.Printf("Previous Hash: %x\n", block.PreviousBlockHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Block Hash: %x\n\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
