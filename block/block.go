package block

import (
	"bytes"
	"crypto/sha256"
)

// SEEMS LIKE A BLOCK CHAIN IS A LINKED LIST WITH HASHES INSTEAD OF POINTERS?
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func (b *Block) DeriveHash() {
	info := bytes.Join(
		[][]byte{
			b.Data,
			b.PrevHash,
		},
		[]byte{},
	)

	hash := sha256.Sum256(info)

	b.Hash = hash[:]
	//If this ^ doesn't make sense, you can look up slice defaults
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash,
		0,
	}

	block.DeriveHash()
	return block
}

// we need to create the OG block for the chain.
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
