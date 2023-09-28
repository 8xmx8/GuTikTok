package checks

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateInput 输入内容验证   2023-9-28 废除此方案
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

// ValidatePassword 密码校验
func ValidatePassword(password string, min, max int) bool {
	// 验证密码长度是否大于8
	if len(password) < min && len(password) > max {
		return false
	}

	// 验证密码是否包含至少两种类型的字符（英文字母、数字、特殊字符）
	typesCount := 0

	// 包含英文字母
	if containsLetter(password) {
		typesCount++
	}

	// 包含数字
	if containsDigit(password) {
		typesCount++
	}

	// 包含特殊字符
	if containsSpecialChar(password) {
		typesCount++
	}

	return typesCount >= 2
}

func containsLetter(password string) bool {
	// 使用正则表达式验证密码是否包含英文字母
	pattern := `[a-zA-Z]`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(password)
}

func containsDigit(password string) bool {
	// 使用正则表达式验证密码是否包含数字
	pattern := `[0-9]`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(password)
}

func containsSpecialChar(password string) bool {
	// 使用字符串包含函数验证密码是否包含特殊字符
	specialChars := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "{", "}", "[", "]", "<", ">", "?", "~"}
	for _, char := range specialChars {
		if strings.Contains(password, char) {
			return true
		}
	}
	return false
}
