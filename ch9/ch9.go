package ch9

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Ch9 第九章入口测试函数
func Ch9() {
	// withdraw()
	// ch9Memo()
	ch9Memo1()
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

// 缓存示例

// Func 函数类型
type Func func(key string) (interface{}, error)

//
type result struct {
	value interface{}
	err   error
}

// Memo 缓存结构
type Memo struct {
	f     Func
	cache map[string]result
}

func new(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

func (memo *Memo) get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 所有url的请求串行执行
func ch9Memo() {
	m := new(httpGetBody)
	urls := []string{
		"https://golang.google.cn/pkg/",
		"https://golang.google.cn/doc/",
		"https://golang.google.cn/pkg/archive/",
		"https://golang.google.cn/pkg/",
		"https://golang.google.cn/doc/",
		"https://golang.google.cn/pkg/archive/",
	}
	for _, url := range urls {
		start := time.Now()
		value, err := m.get(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

// 并行执行版本1
func ch9Memo1() {
	m := new(httpGetBody)
	var wg sync.WaitGroup
	urls := []string{
		"https://golang.google.cn/pkg/",
		"https://golang.google.cn/doc/",
		"https://golang.google.cn/pkg/archive/",
		"https://golang.google.cn/pkg/",
		"https://golang.google.cn/doc/",
		"https://golang.google.cn/pkg/archive/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.get(url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}
