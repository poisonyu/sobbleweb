package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var entering = make(chan client)
var leaving = make(chan client)
var messages = make(chan string)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		// 从entering中获取从每个handleConn中的ch,赋值给cli,那么ch和cli就联系在了一起
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who + "\n"
	messages <- who + "has arrived\n"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text() + "\n"
	}
	leaving <- ch
	messages <- who + " has left\n"
	conn.Close()

}
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprint(conn, msg)
	}

}
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
