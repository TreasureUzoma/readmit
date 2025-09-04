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
	// Return the UUID as a hexadecimal string
	return hex.EncodeToString(uuid), nil
}