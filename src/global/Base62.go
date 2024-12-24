package global

import (
	"strings"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//type Base62 struct {
//}

func EncoderBase62(num uint64) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		result = string(charset[num%62]) + result
		num /= 62
	}
	return result
}

func DecoderBase62(str string) uint64 {
	var num uint64
	for _, char := range str {
		index := int64(strings.IndexRune(charset, char))
		if index == -1 {
			Log.Error("Invalid character '%c' in Base62 string", char)
		}
		num = num*62 + uint64(index)
	}
	return num
}
