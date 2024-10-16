// 9.7 示例 并发的非阻塞缓存
// 用锁建立并发程序
package memo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
func incomingURLs() chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- strconv.Itoa(i)
		}
	}()
	return ch
}

func main() {
	var wg sync.WaitGroup

	m := New(httpGetBody)
	for url := range incomingURLs() {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

// go test -run=TestConcurrent -race -v gopl.io/ch9/memo1

// 这里使用了一个互斥量来保护多个goroutine调用Get时的共享map变量
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready

	}
	return e.res.value, e.res.err
}
