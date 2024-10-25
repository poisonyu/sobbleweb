package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ch := make(chan struct{})
	req, _ := http.NewRequest("GET", "", nil)
	req.Cancel = ch
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(ch)
	}()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}

func crawl(url string) {
	ch := make(chan struct{})
	timer := time.AfterFunc(5*time.Second, func() { close(ch) })
	req, _ := http.NewRequest("GET", url, nil)
	// ctx, cancel := context.WithCancel(context.Background())
	// req.WithContext(ctx)
	req.Cancel = ch
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("[response error]", err)
	}
	defer resp.Body.Close()
	// status := resp.Status
	// fmt.Println("response status: ", status)
	for {
		timer.Reset(2 * time.Second)
		_, err = io.CopyN(io.Discard, resp.Body, 256)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

}

func CnacelRequest(url string) {
	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	// go func() {
	// 	<-req.Context().Done()

	// }()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancel()
	}()
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

}

func Request(ctx context.Context, url string) string {
	var b []byte
	// ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	// go func() {
	// 	os.Stdin.Read(make([]byte, 1))
	// 	cancel()
	// }()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	resp.Body.Read(b)
	return string(b)

}

func mirroredQuery() string {
	var responses = make(chan string, 3)
	ctx, cancel := context.WithCancel(context.Background())
	for _, url := range os.Args[1:] {
		go func(url string) {
			responses <- Request(ctx, url)
		}(url)
	}
	res := <-responses
	cancel()
	return res

}

// func panicrecover() (v int) {
// 	defer func() {
// 		switch p := recover(); p.(type) {
// 		case int:
// 			v = p.(int)
// 		default:
// 			fmt.Println("...")
// 		}
// 	}()
// 	panic(3)
// }
