package utils

import (
	"fmt"
	"strconv"
)

func ConcatPrimitiveValues(values ...interface{}) string {
	var result string
	for i, v := range values {
		if i > 0 {
			result += " "
		}
		switch val := v.(type) {
		case int:
			result += strconv.Itoa(val)
		case float32, float64:
			result += strconv.FormatFloat(val.(float64), 'f', -1, 64)
		case string:
			result += val
		case bool:
			result += strconv.FormatBool(val)
		default:
			result += fmt.Sprint(val)
		}
	}
	return result
}

func Print(values ...interface{}) {
	fmt.Println(ConcatPrimitiveValues(values...))
}
