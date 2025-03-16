package accounts

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"

	"github.com/RDLxxx/AVAFchain/safety/accounts"
	"golang.org/x/crypto/scrypt"
)

type KeyAccount struct {
	Address      string `json:"address"`
	Cipher       string `json:"cipher"`
	CipherText   string `json:"ciphertext"`
	CipherParams struct {
		IV string `json:"iv"`
	} `json:"cipherparams"`
	KDF       string `json:"kdf"`
	KDFParams struct {
		Salt  string `json:"salt"`
		DKLen int    `json:"dklen"`
		N     int    `json:"n"`
		R     int    `json:"r"`
		P     int    `json:"p"`
	} `json:"kdfparams"`
	MAC string `json:"mac"`
}

func CreateKeyAccount(address string, privateKey *ecdsa.PrivateKey, password string) (*KeyAccount, error) {
	salt, err := accounts.GenerateSalt() // salt
	if err != nil {
		return nil, err
	}

	key, err := scrypt.Key([]byte(password), salt, 262144, 8, 1, 32) // scrypt
	if err != nil {
		return nil, err
	}

	iv, err := accounts.GenerateIV() // IV
	if err != nil {
		return nil, err
	}

	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey) // Change *later
	if err != nil {
		return nil, err
	}

	ciphertext, err := accounts.EncryptPrivateKey(privateKeyBytes, key[:16], iv) // Open to Decrypt
	if err != nil {
		return nil, err
	}

	mac := accounts.CreateMAC(key[16:32], ciphertext)

	KeyAccount := &KeyAccount{
		Address:    address,
		Cipher:     "aes-128-ctr",
		CipherText: hex.EncodeToString(ciphertext),
		CipherParams: struct {
			IV string `json:"iv"`
		}{
			IV: hex.EncodeToString(iv),
		},
		KDF: "scrypt",
		KDFParams: struct {
			Salt  string `json:"salt"`
			DKLen int    `json:"dklen"`
			N     int    `json:"n"`
			R     int    `json:"r"`
			P     int    `json:"p"`
		}{
			Salt:  hex.EncodeToString(salt),
			DKLen: 32,
			N:     262144,
			R:     8,
			P:     1,
		},
		MAC: hex.EncodeToString(mac),
	}

	return KeyAccount, nil
}
