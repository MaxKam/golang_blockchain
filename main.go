package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data              []byte
	Hash              []byte
	PreviousBlockHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousBlockHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte(data), []byte{}, prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func initBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := initBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for i, block := range chain.blocks {
		fmt.Printf("Block #: %d\n", i)
		fmt.Printf("Previous Hash: %x\n", block.PreviousBlockHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Block Hash: %x\n\n", block.Hash)
	}
}
