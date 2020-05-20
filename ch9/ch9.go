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
	mux   sync.Mutex
	cache map[string]result
}

func new(f Func) *Memo {
	return &Memo{
		f:     f,
		mux:   sync.Mutex{},
		cache: make(map[string]result),
	}
}

// get 加了信号量保护,完全串行执行，性能很差,下面为输出
// command-line-arguments
// https://golang.google.cn/pkg/, 437.9196ms, 57808 bytes
// https://golang.google.cn/doc/, 656.2899ms, 9008 bytes
// https://golang.google.cn/pkg/archive/, 875.4551ms, 5453 bytes
// https://golang.google.cn/pkg/, 875.4551ms, 57808 bytes
// https://golang.google.cn/pkg/archive/, 875.4551ms, 5453 bytes
// https://golang.google.cn/doc/, 875.4551ms, 9008 bytes
// func (memo *Memo) get(key string) (interface{}, error) {
// 	memo.mux.Lock()
// 	res, ok := memo.cache[key]
// 	if !ok {
// 		res.value, res.err = memo.f(key)
// 		memo.cache[key] = res
// 	}
// 	memo.mux.Unlock()
// 	return res.value, res.err
// }

// 第二个版本
// https://golang.google.cn/pkg/archive/, 323.306ms, 5453 bytes
// https://golang.google.cn/pkg/archive/, 323.306ms, 5453 bytes
// https://golang.google.cn/doc/, 324.3431ms, 9008 bytes
// https://golang.google.cn/doc/, 330.2877ms, 9008 bytes
// https://golang.google.cn/pkg/, 333.3073ms, 57808 bytes
// https://golang.google.cn/pkg/, 334.3063ms, 57808 bytes
func (memo *Memo) get(key string) (interface{}, error) {
	memo.mux.Lock()
	res, ok := memo.cache[key]
	memo.mux.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mux.Lock()
		memo.cache[key] = res
		memo.mux.Unlock()
	}
	return res.value, res.err
}

func (memo *Memo1) get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		fmt.Println(key)
		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}

// Memo1 升级版
type Memo1 struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}
type entry struct {
	res   result
	ready chan struct{}
}

func new1(f Func) *Memo1 {
	return &Memo1{
		f:     f,
		mu:    sync.Mutex{},
		cache: make(map[string]*entry),
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
	m := new1(httpGetBody)
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
			defer wg.Done()
		}(url)
	}
	wg.Wait()
}
