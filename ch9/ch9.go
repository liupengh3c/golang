package ch9

import (
	"fmt"
	"golang/ch9/memo5"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Ch9 第九章入口测试函数
func Ch9() {
	// withdraw()
	// ch9Memo()
	ch9Max()
}

// 如果是两个核的话，0和1的打印次数差不多
func ch9Max() {
	go fmt.Print(0)
	fmt.Print(1)
}

type money struct {
	cnt int32
	ok  bool
}

var deposits = make(chan money)
var balances = make(chan int32)
var hua = make(chan money)

func deposit(amount int32) {
	var mon money
	mon.cnt = amount
	deposits <- mon
}

func balance() int32 {
	cur := <-balances
	return cur
}
func teller() {
	var balance int32 = 200
	for {
		select {
		case dep := <-deposits:
			balance += dep.cnt
		case balances <- balance:
		case cnt := <-hua:
			var draw money
			if cnt.cnt > balance {
				draw.cnt = balance
				draw.ok = false
			} else {
				balance = balance - cnt.cnt
				draw.cnt = balance
				draw.ok = true
			}
			hua <- draw
		}
	}
}

func init() {
	go teller()
}
func withdraw() {
	var consume money
	consume.cnt = 300
	hua <- consume

	cur := <-hua
	if cur.ok {
		fmt.Println("money is enough,the balance=" + fmt.Sprintf("%v", cur.cnt))
	} else {
		fmt.Println("money is not enough,the balance=" + fmt.Sprintf("%v", cur.cnt) + fmt.Sprintf(",draw money=%v", 300))
	}
}
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 所有url的请求串行执行.800ms左右完成
// func ch9Memo() {
// 	m := memo1.New(httpGetBody)
// 	urls := []string{
// 		"https://golang.google.cn/pkg/",
// 		"https://golang.google.cn/doc/",
// 		"https://golang.google.cn/pkg/archive/",
// 		"https://golang.google.cn/pkg/",
// 		"https://golang.google.cn/doc/",
// 		"https://golang.google.cn/pkg/archive/",
// 	}
// 	st := time.Now()
// 	for _, url := range urls {
// 		start := time.Now()
// 		value, err := m.Get(url)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
// 	}
// 	fmt.Printf("time cost all: %v", time.Since(st))
// }

// 并行执行版本1
func ch9Memo() {
	m := memo5.New(httpGetBody)
	var wg sync.WaitGroup
	urls := []string{
		"https://golang.google.cn/pkg/",
		"https://golang.google.cn/doc/",
		"https://golang.google.cn/pkg/archive/",
		"https://golang.google.cn/pkg/",
		"https://golang.google.cn/doc/",
		"https://golang.google.cn/pkg/archive/",
	}
	st := time.Now()
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s, %s, %d bytes %s\n", url, time.Since(start), len(value.([]byte)), start)
			defer wg.Done()
		}(url)
	}
	wg.Wait()
	fmt.Printf("time cost all: %v", time.Since(st))
}
