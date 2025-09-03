package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(uuid), nil
}