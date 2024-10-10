package main

import (
	"fmt"
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

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("connect %s successfully\n", conn.RemoteAddr().String())
// 	done := make(chan struct{})
// 	go func() {
// 		io.Copy(os.Stdout, conn)
// 		log.Println("done")
// 		done <- struct{}{}
// 	}()
// 	mustCopy(conn, os.Stdin)
// 	conn.Close() // 关闭读和写方向的网络连接
// 	// 关闭网络连接中的写方向的连接将导致server程序收到一个文件EOF结束信号
// 	// 关闭网络连接中读方向的连接将导致后台goroutine的io.Copy函数调用返回一个"read from closed connection"的错误。

// 	<-done
// }

func main() {
	raddr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("connect %s successfully\n", conn.RemoteAddr().String())
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done
}
