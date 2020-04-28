package ch6

import (
	"bytes"
	"fmt"
	"math"
)

// 2020-04-27 08:56

// Ch6Enter 第六章入口函数
func Ch6() {
	ch6List()
	ch6Vector()
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

type intset struct {
	words []uint64
}

func (s *intset) add(word int) {
	index, bit := word/64, word%64
	for index >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[index] |= 1 << bit
	return
}
func (s *intset) has(word int) bool {
	index, bit := word/64, word%64
	return index <= len(s.words) && s.words[index]&(1<<bit) != 0
}

func (s *intset) string() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word>>j&0x1 == 1 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*k+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func ch6Vector() {
	var word = new(intset)
	word.add(1)
	word.add(2)
	word.add(3)
	// slice中只有一个值，那就是0xe
	fmt.Println(word.words)
	fmt.Println(word.string())
}
