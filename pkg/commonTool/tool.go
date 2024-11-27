package commonTool

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func BuildPassword(password, sign string) string {
	hash := sha256.Sum256([]byte(password + "@" + sign))
	return hex.EncodeToString(hash[:])
}

func CheckPassword(checkPassword, password, sign string) bool {
	hash := sha256.Sum256([]byte(password + "@" + sign))
	return checkPassword == hex.EncodeToString(hash[:])
}

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}
