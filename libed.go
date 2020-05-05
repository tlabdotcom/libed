package libed

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

func createHash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

// GCMEncrypter is ...
func GCMEncrypter(data []byte, passphrase string) (string, error) {
	block, err := aes.NewCipher(createHash(passphrase))
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// GCMDecrypter is used to ...
func GCMDecrypter(data, passphrase string, res interface{}) error {
	// Decoding from base64
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	key := createHash(passphrase)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := aesgcm.NonceSize()
	if len(decodedData) < nonceSize {
		return fmt.Errorf("authentication failed")
	}
	nonce, ciphertext := decodedData[:nonceSize], decodedData[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal(plaintext, res)
	if err != nil {
		return err
	}

	return nil
}
