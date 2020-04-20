package ch3

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Unicode 测试字符串
func Unicode() {
	str := "人人都是产品经理"
	fmt.Println(str + "-字节数，len=" + fmt.Sprintf("%d", len(str)))

	for i := 0; i < len(str); i++ {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Println(fmt.Sprintf("%d\t%q\t%d", i, r, size))
		i += size
	}
	for _, val := range str {
		fmt.Println(string(val))
	}
	s := []rune("hello,世界")
	fmt.Println(len(s))
}

// Basename 获取文件名
func Basename(str string) {
	var s string
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '/' {
			s = str[i+1:]
			break
		}
	}
	fmt.Println(s)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	fmt.Println(s)
	return
}

// Basename2 简化版，利用库函数
func Basename2(s string) {
	lastIndex := strings.LastIndex(s, "/")
	s = s[lastIndex+1:]
	lastIndex = strings.LastIndex(s, ".")
	s = s[:lastIndex]
	fmt.Println(s)
	return
}

// Comma 递归
func Comma(s string) {
	fmt.Println(comma(s))
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
