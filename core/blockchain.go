package core

import "github.com/RDLxxx/AVAFchain/core/blocks"

// blockchain type ->
type BlockChain struct {
	CBHash []byte
	// db database
	GenesisBlock *blocks.Block
}

func NewBlockchain() BlockChain {

	// For Tests -?
	GenesisTRC := blocks.NewFLTransaction(0, "AVAFuXXXXXXXXXXXXXXXXXXXXXXX", "AVAFuXXXXXXXXXXXXXXXXXXXXXXX", 1488)
	transactions := []*blocks.Transaction{
		&GenesisTRC,
	}
	genesisblock := blocks.CreateBlock(0, transactions)

	bc := BlockChain{
		GenesisBlock: &genesisblock,
	}

	return bc
}
