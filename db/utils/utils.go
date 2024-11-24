package utils

import (
	"encoding/binary"
	"strings"
	"unicode"
)

func ClearStr(str string) string {
	ret := ""
	str = strings.ToLower(strings.TrimSpace(str))
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			ret += string(r)
		}
	}
	return ret
}

func ClearStrSpace(str string) string {
	ret := ""
	str = strings.ToLower(strings.TrimSpace(str))
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			ret += string(r)
		} else if len(ret) > 0 && ret[len(ret)-1] != ' ' {
			ret += " "
		}
	}

	return ret
}

func I2B(num int64) []byte {
	value := make([]byte, 8)
	binary.BigEndian.PutUint64(value, uint64(num))
	return value
}

func B2I(value []byte) int64 {
	return int64(binary.BigEndian.Uint64(value))
}
