package ch6

import (
	"fmt"
	"math"
)

// 2020-04-27 08:56

// Ch6Main 第六章入口函数
func Ch6Main() {
	ch6List()
}

type point struct {
	X float64
	Y float64
}

// 普通函数
func distance(p, q point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// point类型的方法,不通类型的method是可以同名的
// 因为每个类型都有自己的命名空间，不冲突
func (p point) distance(q point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// --------------实现一个链表----------------- //
type node struct {
	val  int32
	next *node
}

// 计算列表元素的和，递归
func (n *node) sum() int32 {
	if n == nil {
		return 0
	}
	return n.val + n.next.sum()
}

// 链表头
type list struct {
	head *node
}

// 初始化链表头
func initList() *list {
	head := new(list)
	node := new(node)
	node.val = 0
	node.next = nil
	head.head = node
	return head
}

// 判断链表是否为空，list为链表头指针
func (list *list) empty() bool {
	if list.head == nil {
		return true
	}
	return false
}

// 从链表头加入元素
func (list *list) add(data int32) *node {
	n := new(node)
	n.next = list.head
	n.val = data

	list.head = n
	return n
}

// 从链表尾部加入元素
func (list *list) append(data int32) *node {
	n := new(node)
	n.val = data
	n.next = nil
	if list.empty() {
		list.head = n
	} else {
		cur := list.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = n
	}

	return n
}

func ch6List() {
	list := initList()
	list.head.val = 1
	for i := 1; i < 100; i++ {
		list.append(int32(i + 1))
	}
	node := list.head
	for {
		fmt.Printf("the %dth element:%[1]d\n", node.val)
		if node.next == nil {
			break
		}
		node = node.next
	}
}
