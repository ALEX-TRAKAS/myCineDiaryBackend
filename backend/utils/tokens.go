package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func SHA256(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	return hex.EncodeToString(h.Sum(nil))
}

func HashToken(token string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(token), salt, 1, 64*1024, 4, 32)
	return fmt.Sprintf("%x.%x", salt, hash), nil
}

func CompareHashAndToken(hash string, token string) bool {
	parts := strings.Split(hash, ".")
	if len(parts) != 2 {
		return false
	}

	salt, _ := hex.DecodeString(parts[0])
	hashBytes, _ := hex.DecodeString(parts[1])

	newHash := argon2.IDKey([]byte(token), salt, 1, 64*1024, 4, 32)
	return subtle.ConstantTimeCompare(newHash, hashBytes) == 1
}
