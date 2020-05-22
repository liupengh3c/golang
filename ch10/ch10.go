package ch10

import (
	"bytes"
	"crypto/rand"
	"fmt"
	mrand "math/rand"
	"runtime"
	"time"
)

// 两个名字一样的包，导入到第三个包中时，导入
// 声明就必须至少为其中一个指定一个替代名字

// Ch10 入口函数
func Ch10() {
	mathRand()
	chEnv()
}

func mathRand() {
	fmt.Println(mrand.Int31n(100))
}
func crypRand() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(bytes.Equal(b, make([]byte, c)))
}

func chEnv() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	time.Now()
}
