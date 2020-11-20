package cstools

import "encoding/json"

func HasKey(m map[string]interface{}, key string) bool {
	if _, ok := m[key]; ok {
		return true
	} else {
		return false
	}
}

func MapToJsonStr(m interface{}) string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	s := string(b)
	return s
}

func MapToStruct(m interface{}, target interface{}) {
	s := MapToJsonStr(m)
	json.Unmarshal([]byte(s), &target)
}
