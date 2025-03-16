package accounts

import (
	"crypto/aes"
	"crypto/cipher"
	crp "crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/RDLxxx/AVAFchain/accounts"
	sf "github.com/RDLxxx/AVAFchain/safety/accounts"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/sha3"
)

func GetPrivateKeyFromAP(ka *accounts.KeyAccount, password string) (*crp.PrivateKey, error) {
	salt, err := hex.DecodeString(ka.KDFParams.Salt)
	if err != nil {
		return nil, fmt.Errorf("failed to decode salt: %v", err)
	}

	key, err := scrypt.Key([]byte(password), salt, ka.KDFParams.N, ka.KDFParams.R, ka.KDFParams.P, ka.KDFParams.DKLen)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %v", err)
	}

	iv, err := hex.DecodeString(ka.CipherParams.IV)
	if err != nil {
		return nil, fmt.Errorf("failed to decode IV: %v", err)
	}

	ciphertext, err := hex.DecodeString(ka.CipherText)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ciphertext: %v", err)
	}

	mac := sf.CreateMAC(key[16:32], ciphertext)
	if hex.EncodeToString(mac) != ka.MAC {
		return nil, fmt.Errorf("invalid MAC: password may be incorrect")
	}

	privateKeyBytes, err := DecryptPrivateKey(ciphertext, key[:16], iv)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt private key: %v", err)
	}

	privateKey, err := x509.ParseECPrivateKey(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	return privateKey, nil
}

func DecryptPrivateKey(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

func IsGoodPrv(account accounts.Account, prvKey crp.PrivateKey) bool {
	pubKeyFromPrv := prvKey.PublicKey
	chkgAdrs := account.Address
	hash := sha3.NewLegacyKeccak256()
	hash.Write(elliptic.Marshal(elliptic.P256(), pubKeyFromPrv.X, pubKeyFromPrv.Y))

	pkpvAdrs := "AVAFu" + hex.EncodeToString(hash.Sum(nil)[12:])
	return chkgAdrs == pkpvAdrs
}
