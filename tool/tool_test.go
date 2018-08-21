package tool

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func Test_IntToHex(t *testing.T) {
	//var hexStr = "0x00000000000000000000000000000000000000000000000000000000000000050000000000000000000000002333e8406e0a80700f9c787cb96d1053ccd437550000000000000000000000000000000000000000000000000000000059060a00000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000f4240000000000000000000000000ac00155169e278c2f1af7435e0b5407e73040dd9"
	//dMap := ParseEventData(hexStr)
	//log.Printf("dMap: %s", IntToHex(19))
	//bigInt := new(big.Int)
	//bigInt := HexToBigInt(hexStr)

	//hex :=

	//text, err := bigInt.MarshalText()
	//log.Printf("bigNum: %s", HexToIntStr(hexStr))

	fmt.Println(ToValue("0.1", 4))
	fmt.Println(ToValue("0.1234", 4))
	fmt.Println(ToValue("0.12345678", 4))
	fmt.Println(ToValue("1", 4))
	fmt.Println(ToValue("1.0", 4))
	fmt.Println(ToValue("1.1", 4))
	fmt.Println(ToValue("1.1234", 4))
	fmt.Println(ToValue("1.12345678", 4))
	fmt.Println(ToValue("1234.1", 4))
	fmt.Println(ToValue("1234.1234", 4))
	fmt.Println(ToValue("1234.12345678", 4))
	fmt.Println(strconv.ParseFloat("01234.5678", 64))

}

func Test_ParamToStringWithSort(t *testing.T) {
	// m := make(map[string]string)
	// m["hello"] = "echo hello"
	// m["world"] = "echo world"
	// m["go"] = "echo go"
	// m["is"] = "echo is"
	// m["cool"] = "echo cool"
	// reply := ParamToStringWithSort(m)
	// log.Printf("Reply: %s", reply)
	//fmt.Println(Guid())
	fmt.Println(EtherToHex("10000000"))
	fmt.Println(EtherToHex("1000.00001"))
	fmt.Println(EtherToHex("1"))
	fmt.Println(EtherToHex("0.1"))
	fmt.Println(EtherToHex("0.0000000000001"))
	fmt.Println(ToValue("10", 18))
	fmt.Println(ToValue("1.0", 18))
	fmt.Println(ToValue("0.01", 18))
	fmt.Println(ToBalance("10000000000000000000", 18)) //10eth
	fmt.Println(ToBalance("10100000000000000000", 18)) //10.1eth
	fmt.Println(ToBalance("10000000000000000", 18))    //0.011eth
}

// var isAddress = function (address) {
//     if (!/^(0x)?[0-9a-f]{40}$/i.test(address)) {
//         // check if it has the basic requirements of an address
//         return false;
//     } else if (/^(0x)?[0-9a-f]{40}$/.test(address) || /^(0x)?[0-9A-F]{40}$/.test(address)) {
//         // If it's all small caps or all all caps, return true
//         return true;
//     } else {
//         // Otherwise check each case
//         return isChecksumAddress(address);
//     }
// };

// var isChecksumAddress = function (address) {
//     // Check each case
//     address = address.replace('0x','');
//     var addressHash = sha3(address.toLowerCase());

//     for (var i = 0; i < 40; i++ ) {
//         // the nth letter should be uppercase if the nth digit of casemap is 1
//         if ((parseInt(addressHash[i], 16) > 7 && address[i].toUpperCase() !== address[i]) || (parseInt(addressHash[i], 16) <= 7 && address[i].toLowerCase() !== address[i])) {
//             return false;
//         }
//     }
//     return true;
// };

func Test_IsAddress(t *testing.T) {
	addr := "0x1a30cb37962a736b23c97a888140cca42d4846eb"
	t.Error(isAddress(addr))
}

func isAddress(addr string) bool {
	reg := regexp.MustCompile("/^(0x)?[0-9a-f]{40}$/")
	if reg.FindString(addr) != "" {
		return false
	}
	return true
}
