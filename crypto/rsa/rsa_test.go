package rsa

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"testing"
)

var (
	publicKeyStr = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzelMahQy9iAXw9xCVxQg
sdW/lOxf6Pn91BcTCPf+vnfMnboqRIUazEjLJMSXeqBPLpLS/su4lxWmTw9twKg1
IfbE/6n6euCz5v+6R3BaVU7g2iopoMAoJqrhAoPsugw/g3WaCyzOBCq7sjRkODXd
DRR2kqoKb4HVCjItf6lU8TCIQe2otrEcoXexcbuoMUFAoHXt1kHJ5PHJO2KaxFK+
notQ98WUOWnBuiHDfsWqIKocaE0g8agjp4+I4eSTy9qQMQRNwoQhnuVHBUh3BcFg
gEqnHjdBWNba2eJ9VDbaBq30H5dOXgXNaVqh6ZZqtbWdpkGu/SgTGzi3N3gFUHqC
eQIDAQAB
-----END PUBLIC KEY-----`

	privateKeyStr = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzelMahQy9iAXw9xCVxQgsdW/lOxf6Pn91BcTCPf+vnfMnboq
RIUazEjLJMSXeqBPLpLS/su4lxWmTw9twKg1IfbE/6n6euCz5v+6R3BaVU7g2iop
oMAoJqrhAoPsugw/g3WaCyzOBCq7sjRkODXdDRR2kqoKb4HVCjItf6lU8TCIQe2o
trEcoXexcbuoMUFAoHXt1kHJ5PHJO2KaxFK+notQ98WUOWnBuiHDfsWqIKocaE0g
8agjp4+I4eSTy9qQMQRNwoQhnuVHBUh3BcFggEqnHjdBWNba2eJ9VDbaBq30H5dO
XgXNaVqh6ZZqtbWdpkGu/SgTGzi3N3gFUHqCeQIDAQABAoIBAQCsAPtVDWR4ltxj
PuWnyed5xhzQMRf5DIMNHO1Iq6h/wKELDIzsSefVx+Tx5MrIo4shU4KvVsvuYSZY
moHK0nf31CRBkOLsrDF7gBlCPccnxcksVNYLMxkXG9zz9fHUhBC2JpG0TgwWDQBX
X05sagoqN/LIlwQ6m1CzwLCjGHcdNTW0ZHDK9DA+Wcmupm/06G/upEredLyOvsVf
giQ2kcxzoGsRaNsE+cGWa9p9wqKJTHnFuXS+roluq0s0EQX+eDB0bZxclSSNdSyv
85yj/6wY0g1dJGNmDED8EoL3JhIglnxTMzJ4og1hCo8PUSKXAZWUT7VJXsQExgO6
J1tx7wnFAoGBANV3RShzBMxfHSQ9NxKKqQz9MR66/XWU41Yk1dFEjM2oS+e9YWEH
SjriS8O/H099OtYcGNYkJj58+GbxK73+bJgRZ1fkAIXBRZXXQH7stg3BiZeFWOwS
+R/7C6SBCDLrLKhBMTbNzyoxDrFyB5BC74GalPpA7DA/OlIgMP2uLJUbAoGBAPbw
rNzzMgGCT3DStZt3ZBACXQxVYVX+cNPUG5O0hV8wG3T8JE+2U2bzToPfXjn/qL2L
kOl5z50QC3HoIqmdT3Km2WpowUhwZKKnpPOzBrZGu8HZa5cKwDtv5exGYrUMgp2P
KF9ALDWkNwIQckI1XWbEe3/8sgne3pSGmWwQXwP7AoGAJS5EDnqSMGK0ubYr3H/o
WAnVv2uEcDGBs2RxFaUh/UQ+DFwxFOxnIoB9/9dPRdIjKF32eX4MZz/vKEcDfnFN
SuNlI6rj6Gg0jZfTdQgX4ad/JrQkO+JGICri6UFMQ0oxGhFY2Bna5pdq3r9kz9zI
yMM7Bae/O9wXdWyD+/uValcCgYA2munewfhg1Qv0CuQVyMTbtWoV/BtWBLm2XcTr
WJPVhLHNoKP27H5s2YiXKKGRebM6ls4oksMSHCYrvgVMNHkJBVQ2b4uuFQxr215i
dUgarnF+YDGmaL4xZoEVSksxdd68MJfg2DPueK2hSzm44kwRGYmlt583B414knsC
pmwcWQKBgGwF9/PEHe+EKwMYOHSlXBSrt8psXMY1oE4qSmymXdsPuyqJsWiR+HIx
/7w/zDSZdE5j5gUugspTFv9jbPukVvZ/Sl+7NxO5PR1hqVx6n4ld2y2ZoPpOiRJ7
azgM2kitxrYGna7AzJylcbc7zSdZtoKyroCB2gVBg7BUK4cjSBud
-----END RSA PRIVATE KEY-----`

publicKeyStr2=`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw3TGU0B0AnYFq1lwmwviCh5Ig7pAx9YgSKAOjjBXCJhU8YYL/QzynLRcYS0JEmQQfkKXYd+4UXhmhxoHmBTDclafXxR/UJayCbmXPJaw1Y4qkaKvOpSEnV2WK0rGjEWE7mGBKkaCo2dk28Hso4CoQkZBgneMMi9Ff02GFDhTc0N2jEsoFtbgtpD2zCuJj6KBV8RT21OSmQYpAtQcJ0DSS96Wmj5QM75udJW1RGz0eaBX01K6aEBMKaHCfe+wy15l3GjPuC3qeKQpaBk2LK9FXVNYVLGMyV10itADdekZz/RNpULsWnTA+7GRGk7V/pHmVkt9mvm6tPkcAMCTAmhDSwIDAQAB`
privateKeyStr2=`MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDDdMZTQHQCdgWrWXCbC+IKHkiDukDH1iBIoA6OMFcImFTxhgv9DPKctFxhLQkSZBB+Qpdh37hReGaHGgeYFMNyVp9fFH9QlrIJuZc8lrDVjiqRoq86lISdXZYrSsaMRYTuYYEqRoKjZ2TbweyjgKhCRkGCd4wyL0V/TYYUOFNzQ3aMSygW1uC2kPbMK4mPooFXxFPbU5KZBikC1BwnQNJL3paaPlAzvm50lbVEbPR5oFfTUrpoQEwpocJ977DLXmXcaM+4Lep4pCloGTYsr0VdU1hUsYzJXXSK0AN16RnP9E2lQuxadMD7sZEaTtX+keZWS32a+bq0+RwAwJMCaENLAgMBAAECggEBAJTFHmuhpGt3H8uTkUVCXzOqZzF0o/g3QG1m/o01uBk/sXAAwsFCiKvEEIiaucv4xCEGWKlTmrzQMuHkayFTtIwj6Lx7IojZmvwR+k5QuJpj0nz1YgDpqKMK59Fd4hBEnfC/4IuoDamIellWmuK8e9WuGhDPI7PTDUffXw2m7cXNoALvJMu42c6PU7X3ueb3A1/T6QqY1N0XKrLpxNpeRYCN2mU/LhB55awwpcbNTNgkbLIzTgMg0IN+JLi88iaO7bv8awie7pSCFDwH23dMlM0/Ua2cyj7m9x8y2bA+Orp1tbFnN9SalrmhuqcbQ1mRsDLTlleIgOu6s7Gy9zioQwkCgYEA57YBxPzdtJN/R/fPqrTSxku3OkMBW0/cKKgr/0VNZSRbr95GB2FX+exTW6F86Qt6kulQtgSriXJ8RuDpEM0TYryiXEVUak6hOaQ33lpE0sS9R5u8gDXtoeaOwrhU3roNmLOFF4ukZw6Gs7XNnpIWzG6UYHFcgzEVk1488/XhHH8CgYEA1/Hdg4AiBo7thXq7MXDLesedR4Dbtyu0yEF3ZVc7Z0oS1snhOYkPElyS6Z95IffALpGVVzO16kwJu6Vki+N/lo/iFjR4YFfopYkEAIkye4lw5lJ6lwDQcWo09loCGiUQZLHKwfLak+Ley/XC9de5aPyb2hDo0GwVVnO081gYIzUCgYEAq2lnkbSGxqk+xZy7kctHCc7Fc2JSRJylf6Y5NhSslp/4+/dw0tDeZlK/r8+dOkF1oezb+msmAv4glcaYZAdxyd9GNQBM3H/ioWOsuN8KfulwJOM+5ZH/g3+uKLp4fnQgztAvKyXwrxR97cAWprHoD7/WICp8h8jt7yEN8mP47j0CgYAD8onoE3mLwSUaYYn2d2dg0TFQ00ww5v2hA4FZOuT9GF+LyZjyYk0COur8lkuykULUFxkxxOI4bDdpVLanz/rPF8Y8Pa1NpY29KOoH0Ho5w+NqcmuHQx6MVDKvpimPrMnF7XIVZYkVVvpXpCByOgVLpAJ9U/3NgYxKTkcqg5u9WQKBgQDMWhQHwOg2++mlVxkyWZkm8dDhaCHqyO41t5EDqLy6NXTM8kvRY5b6g/vCXb01eNr5NPnYm8dVZSvcrhSOoEwmt4eYGTdYckzFDYj3ot/XnC4VpaY9x3Q+PSrInstsdly3vC8whNs5RZVMBKzDoTL+g4EyoHs07eA0Yet2VruTfw==`

)

