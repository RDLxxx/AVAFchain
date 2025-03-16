package accounts

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// IV
func GenerateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	return iv, nil
}

// SALT
func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}
	return salt, nil
}

// AES-128-CTR
func EncryptPrivateKey(privateKey, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(privateKey))
	stream.XORKeyStream(ciphertext, privateKey)

	return ciphertext, nil
}

// Message Authentication Code
func CreateMAC(key, ciphertext []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(ciphertext)
	return mac.Sum(nil)
}
