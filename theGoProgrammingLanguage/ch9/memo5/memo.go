// 9.7 示例 并发的非阻塞缓存
// 用通信建立并发程序
package main

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{make(chan request)}
	// monitor goroutine
	go memo.server(f)
	return memo
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key: key, response: response}
	res := <-response
	return res.value, res.err
}
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}
func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

// 把map变量限制在一个单独的monitor goroutine
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)

		}
		go e.deliver(req.response)
	}
}
