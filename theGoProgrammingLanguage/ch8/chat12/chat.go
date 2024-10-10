package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const clientKeepTime = 30 * time.Second

type client chan<- string

type clientchan struct {
	name string
	client
}

var entering = make(chan *clientchan)
var leaving = make(chan *clientchan)
var messages = make(chan string)

func broadcaster() {
	clients := make(map[*clientchan]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				select {
				case cli.client <- msg:
				case <-time.After(30 * time.Second):
					// 这边进入下一个循环还不清楚用break还是continue
					continue
				}

			}
		// 从entering中获取从每个handleConn中的ch,赋值给cli,那么ch和cli就联系在了一起
		case cli := <-entering:
			clients[cli] = true
			var users []string
			for cli := range clients {
				users = append(users, cli.name)
			}
			cli.client <- fmt.Sprintf("%v\n", users)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.client)
		}
	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string, 10)

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who + "\n"
	messages <- who + " has arrived\n"
	cc := clientchan{who, ch}
	entering <- &cc

	ticker := time.NewTicker(clientKeepTime)
	tickerChan := make(chan struct{})
	input := bufio.NewScanner(conn)
	go func() {
		for {
			select {
			case <-tickerChan:
				ticker = time.NewTicker(clientKeepTime)
			case <-ticker.C:
				conn.Close()
			}
		}
	}()
	for input.Scan() {
		tickerChan <- struct{}{}
		messages <- who + ": " + input.Text() + "\n"
	}
	leaving <- &cc
	messages <- who + " has left\n"
	conn.Close()

}
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprint(conn, msg)
	}

}
func main() {
	address := "localhost:8000"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listen in %s\n", address)
	// listener, err := net.Listen("tcp", "localhost:8000")
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
