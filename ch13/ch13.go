package ch13

import (
	"fmt"
	"unsafe"
)

// Ch13 第13章入口函数
func Ch13() {
	fmt.Println(unsafe.Sizeof(float64(7)))
	first()
}

type node struct {
	b   bool
	age float64
	num int32
}

type node1 struct {
	age float64
	num int32
	b   bool
}

// 13.1 示例
func first() {
	var n node
	fmt.Println("n:", unsafe.Sizeof(n))
	fmt.Println("n.b offset:", unsafe.Offsetof(n.b))
	fmt.Println("n.age offset:", unsafe.Offsetof(n.age))
	fmt.Println("n.num offset:", unsafe.Offsetof(n.num))

	var n1 node1
	fmt.Println("n1:", unsafe.Sizeof(n1))
	fmt.Println("n1.age offset:", unsafe.Offsetof(n1.age))
	fmt.Println("n1.num offset:", unsafe.Offsetof(n1.num))
	fmt.Println("n1.b offset:", unsafe.Offsetof(n1.b))
}
