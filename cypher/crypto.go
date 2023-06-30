package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/sword/api-backend-challenge/config"
	"github.com/sword/api-backend-challenge/log"
)

type Crypto struct {
	block cipher.Block
}

func NewCrypto() *Crypto {
	logger := log.NewEntry()
	block, err := aes.NewCipher([]byte(config.GetEnv().Security.SecretKey))
	if err != nil {
		logger.WithError(err).Fatal()
	}
	return &Crypto{block}
}

// Encrypt method is to encrypt or hide any classified text
func (c *Crypto) Encrypt(text string) (string, error) {
	plaintext := []byte(text)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(c.block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt method is to extract back the encrypted text
func (c *Crypto) Decrypt(text string) (string, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(text)
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(c.block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
