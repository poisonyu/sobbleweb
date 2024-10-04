package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8000", "use specific port")

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()
	address := "localhost:" + *port
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
