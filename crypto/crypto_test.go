package crypto

import (
	"crypto/aes"
	//"encoding/base58"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"go-anyth/crypto/address/base58"
	"log"
	"testing"
	"time"

	"github.com/neo4l/x/randentropy"
)

func Test_MD5(t *testing.T) {
	log.Printf("Method: Test_MD5...")
	t1 := time.Now().UnixNano()
	pwd := Md5("123456")
	t.Logf("pwd: %s,%d", pwd, time.Now().UnixNano()-t1)

	bytes, _ := hex.DecodeString("dfe8f2f7e5f94072d753017f63420763f9ea7d455ad1dcbd3eff199ece8a5c56")
	fmt.Println(base58.Encode(bytes))

}

func Test_AesCBCEncrypt(t *testing.T) {
	iv := randentropy.GetEntropyCSPRNG(aes.BlockSize)
	srcText := "0123456789012345"
	key := randentropy.GetEntropyCSPRNG(aes.BlockSize)
	log.Printf("key: %s", hex.EncodeToString(key))
	log.Printf("iv: %s", hex.EncodeToString(iv))

	encText, err := AesCBCEncrypt(key, []byte(srcText), iv)

	log.Printf("encText: %s, %s", hex.EncodeToString(encText), err)

	decText, err := AesCBCDecrypt(key, encText, iv)

	log.Printf("decText: %s, %s", string(decText), err)
}

func Test_EncPasswd(t *testing.T) {
	pw := EncPasswd("9b99031f979b46279de69c5fd9250006", "212aeafdbadcf2e2067d8edf")
	fmt.Printf("passwd: %s\n", base64.StdEncoding.EncodeToString(pw))

	fmt.Printf("base64: %s\n", base64.StdEncoding.EncodeToString(PBKDF2WithHmacsha256("1111", "1234", 3000, 32)))

	hm, err := HmacSHA256Sign("123", "123")
	fmt.Printf("hamac256: %s,%s\n", hm, err)

}
