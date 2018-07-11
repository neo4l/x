package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("./rsa_private_key_bhbe.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	//openssl pkcs8 -topk8 -in rsa_private_key.pem -out pkcs8_rsa_private_key.pem -nocrypt

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("./rsa_public_key_bhbe.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

// func main() {

// 	// Generate RSA Keys
// 	miryanPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

// 	if err != nil {
// 		fmt.Println(err.Error)
// 		os.Exit(1)
// 	}

// 	miryanPublicKey := &miryanPrivateKey.PublicKey

// 	raulPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

// 	if err != nil {
// 		fmt.Println(err.Error)
// 		os.Exit(1)
// 	}

// 	raulPublicKey := &raulPrivateKey.PublicKey

// 	fmt.Println("Private Key : ", miryanPrivateKey)
// 	fmt.Println("Public key ", miryanPublicKey)
// 	fmt.Println("Private Key : ", raulPrivateKey)
// 	fmt.Println("Public key ", raulPublicKey)

// 	//Encrypt Miryan Message
// 	message := []byte("the code must be like a piece of music")
// 	label := []byte("")
// 	hash := sha256.New()

// 	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, raulPublicKey, message, label)

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(message), ciphertext)
// 	fmt.Println()

// 	// Message - Signature
// 	var opts rsa.PSSOptions
// 	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
// 	PSSmessage := message
// 	newhash := crypto.SHA256
// 	pssh := newhash.New()
// 	pssh.Write(PSSmessage)
// 	hashed := pssh.Sum(nil)

// 	signature, err := rsa.SignPSS(rand.Reader, miryanPrivateKey, newhash, hashed, &opts)

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("PSS Signature : %x\n", signature)

// 	// Decrypt Message
// 	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, raulPrivateKey, ciphertext, label)

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", ciphertext, plainText)

// 	//Verify Signature
// 	err = rsa.VerifyPSS(miryanPublicKey, newhash, hashed, signature, &opts)

// 	if err != nil {
// 		fmt.Println("Who are U? Verify Signature failed")
// 		os.Exit(1)
// 	} else {
// 		fmt.Println("Verify Signature successful")
// 	}
// }

func EncryptOAEP(publicKey *rsa.PublicKey, message []byte) ([]byte, error) {

	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, []byte(""))
}

func DecryptOAEP(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {

	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, []byte(""))
}

func ToPublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pubInterface.(*rsa.PublicKey), nil
}

func ToPrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func SignPSS(privateKey *rsa.PrivateKey, message []byte) ([]byte, error) {

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(message)
	hashed := pssh.Sum(nil)

	return rsa.SignPSS(rand.Reader, privateKey, newhash, hashed, &opts)
}

func VerifyPSS(publicKey *rsa.PublicKey, message, sign []byte) error {

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(message)
	hashed := pssh.Sum(nil)

	return rsa.VerifyPSS(publicKey, newhash, hashed, sign, &opts)
}

func Encrypt(publicKey *rsa.PublicKey, origData []byte) ([]byte, error) {

	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, origData)
}

func Decrypt(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {

	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
}

func Sign(privateKey *rsa.PrivateKey, message []byte) ([]byte, error) {

	newhash := crypto.SHA256
	hash := newhash.New()
	hash.Write(message)
	hashed := hash.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, privateKey, newhash, hashed)
}

func Verify(publicKey *rsa.PublicKey, message, sign []byte) error {

	newhash := crypto.SHA256
	hash := newhash.New()
	hash.Write(message)
	hashed := hash.Sum(nil)

	return rsa.VerifyPKCS1v15(publicKey, newhash, hashed, sign)
}
