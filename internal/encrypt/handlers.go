package encrypt

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var key []byte = make([]byte, 32)

type EncryptBody struct {
	Text string `json:"text"`
}
type DecryptBody struct {
	Encrypted string `json:"encrypted"`
}
type HashBody struct {
	Base string `json:"base"`
}

func EncryptHandler(c *gin.Context) {
	var encryptBody EncryptBody
	if err := c.ShouldBindJSON(&encryptBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if encryptBody.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "encryptBody.Text is required"})
		return
	}

	ciphertext, err := Encrypt([]byte(encryptBody.Text), key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "encryption failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ciphertext": ciphertext})

}

func DecryptHandler(c *gin.Context) {
	var decryptBody DecryptBody
	if err := c.ShouldBindJSON(&decryptBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if decryptBody.Encrypted == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ciphertext is required"})
		return
	}

	plaintext, err := Decrypt(decryptBody.Encrypted, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "decryption failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plaintext": string(plaintext)})
}

func HashingHandler(c *gin.Context) {
	var hashBody HashBody
	if err := c.ShouldBindJSON(&hashBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if hashBody.Base == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "encryptBody.Text is required"})
		return
	}

	baseHash := hashBody.Base // Replace with your base hash
	var keys []string

	for i := 1; i <= 10000; i++ {
		keys = append(keys, strconv.Itoa(i))
	}

	hashes, err := Hash(baseHash, keys)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "encryption failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"hashes": hashes})

}
