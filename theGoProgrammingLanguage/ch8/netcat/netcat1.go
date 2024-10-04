package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func mustCopy(dst io.Writer, src io.Reader) {

	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func test(c net.Conn, wg *sync.WaitGroup) {
	var b []byte
	c.Read(b)
	fmt.Printf("b: %s\n", b)
	if string(b) == "close\n" {
		c.Close()
	} else {
		io.Copy(os.Stdout, c)
	}
	wg.Done()

}
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go test(conn, &wg)
	mustCopy(conn, os.Stdin)
	wg.Wait()
	// io.WriteString(conn, cmd)
	// if strings.HasPrefix(cmd, "get ") {
	// 	file := strings.Trim(cmd, "get ")
	// 	if file == "" {
	// 		mustCopy(os.Stdout, conn)
	// 		continue
	// 	}
	// 	file = filepath.Base(file)
	// 	f, _ := os.Create(file)
	// 	mustCopy(f, conn)
	// 	f.Close()
	// } else if strings.HasPrefix(cmd, "send ") {
	// 	// file := strings.Trim(cmd, "send ")
	// 	// if file == "" {
	// 	// 	mustCopy(os.Stdout, conn)
	// 	// 	continue
	// 	// }
	// 	continue
	// } else {
	// 	mustCopy(os.Stdout, conn)
	// }

}
