package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256Bytes(content []byte) []byte {
	h := sha256.New()
	h.Write(content)
	sum := h.Sum(nil)
	return sum
}

func HashSHA256(content []byte) string {
	return hex.EncodeToString(HashSHA256Bytes(content))
}

func HashSHA256Str(str string) string {
	content := []byte(str)
	return hex.EncodeToString(HashSHA256Bytes(content))
}
