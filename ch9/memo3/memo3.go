package memo3

import "sync"

// Memo 缓存结构
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

// Func 函数类型
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New 初始化实例
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get 发送请求
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical sections, several goroutines
		// may race to compute f(key) and update the map.
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
