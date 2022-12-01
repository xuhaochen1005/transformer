package utils

import (
	"strconv"
)

func FloatToString(a float32) string {
	var strings string
	strings = strconv.FormatFloat(float64(a), 'f', 4, 64)
	return strings
}
