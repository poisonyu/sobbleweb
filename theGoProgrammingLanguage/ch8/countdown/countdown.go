// 一个没有任何case的select语句写作select{}，会永远地等待下去
// 如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	ticker := time.NewTicker(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-ticker.C:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("stop launch")
			ticker.Stop()
			return
		}
	}
	ticker.Stop()
	fmt.Println("start launch")
}
