package utils

import (
	"encoding/json"
)

func StringToArray(str string) []float32 {
	var arr []float32
	err := json.Unmarshal([]byte(str), &arr)
	if err != nil {
		return nil
	}
	return arr
}
