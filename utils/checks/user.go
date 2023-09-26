package checks

import (
	"fmt"
	"regexp"
)

// ValidateInput 输入内容验证
func ValidateInput(minLen, maxLen int, s ...string) string {
	for i := range s {
		n := len(s[i])
		switch {
		case n == 0:
			return "不可以为空值!"
		case n < minLen:
			return fmt.Sprintf("长度太短，最短为%d个字符!", minLen)
		case n > maxLen:
			return fmt.Sprintf("长度太长，最长为%d个字符!", maxLen)
		case !isValidString(s[i]):
			return "不合法,只能使用大小写字母，数字与.!@$%#特殊符号!"
		}
	}
	return ""
}

func isValidString(s string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9\\.!@$%#]+$", s)
	return ok
}
