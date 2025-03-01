package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func Hash(baseHash string, keys []string) ([]string, error) {
	var concatenatedHashes []string

	for _, key := range keys {
		combined := baseHash + key
		hash := sha256.Sum256([]byte(combined))
		concatenatedHashes = append(concatenatedHashes, hex.EncodeToString(hash[:]))
	}

	return concatenatedHashes, nil
}

func Encrypt(plaintext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using AES-GCM.
func Decrypt(ciphertext string, key []byte) ([]byte, error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()
	if len(decodedCiphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertextOnly := decodedCiphertext[:nonceSize], decodedCiphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertextOnly, nil)

	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
