package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch() {
	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		resp.Status = "302"
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// b, err := io.ReadAll(resp.Body)
		f, _ := os.Create("index.html")
		io.Copy(f, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		// fmt.Printf("%s\n", b)
	}
}

func main() {
	Fetch()
}
