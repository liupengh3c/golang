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
	appendInt()
	noempty()
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
	// fmt.Printf("%T %[1]d", pc)
}

func appendInt() {
	var x []int
	fmt.Println(len(x), cap(x))
	for i := 0; i < 20; i++ {
		x = append(x, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(x), x)
	}
}

func noempty() {
	s := []string{"ab", "", "art"}
	// fmt.Println(nonempty(s))
	// fmt.Printf("%q\n", nonempty(s))
	fmt.Println("nonempt2 function test,appen func")
	fmt.Printf("%q\n", nonempty2(s))
}
func nonempty(str []string) []string {
	i := 0
	for _, v := range str {
		if v != "" {
			str[i] = v
			i++
		}
	}
	return str[:i]
}

func nonempty2(str []string) []string {
	var s []string
	for _, v := range str {
		if v != "" {
			fmt.Println(v)
			s = append(s, v)
		}
	}
	fmt.Println(s)
	return s
}

// 二叉树实现实现插入排序
type tree struct {
	value int
	left  *tree
	right *tree
}

// add 将value插入二叉树中，左小右大
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendVal(values []int, t *tree) []int {

	return values
}
