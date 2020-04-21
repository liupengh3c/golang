package ch4

import (
	"crypto/sha256"
	"fmt"
)

// Ch4 第四章测试函数
func Ch4() {
	s := [2]int{4, 8}
	sha()
	// 数组指针测试
	arrayP(&s)
	popcount(14)
}

// Sha 加密
func sha() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
}

// arrayP 数组指针
func arrayP(p *[2]int) {
	for i := range p {
		p[i] = 32
	}
	fmt.Println(p)
}

func popcount(x int64) {
	var pc [256]byte
	cnt := 0
	for k := range pc {
		pc[k] = pc[k/2] + byte(k&1)
		fmt.Printf("%d-%d\n", k, pc[k])
	}
	cnt = int(pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] + pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] + pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] + pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])
	fmt.Println(cnt)
}
