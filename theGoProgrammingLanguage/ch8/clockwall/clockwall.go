package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// var Addr map[string]string

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		a := strings.Split(arg, "=")
		go Connect(a[0], a[1])
	}
	time.Sleep(5 * time.Minute)
}

func Connect(city, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("connect to %s in %s successfully...\n", addr, city)
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
