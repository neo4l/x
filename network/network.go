package network

import (
	"fmt"
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"andui/crypto/rsa"
	"andui/crypto"
	//"encoding/hex"
)

func Get(url string, params map[string]string) ([]byte, error) {
	paramStr := ParamToString(params)
	log.Printf("http get: %s\nParam: %s", url, paramStr)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(url + "?" + paramStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	reply, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return reply, err
}

func Post(url string, params map[string]string) ([]byte, error) {
	paramStr := ParamToString(params)
	log.Printf("http post: %s\nParam: %s", url, paramStr)
	body := bytes.NewBuffer([]byte(paramStr))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	reply, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return reply, err
}

func ParamToString(params map[string]string) string {
	pstr := ""
	if params != nil && len(params) > 0 {
		var keys []string
		for k := range params {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			if pstr != "" {
				pstr += "&"
			}
			pstr += key + "=" + params[key]
		}
	}
	return pstr
}

func Hamcsha256SignParam(secret string, params map[string]string) (sign string,err error) {
	data := ParamToString(params)
	fmt.Printf("data: %s\n", data)
	return Hamcsha256Sign(secret, data)
}

func Hamcsha256Sign(secret, data string) (sign string, err error) {
	return crypto.HmacSHA256Sign(secret, data)
}

func RSASignParam(privateKey string, params map[string]string) ([]byte,error) {
	pk, err := rsa.ToPrivateKey([]byte(privateKey))
	if err != nil {
		return nil,err
	}
	return rsa.Sign(pk, []byte(crypto.SHA256(ParamToString(params))))
}

func RSAVerifySign(publicKey string, params map[string]string, sign string) bool {
	pk, err := rsa.ToPublicKey([]byte(publicKey))
	if err != nil {
		return false
	}
	msg := crypto.SHA256(ParamToString(params))
	fmt.Printf("pdata: %s\n", ParamToString(params))
	fmt.Printf("hash1: %s\n", string(msg))
	fmt.Printf("sign: %s\n", sign)
	signBytes := []byte(sign)
	// signBytes,err := hex.DecodeString(sign)
	// if err != nil {
	// 	fmt.Printf("DecodeString: %s\n", err)
	// 	return false
	// }
	return rsa.Verify(pk, []byte(msg), signBytes) == nil
}