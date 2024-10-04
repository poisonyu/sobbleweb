// package main

// import (
// 	"flag"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"time"
// )

// var port = flag.String("port", "8000", "use specific port")

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
// 		if err != nil {
// 			return
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	flag.Parse()
// 	address := "localhost:" + *port
// 	listener, err := net.Listen("tcp", address)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Listen in %s\n", address)
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Print(err)
// 			continue
// 		}
// 		go handleConn(conn)
// 	}
// 	// current, _ := os.Getwd()
// 	// path, err := filepath.Abs("ch8")
// 	// fmt.Println(current, path, err)
// }

package bank

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}
func init() {
	go teller()
}

type Cake struct {
	state string
}

var cooked = make(chan *Cake)
var iced = make(chan *Cake)

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake
	}
}
