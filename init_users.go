package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"go-users-simple-crud/model"
	"strings"
	"time"
)

const (
	key              = "I5Fkl7pS2cQcEEKiR6icrxHgA+cX09vtGEMVXiXEXwk="
	encryptedContent = `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`
)

func initDatabase(container *Container) {
	content, err := decryptContent(encryptedContent, key)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(content, "\n")
	for _, line := range lines[1:] {
		columns := strings.Split(line, ",")
		birthDate, err := time.Parse("02-01-2006", columns[2])
		if err != nil {
			panic(err)
		}
		user := model.User{
			Name:      columns[0],
			Email:     columns[1],
			BirthDate: birthDate,
		}
		if _, err = container.UserInMemoryRepository.Save(user); err != nil {
			panic(err)
		}
	}
}

func decryptContent(ciphertextStr, key string) (string, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextStr)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(decodedKey)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short: %w", err)
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
