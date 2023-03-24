package helper

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(stringPassword *string) string {
	var newSha = sha256.New()
	newSha.Write([]byte(*stringPassword))
	encryptedString := fmt.Sprintf("%x", newSha.Sum(nil))
	return encryptedString
}
