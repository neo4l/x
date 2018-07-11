package tool

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"sort"

	"github.com/tealeg/xlsx"
)

func IsEmpty(obj interface{}) bool {
	if obj == nil {
		return true
	}
	switch v := obj.(type) {
	case string:
		return v == ""
	}
	return true

}

// func HexToInt(hexStr string) (int64, error) {
//  log.Printf("HexToInt: %s", hexStr)
//  if IsEmpty(hexStr) {
//      return 0, nil
//  } else {
//      return strconv.ParseInt(hexStr, 0, 0)
//  }
// }

// // func IntToHex(num int) (int, error) {
// //   if IsEmpty(hexStr) {
// //       return 0, nil
// //   } else {
// //       return
// //   }
// // }

func AToInt64WithoutErr(str string) int64 {
	num, _ := strconv.ParseInt(str, 0, 64)
	return num
}

func StrToHex(str string) (string, error) {
	if IsEmpty(str) {
		return "0x", nil
	} else if strings.HasPrefix(str, "0x") {
		return str, nil
	}
	bigInt, ok := new(big.Int).SetString(str, 10)
	if ok {
		return fmt.Sprintf("0x%x", bigInt), nil
	}
	return "", fmt.Errorf("parse error")
}

func StrToHexWithoutError(str string) string {
	hexStr, err := StrToHex(str)
	if err != nil {
		return "0x"
	}
	return hexStr
}

func HexToInt(hexStr string) (int64, error) {
	if IsEmpty(hexStr) || hexStr == "0x" {
		return 0, nil
	}
	return strconv.ParseInt(Strip0x(hexStr), 16, 64)
}

func HexToIntWithoutError(hexStr string) int64 {
	reply, err := HexToInt(hexStr)
	if err != nil {
		return 0
	}
	return reply
}

func HexToUintWithoutError(hexStr string) uint64 {
	if IsEmpty(hexStr) || hexStr == "0x" {
		return 0
	}
	reply, err := strconv.ParseUint(Strip0x(hexStr), 16, 64)
	if err != nil {
		return 0
	}
	return reply
}

func HexToBigInt(hexStr string) *big.Int {
	bigInt := new(big.Int)
	if IsEmpty(hexStr) {
		bigInt.SetString("0", 0)
	} else if strings.HasPrefix(hexStr, "0x") {
		bigInt.SetString(hexStr, 0)
	} else {
		bigInt.SetString(hexStr, 16)
	}
	return bigInt
}

func IntToHex(num int64) string {
	return "0x" + strconv.FormatInt(num, 16)
}

func HexToIntStr(hexStr string) string {
	return HexToBigInt(hexStr).String()
}

func Strip0x(input string) string {
	if len(input) >= 2 && strings.HasPrefix(input, "0x") {
		return Substr(input, 2, len(input))
	}
	return input
}

func Add0x(input string) string {
	if len(input) < 2 || !strings.HasPrefix(input, "0x") {
		return "0x" + input
	}
	return input
}

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func ParseLogData(hexStr string) (data []string) {
	//log.Printf("Method: ParseEventData... %s", hexStr)
	dataArray := []string{}
	if len(hexStr) <= 2 {
		return append(dataArray, hexStr)
	}
	hexStr = Strip0x(hexStr)
	len := len(hexStr)
	//log.Printf("hexStr: %d,%s", len, hexStr)
	if len > 0 && len%64 == 0 {
		n := len / 64
		for index := 0; index < n; index++ {
			dataArray = append(dataArray, Substr(hexStr, index*64, 64))
		}
	}
	return dataArray
}

func CurDate() string {
	//date := time.Now()
	return time.Now().Format("20060102") //fmt.Sprintf("%d%d%d", date.Year(), date.Month(), date.Day())
}

func ParseTime(format, timeStr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	return time.ParseInLocation(format, timeStr, loc)
}

func ReadJsonFile(fileName string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}
	var jsonObj map[string]interface{}
	if err := json.Unmarshal(bytes, &jsonObj); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return jsonObj, nil
}

func ID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	bytes := Md5Hash(b)
	return hex.EncodeToString(bytes)
}

func UUID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	bytes := Md5Hash(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
}

