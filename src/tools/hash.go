package tools

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateHash(s string) string {
	return hex.EncodeToString(sha256.New().Sum([]byte(s)))
}

func GenerateSSID(id, token int, email, password string) string {
	return fmt.Sprintf("%d%d%s%s", id, token, email, password)
}

func HashPassword(password string, token int) string {
	return GenerateHash(fmt.Sprintf("%s%d", password, token))
}
