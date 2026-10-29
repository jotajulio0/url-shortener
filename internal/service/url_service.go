package service

import (
	"crypto/sha256"
	"encoding/binary"
)

const (
	codeSize = 6
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func base62Encode(num uint64) string {
	result := make([]byte, codeSize)
	base := uint64(len(charset))

	for i := 0; i < codeSize; i++ {
		result[i] = charset[num%base]
		num /= base
	}
	return string(result)
}

func GenerateShortCode(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))
	num := binary.BigEndian.Uint64(hash[:8])
	return base62Encode(num)
}
