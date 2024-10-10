package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	ch := make(chan int)
	go func(input *bufio.Scanner) {
		for input.Scan() {
			ch <- 1
		}
		c.Close()
	}(input)
	for {
		select {
		case <-time.After(60 * time.Second):
			fmt.Fprintln(c, "timeout, disconnect...")
			c.Close()
			wg.Wait()
			return
		case <-ch:
			wg.Add(1)
			go echo(c, input.Text(), 1*time.Second, &wg)
		}
	}

}

func main() {
	address := "localhost:8000"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listen in %s\n", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

	// current, _ := os.Getwd()
	// path, err := filepath.Abs("ch8")
	// fmt.Println(current, path, err)
}
