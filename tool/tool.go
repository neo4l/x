package tool

import(
	"encoding/json"
)

func ToJson(bean interface{}) []byte {
	bytes, err := json.Marshal(bean)
	if err != nil {
		return nil
	}
	return bytes
}