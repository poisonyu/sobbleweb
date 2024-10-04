package main

import (
	"fmt"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	// ...
	return []string{"i can fly", "dog", url}
}

func main1() {
	worklist := make(chan []string)

	go func() {
		worklist <- os.Args[1:]
	}()
	seen := make(map[string]bool)
	for list := range worklist {
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				go func(url string) {
					worklist <- crawl(url)
				}(url)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	tokens <- struct{}{}
	var res = []string{"dog", "cat", "elelpent", url}
	// todo crawl
	<-tokens
	return res
}

func main2() {
	worklist := make(chan []string)
	var n int
	n++
	go func() {
		worklist <- os.Args[1:]
	}()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				n++
				go func(url string) {
					worklist <- crawl2(url)
				}(url)
			}
		}
	}
}
func main3() {
	worklist := make(chan []string)
	unseenlist := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for url := range unseenlist {
				foundUrl := crawl2(url)
				go func() {
					worklist <- foundUrl
				}()
			}
		}()
	}
	seen := make(map[string]bool)
	for list := range worklist {
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				go func() {
					unseenlist <- url
				}()
			}
		}
	}
}

const depth = 3

type itemlist struct {
	s     []string
	depth int
}

func crawl4(url string, depth int) itemlist {
	tokens <- struct{}{}
	var res = []string{"dog", "cat", "elelpent", url}
	// todo crawl
	<-tokens
	return itemlist{res, depth + 1}
}
func main4() {
	worklist := make(chan itemlist)
	var n int
	n++
	go func() {
		worklist <- itemlist{os.Args[1:], 0}
	}()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, url := range list.s {
			if !seen[url] && list.depth < depth {
				seen[url] = true
				n++
				go func(url string) {
					worklist <- crawl4(url, list.depth)
				}(url)
			}
		}
	}
}
