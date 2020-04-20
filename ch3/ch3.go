package ch3

import (
	"bytes"
	"fmt"
	"strconv"
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

// Strings 字符串函数
func Strings(s, sub string) {
	if strings.Contains(s, sub) {
		fmt.Println(sub + " is in " + s)
	} else {
		fmt.Println(sub + " is not in " + s)
	}
	fmt.Println(s + " have " + fmt.Sprintf("%d", strings.Count(s, sub)) + " " + sub)
	if strings.HasPrefix(s, sub) {
		fmt.Println(s + " is begin with " + sub)
	} else {
		fmt.Println(s + " is not begin with " + sub)
	}

	if strings.HasSuffix(s, sub) {
		fmt.Println(s + " is end with " + sub)
	} else {
		fmt.Println(s + " is not end with " + sub)
	}

	fmt.Println("index is " + fmt.Sprintf("%d", strings.Index(s, sub)))

	fmt.Println([]string{"i", "am", "student"}, "&")
}

// Bytes 函数
func Bytes(s, sub []byte) {
	if bytes.Contains(s, sub) {
		fmt.Println(string(sub) + " is in " + string(s))
	}
	fmt.Println(string(s) + " have " + fmt.Sprintf("%d", bytes.Count(s, sub)) + " " + string(sub))
}

// IntsToString 将数组以字符串形式打印出来
func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for k, val := range values {
		if k > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", val)
	}
	buf.WriteByte(']')
	return buf.String()
}

// Translate 字符串、数字转换
func Translate() {
	s := "123"
	t, _ := strconv.Atoi(s)
	a := strconv.Itoa(t + 1)
	fmt.Println(a)
}

// Const 常量
func Const() {
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(b)
	const (
		s = 1 << iota
		t
		r
		u
	)
	v := 15
	w := 1
	fmt.Println(fmt.Sprintf("%b", u))
	fmt.Println(w &^ v)
}
