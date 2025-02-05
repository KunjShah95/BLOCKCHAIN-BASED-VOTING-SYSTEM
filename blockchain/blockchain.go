package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Index     int    `json:"index"`
	Vote      string `json:"vote"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prev_hash"`
}

type Blockchain struct {
	Blocks []Block
	mu     sync.Mutex
}

func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []Block{}}
}

func (bc *Blockchain) AddVote(vote string) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	var prevHash string
	if len(bc.Blocks) > 0 {
		prevHash = bc.Blocks[len(bc.Blocks)-1].Hash
	}

	block := Block{
		Index:    len(bc.Blocks) + 1,
		Vote:     vote,
		Hash:     calculateHash(len(bc.Blocks)+1, vote, prevHash),
		PrevHash: prevHash,
	}

	bc.Blocks = append(bc.Blocks, block)
}

func calculateHash(index int, vote, prevHash string) string {
	record := fmt.Sprintf("%d%s%s", index, vote, prevHash)
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}