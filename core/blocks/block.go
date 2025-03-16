package blocks

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Block struct {
	// HEADER
	BID int `json:"block_id"`
	// Hash      []byte
	// PrevHash  []byte
	// Validator "ADDRESS"
	// AFuelLimit (Fee)
	// AFuelUsed (Fee)
	StateRoot []byte `json:"state_root"`
	Timestamp int64  `json:"timestamp"`

	// Transactions
	Transf []*Transaction
}

func CreateBlock(id int, Transactions []*Transaction) Block {
	block := Block{
		BID:       id,
		Timestamp: time.Now().Unix(),
		Transf:    Transactions,
		StateRoot: []byte{},
	}

	return block
}

func HashBlock(block Block) string {
	blockBytes, _ := json.Marshal(block)
	hash := sha256.Sum256(blockBytes)
	return string(hash[:])
}
