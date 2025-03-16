package blocks

import (
	"crypto/ecdsa"
	crp "crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/RDLxxx/AVAFchain/accounts"
	utilsa "github.com/RDLxxx/AVAFchain/utils/accounts"
)

type Transaction struct {
	TID int `json:"transaction_id"`
	// AfuelUsed
	// AfuelPrice
	From      string  `json:"from"`
	To        string  `json:"to"`
	Value     float64 `json:"value"`
	Data      []byte  `json:"data"`
	IsValid   bool    `json:"valid"`
	Signature string  `json:"signature"`
}

func NewFLTransaction(ID int, Frm string, Tt string, Val float64) Transaction {
	Trc := Transaction{
		TID:       ID,
		From:      Frm,
		To:        Tt,
		Value:     Val,
		IsValid:   false,
		Signature: "None",
	}

	return Trc
}
func Sign(trc Transaction, account accounts.Account, privateKey crp.PrivateKey) Transaction {
	hash := computeTransactionHash(trc)
	if utilsa.IsGoodPrv(account, privateKey) {
		if account.Address == trc.From {
			r, s, _ := ecdsa.Sign(rand.Reader, &privateKey, hash[:])
			signature := append(r.Bytes(), s.Bytes()...)
			hcc := hex.EncodeToString(signature)
			trc := Transaction{
				IsValid:   utilsa.IsGoodPrv(account, privateKey),
				Signature: hcc,
			}

			return trc
		}
	}
	return trc
}

func computeTransactionHash(trc Transaction) [32]byte {
	data := []byte{}
	data = append(data, []byte(string(trc.TID))...)
	data = append(data, []byte(trc.From)...)
	data = append(data, []byte(trc.To)...)
	data = append(data, []byte(fmt.Sprintf("%f", trc.Value))...)
	data = append(data, trc.Data...)
	hash := sha256.Sum256(data)
	return hash
}
