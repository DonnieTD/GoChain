package blockchain

import (
	"BlockChain/block"
	"fmt"
)

type BlockChain struct {
	Blocks []*block.Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := block.CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Guessing this is how analysis is done in general on these kinds of systems ( does this work even on MOnero ? )
func (chain *BlockChain) Crawl() {
	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}
}

func Init() *BlockChain {
	return &BlockChain{
		Blocks: []*block.Block{block.Genesis()},
	}
}
