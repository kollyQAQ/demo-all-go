package string

import (
	"regexp"
	"unicode"
)

// 判断是否是中文
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// regex 判断 text 包含字符
func Reg(regex, text string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(text)
}
