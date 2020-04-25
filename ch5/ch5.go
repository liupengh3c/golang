package ch5

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	getHTML()
}

func visit(links []string, n *html.Node) []string {
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

// fetch 简易版爬虫
func fetch(url string) *os.File {
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
	io.Copy(os.Stdout, resp.Body)
	return os.Stdout
}

func getHTML() {
	doc, err := html.Parse(fetch("http://golang.google.cn/"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "finlink1:%v\n", err)
	}
	fmt.Println(doc.Data)
	// for _, val := range visit(nil, doc) {
	// 	fmt.Println(val)
	// }
}
