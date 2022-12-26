package strings

import (
	"encoding/json"
)

func JsonIndent(bytesJson []byte) string {
	m := map[string]interface{}{}
	
	_ = json.Unmarshal(bytesJson, &m)
	
	bytesData, _ := json.MarshalIndent(m, "", "    ")
	
	return string(bytesData)
}
