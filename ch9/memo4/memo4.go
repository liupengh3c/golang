package memo4

import "sync"

// Func 函数类型
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//!+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// New 初始化实例
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Memo 缓存结构
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

// Get 发送请求
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // 广播数据准备完毕
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // 等待数据准备完毕
	}
	return e.res.value, e.res.err
}
