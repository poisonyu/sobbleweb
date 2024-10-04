package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

}

func main() {
	// conn, err := net.Dial("tcp", "localhost:8000")
	raddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()
	// done := make(chan struct{})
	// done := make(chan int)

	// go func() {
	// 	io.Copy(os.Stdout, conn)
	// 	fmt.Println("done")
	// 	done <- struct{}{}
	// 	// done <- 1
	// }()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
	// conn.CloseWrite()
	// <-done
}