func Test_GenRsaKey(t *testing.T) {
	GenRsaKey(2048)
}

func Test_Sign2(t *testing.T) {
	pubk,err := base64.StdEncoding.DecodeString(publicKeyStr2)
	publicKey, err := ToPublicKey(pubk)

	if err != nil {
		t.Errorf("ToPublicKey err,%s", err)
		return
	}
	t.Logf("PublicKey: %s", publicKey)

	prik,err := base64.StdEncoding.DecodeString(privateKeyStr2)
	privateKey, err := ToPrivateKey(prik)

	if err != nil {
		t.Errorf("ToPrivateKey err,%s", err)
		return
	}
	t.Logf("PrivateKey: %s", privateKey)

	message := "cae8fb30897bf4a0d1dbc9572a391026c11581b8218479b6e4aa9c2337899469"

	sign, err := Sign(privateKey, []byte(message))
	if err != nil {
		t.Errorf("sign err,%s", err)
		return
	}
	log.Printf("sign: %s", hex.EncodeToString(sign))


	err = Verify(publicKey, []byte(message), sign)
	if err != nil {
		t.Errorf("Verify err,%s", err)
		return
	}
	t.Logf("Verify: %d,%s", err == nil, err)

	message2 := "cae8fb30897bf4a0d1dbc9572a391026c11581b8218479b6e4aa9c2337899469"
	sign2Hex := `618d2275d087e945040e92756e57f392ce189a0a24aaafd299cca1a1625c4a9411f79331ae0bf61e519059eade80ef5d9f40c8fd61028ac7a578018146c8284af7ddbcb3ad32bcffa6e6675fe85543b3ebece3bb57f2bb8e041c871063f04fb6b430ba2f5081dc02a6f2c534f2d60098b1b93f9d2d67d47d3110653437cfb701d852a6f8361db969bc198c04e225a1fdd3b0e5e051b130cf544cb519f2f227f00b077154453b90aaa9efeb9cf4a1e6ee60b0ec632cbbf6c5d685ab083a604e2c8cf6512c755d45dc395bc7c6fab213ee6a72c3fdcced5e04fe360c1307eec62cb1cad7dc8f6af357c0a012583f34a1b3b6c8cc3dd9d2ccddab10876eab11668d`
	sign2, err := hex.DecodeString(sign2Hex)
	err2 := Verify(publicKey, []byte(message2), sign2)
	if err != nil {
		log.Printf("Verify2 err,%s", err2)
		return
	}
	log.Printf("Verify2: %d,%s", err2 == nil, err2)
}

