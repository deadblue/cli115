package util

import "encoding/json"

func MustMarshal(value interface{}) []byte {
	if data, err := json.Marshal(value); err != nil {
		return []byte{}
	} else {
		return data
	}
}
