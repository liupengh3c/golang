package ch1

import (
	"fmt"
	"os"
)

// GetStdCmd 第一章示例代码，第1个
func GetStdCmd() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	return
}

// GetStdCmd2 第一章示例代码，第2个
func GetStdCmd2() {
	s, sep := "", ""
	for _, val := range os.Args[1:] {
		s += sep + val
		sep = " "
	}
	fmt.Println(s)
	return
}