func Md5Hash(b []byte) []byte {
	h := md5.New()
	h.Write(b)
	return h.Sum(nil)
}

func CreateSalt() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func ParamToStringWithSort(params map[string]string) string {
	if len(params) <= 0 {
		return ""
	}
	sortedKeys := make([]string, 0)
	for key, _ := range params {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)
	reply := ""
	for _, key := range sortedKeys {
		if len(reply) > 0 {
			reply += "&"
		}
		reply += key + "=" + params[key]
	}

	return reply
}

func LeftPadString(stringToPad, padChar string, length int) string {
	var repreatedPadChar = ""
	count := length - len(stringToPad)
	for index := 0; index < count; index++ {
		repreatedPadChar += padChar
	}
	return stringToPad + repreatedPadChar
}

func RightPadString(stringToPad, padChar string, length int) string {
	var repreatedPadChar = ""
	count := length - len(stringToPad)
	for index := 0; index < count; index++ {
		repreatedPadChar += padChar
	}
	return repreatedPadChar + stringToPad
}

func ToEther(hexValue string) string {
	return ToBalance(HexToIntStr(hexValue), 18)
}

func ToBalance(value string, decimals int) string {
	val := RightPadString(value, "0", decimals+1)
	prefixVal := Substr(val, 0, len(val)-decimals)
	return prefixVal + "." + Substr(val, len(val)-decimals, decimals)
}

func ToValue(value string, decimals int) string {
	if strings.HasPrefix(value, ".") {
		value = Substr(value, 1, len(value)-1)
	}
	if strings.HasSuffix(value, ".") {
		value = Substr(value, 0, len(value)-1)
	}
	if strings.Contains(value, ".") {
		index := strings.Index(value, ".")
		suffix := Substr(value, index+1, len(value)-index-1)
		suffix = LeftPadString(suffix, "0", decimals)
		return Substr(value, 0, index) + suffix
	}
	return LeftPadString(value, "0", decimals+len(value))
}

func EtherToHex(ether string) string {
	if ether == "" || ether == "0" {
		return "0x"
	}
	index := strings.Index(ether, ".")
	value := ""
	if index < 0 {
		value = ether + "000000000000000000"
	} else {
		suffix := Substr(ether, index+1, len(ether)-index-1)
		suffix = LeftPadString(suffix, "0", 18)
		value = Substr(ether, 0, index) + suffix
	}
	return StrToHexWithoutError(value)
}

func GWeiToHex(gwei string) string {
	if gwei == "" || gwei == "0" {
		return ""
	}
	index := strings.Index(gwei, ".")
	value := ""
	if index < 0 {
		value = gwei + "000000000000000000"
	} else {
		suffix := Substr(gwei, index+1, len(gwei)-index-1)
		suffix = LeftPadString(suffix, "0", 9)
		value = Substr(gwei, 0, index) + suffix
	}
	return StrToHexWithoutError(value)
}

func EtherToWei(ether string) string {
	// val := new(big.Int).Mul(HexToBigInt(hexValue), new(big.Int).SetInt64(1000000000000000000))
	// return val.String()
	hexValue := EtherToHex(ether)
	return HexToIntStr(hexValue)
}

func WeiToEther(value string) string {
	val := new(big.Int).Div(HexToBigInt(value), new(big.Int).SetInt64(1000000000000000000))
	return val.String()
}

func WeiToGWei(value string) string {
	val := RightPadString(HexToIntStr(value), "0", 10)
	prefixVal := Substr(val, 0, len(val)-9)
	return prefixVal + "." + Substr(val, len(val)-9, 9)
}

func GWeiToWei(gwei string) string {
	// val := new(big.Int).Mul(HexToBigInt(value), new(big.Int).SetInt64(1000000000))
	// return val.String()
	hexValue := GWeiToHex(gwei)
	return HexToIntStr(hexValue)
}

func ToJson(obj interface{}) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil
	} else {
		return bytes
	}
}

func ReadLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)

	if file, err = os.Open(path); err != nil {
		return
	}

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))

	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func ExportExcl(items []string, fileName string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	for _, item := range items {
		row := sheet.AddRow()
		values := strings.Split(item, ",")
		for _, value := range values {
			cell := row.AddCell()
			cell.Value = value
		}
	}

	return file.Save(fileName)
}
