package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

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
			fmt.Println(err)
			continue
		}
		fmt.Printf("connect to remote address %s\n", conn.RemoteAddr().String())
		go handleConn(conn)

	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		cmd(c, input.Text())
	}

}

func cmd(c net.Conn, s string) {
	if s == "ls" {
		dirs, _ := os.ReadDir(".")
		var directory string
		for _, dir := range dirs {
			directory = dir.Name() + " "
		}
		io.WriteString(c, directory+"\n")
	} else if strings.HasPrefix("cd ", s) {
		path := strings.Trim(s, "cd ")
		current, _ := os.Getwd()
		if !filepath.IsAbs(path) {
			fmt.Printf("current directory %s\n", current)
			path = filepath.Join(current, path)
		} else if path == ".." {
			path = filepath.Dir(current)
		}
		fmt.Println("change directory to ", path)
		err := os.Chdir(path)
		if err != nil {
			io.WriteString(c, err.Error())
		}
		io.WriteString(c, fmt.Sprintf("change directory to %s successfully...\n", path))
	} else if s == "close" {
		io.WriteString(c, "close")
		c.Close()
		// } else if strings.HasPrefix(s, "get ") {
		// 	file := strings.Trim(s, "get ")
		// 	if file == "" {
		// 		io.WriteString(c, "need file name")
		// 	}
		// 	if !filepath.IsAbs(file) {
		// 		current, _ := os.Getwd()
		// 		file = filepath.Join(current, file)
		// 	}

		// 	// os.ReadFile(file)
		// 	f, err := os.Open(file)
		// 	if err != nil {
		// 		io.WriteString(c, err.Error())
		// 	}
		// 	_, err = io.Copy(c, f)
		// 	if err != nil {
		// 		io.WriteString(c, err.Error())
		// 	}
		// 	f.Close()
	} else {
		fmt.Fprintln(c, "command error")
	}
}
func NewChanString(conn net.Conn) chan string {
	c := make(chan string)
	go func() {
		for {
			var b []byte
			n, err := conn.Read(b)
			if err != nil {
				fmt.Println(n, err)
				continue
			}
			c <- string(b)
		}
	}()
	return c
}
