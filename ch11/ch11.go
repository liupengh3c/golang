package ch11

import (
	"fmt"
	"unicode"
)

// Ch11 第十一章入口函数
func Ch11() {
	fmt.Println(IsPalindrome("abcba"))
}

// IsPalindrome 是否回文字符串
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