func Test_Sign(t *testing.T) {

	publicKey, err := ToPublicKey([]byte(publicKeyStr))

	if err != nil {
		t.Errorf("ToPublicKey err,%s", err)
		return
	}
	t.Logf("PublicKey: %s", publicKey)

	privateKey, err := ToPrivateKey([]byte(privateKeyStr))

	if err != nil {
		t.Errorf("ToPrivateKey err,%s", err)
		return
	}
	t.Logf("PrivateKey: %s", privateKey)

	message := "明文123456"

	encMsg, err := Encrypt(publicKey, []byte(message))
	if err != nil {
		t.Errorf("Encrypt err,%s", err)
		return
	}
	t.Logf("encMsg: %s", hex.EncodeToString(encMsg))

	sign, err := Sign(privateKey, []byte(message))
	if err != nil {
		t.Errorf("sign err,%s", err)
		return
	}
	log.Printf("sign: %s", hex.EncodeToString(sign))

	decMsg, err := Decrypt(privateKey, encMsg)
	if err != nil {
		t.Errorf("Decrypt err,%s", err)
		return
	}
	t.Logf("decMsg: %s", string(decMsg))

	err = Verify(publicKey, []byte(message), sign)
	if err != nil {
		t.Errorf("Verify err,%s", err)
		return
	}
	t.Logf("Verify: %d,%s", err == nil, err)

	message2 := "明文123456"
	sign2Hex := `8445158d80534f59fdee7cebcb48c530109bf763210b2db99fad7d2c1b969f029746b80f9b44bae1ea60df4b0417a682c575c4445a03703f03cd62c62f43fa9a1eec4847b623addab8c259527dcbf572350d3760b485a2eed25fc59fd6e3b108f2a4f69fd058cf36ee81afcd6cb31db08f3e2fb1c058c5d76198584555be089c7f4204309c998d9ab7d81bfcadf3f9ef2eefd499df8d75ea1d5c9ee34e3f50beab9a9f841caeb006fb86d4bf00ff1823499b19a32cfd8d9159d8a8a8f745f83de39287c740bf144d5f24f9c9a6ba3a50228ffd804fdc9ca3e192312239292c406359e3e628cb10d475f4fe16e464961e8c3e644c1285b2ba98e5bc87c1f89c01`
	sign2, err := hex.DecodeString(sign2Hex)
	err2 := Verify(publicKey, []byte(message2), sign2)
	if err != nil {
		log.Printf("Verify2 err,%s", err2)
		return
	}
	log.Printf("Verify2: %d,%s", err2 == nil, err2)
}

