package accounts

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

type Account struct {
	// Code []byte *later
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
	// Count int big (count transaction) *later
}

func NewAccount(balance float64, Password string) (*Account, *KeyAccount, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	publicKey := privateKey.PublicKey

	hash := sha3.NewLegacyKeccak256()
	hash.Write(elliptic.Marshal(elliptic.P256(), publicKey.X, publicKey.Y))

	address := "AVAFu" + hex.EncodeToString(hash.Sum(nil)[12:])

	ka, err := CreateKeyAccount(address, privateKey, Password)
	if err != nil {
		return nil, nil, err
	}

	return &Account{
		Address: address,
		Balance: balance,
	}, ka, err
}
