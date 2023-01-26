package main

import "BlockChain/blockchain"

func main() {
	chain := blockchain.Init()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")
	chain.AddBlock("Fourth block after Genesis")
	chain.AddBlock("Fith block after Genesis")

	chain.Crawl()
}
