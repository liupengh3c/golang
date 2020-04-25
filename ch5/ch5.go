package ch5

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

//------------------ 5.2递归-----------------//

// NodeType 节点类型
type NodeType int32

// Attribute 属性
type Attribute struct {
	key   string
	value string
}

// Node 节点属性
type Node struct {
	Type                    NodeType
	data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

// Ch5 第五章函数入口
func Ch5() {
	// getHTML()
	// proErr()
	// getHTML2()
	ch5Sort()
}

// fetch 简易版爬虫
func fetch(url string) *os.File {
	var h = new(os.File)
	s := time.Now()
	fmt.Println("begin request")
	resp, err := http.Get(url)
	t := time.Since(s)
	fmt.Println(t)
	if err != nil {
		fmt.Println("request " + url + " fail")
		return nil
	}
	defer resp.Body.Close()
	// b, err := ioutil.ReadAll(resp.Body)
	io.Copy(h, resp.Body)
	return h
}

// n是一个链表，初始化后是html的最开始位置··16·2·2]

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		fmt.Println(n.Data)
	}
	if n != nil && n.Type == html.ElementNode && n.Data == "a" {
		for _, val := range n.Attr {
			if val.Key == "href" {
				links = append(links, val.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func getHTML() {
	f, _ := os.Open("2.html")
	defer f.Close()
	doc, err := html.Parse(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "finlink1:%v\n", err)
	}

	for _, val := range visit(nil, doc) {
		fmt.Println(val)
	}
}

// 函数变量+错误处理

func add(r rune) rune {
	return r + 1
}
func proErr() error {
	f := add
	fmt.Println(add(2))
	fmt.Println(strings.Map(add, "HAL-9000"))
	fmt.Println(f(3))
	return nil
}

func forEachMode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachMode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth--
	}
}

func getHTML2() {
	f, _ := os.Open("2.html")
	defer f.Close()
	doc, err := html.Parse(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "finlink1:%v\n", err)
	}
	forEachMode(doc, startElement, endElement)
}

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	// 匿名函数，深度优先算法
	visitAll = func(items []string) {
		for key, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				fmt.Println("key:", key, item)
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	// sort.Strings(keys)
	fmt.Println(keys)
	visitAll(keys)
	return order
}

func ch5Sort() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
