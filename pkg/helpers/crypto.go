package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"app/core/config"
)

func newGCM() (cipher.AEAD, error) {
	block, err := aes.NewCipher(config.GetAppKey())
	if err != nil {
		return nil, err
	}

	return cipher.NewGCM(block)
}

func Encrypt(plaintext string) (string, error) {
	gcm, err := newGCM()
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encoded string) (string, error) {
	gcm, err := newGCM()
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("ciphertext too short")
	}

	nonce, body := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, body, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
