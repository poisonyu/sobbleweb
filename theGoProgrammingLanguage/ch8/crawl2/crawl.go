package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	go func() {
		os.Stdin.Read(make([]byte, 1))

	}()

}

func crawl(url string) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("[response error]", err)
	}
	resp
}
