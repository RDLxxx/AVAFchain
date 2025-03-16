package merkle

import (
	"fmt"

	"github.com/RDLxxx/AVAFchain/accounts"
	borsh "github.com/RDLxxx/AVAFchain/storage/Borsh"
	"github.com/RDLxxx/AVAFchain/storage/merkle"
)

func createAccountNode(account *accounts.Account) (*merkle.ShortNode, error) {
	data, err := borsh.DeserializeAccount(*account)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize account: %w", err)
	}

	key := []byte(account.Address)
	shortNode := &merkle.ShortNode{
		Key:   key,
		Val:   merkle.ValueNode(data),
		flags: merkle.NodeFlag{},
	}

	return shortNode, nil
}