func Test_SignOAEP(t *testing.T) {

	publicKey, err := ToPublicKey([]byte(publicKeyStr))

	if err != nil {
		t.Errorf("ToPublicKey err,%s", err)
		return
	}
	t.Logf("PublicKey: %s", publicKey)
	publicKeyBase64Str := base64.StdEncoding.EncodeToString([]byte(publicKeyStr))
	log.Printf("publicKeyBase64Str: %s, %s", publicKeyBase64Str, err)

	privateKey, err := ToPrivateKey([]byte(privateKeyStr))

	privateKeyBase64Str := base64.StdEncoding.EncodeToString([]byte(privateKeyStr))

	log.Printf("privateKeyBase64Str: %s, %s", privateKeyBase64Str, err)

	if err != nil {
		t.Errorf("ToPrivateKey err,%s", err)
		return
	}
	t.Logf("PrivateKey: %s", privateKey)

	message := "12345678901234567890123456789012345678901234567890123456789012345678901234567890"

	encMsg, err := EncryptOAEP(publicKey, []byte(message))
	if err != nil {
		t.Errorf("Encrypt err,%s", err)
		return
	}
	t.Logf("encMsg: %s", hex.EncodeToString(encMsg))

	sign, err := SignPSS(privateKey, []byte(message))
	if err != nil {
		t.Errorf("sign err,%s", err)
		return
	}
	t.Logf("sign: %s", hex.EncodeToString(sign))

	decMsg, err := DecryptOAEP(privateKey, encMsg)
	if err != nil {
		t.Errorf("Decrypt err,%s", err)
		return
	}
	t.Logf("decMsg: %s", string(decMsg))

	err = VerifyPSS(publicKey, []byte(message), sign)
	if err != nil {
		t.Errorf("Verify err,%s", err)
		return
	}
	t.Logf("Verify: %d,%s", err == nil, err)
}
