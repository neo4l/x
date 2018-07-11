package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto/sha3"
)

//生成32位md5字串
func Md5(s string) string {
	return hex.EncodeToString(Md5Hash([]byte(s)))
}

func Md5Hash(b []byte) []byte {
	h := md5.New()
	h.Write(b)
	return h.Sum(nil)
}

func SHA256(s string) string {
	return hex.EncodeToString(SHA256Hash([]byte(s)))
}

func SHA256Hash(b []byte) []byte {
	hash := sha256.New()
	hash.Write(b)
	return hash.Sum(nil)
}

func SHA3_256(b []byte) (h []byte) {
	hash := sha3.NewKeccak256()
	hash.Write(b)
	return hash.Sum(h[:0])
}

func SHAHash(b []byte) ([]byte, error) {
	sha := sha1.New()
	_, err := sha.Write(b)
	if err != nil {
		return nil, err
	}
	return sha.Sum(nil), nil
}

func HmacSHA256Sign(secret, data string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func HmacSHA512Sign(secret, data string) (string, error) {
	mac := hmac.New(sha512.New, []byte(secret))
	_, err := mac.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func HmacSHA1Sign(secret, data string) (string, error) {
	mac := hmac.New(sha1.New, []byte(secret))
	_, err := mac.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func HmacMD5Sign(secret, data string) (string, error) {
	mac := hmac.New(md5.New, []byte(secret))
	_, err := mac.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}
