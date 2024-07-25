package utils

import (
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
