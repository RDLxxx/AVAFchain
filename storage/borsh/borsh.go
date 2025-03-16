package borsh

import (
	"github.com/RDLxxx/AVAFchain/accounts"
	"github.com/near/borsh-go"
)

// Account
func Serialize(a accounts.Account) ([]byte, error) {
	return borsh.Serialize(a)
}

// Account
func DeserializeAccount(data []byte) (*accounts.Account, error) {
	var account accounts.Account
	err := borsh.Deserialize(&account, data)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
