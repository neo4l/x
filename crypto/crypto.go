package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func PBKDF2WithHmacsha256(password, salt string, iter, keyLen int) []byte {
	return pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, sha256.New)
}

func EncPasswd(password, salt string) []byte {
	return PBKDF2WithHmacsha256(password, salt, 3000, 32)
}
