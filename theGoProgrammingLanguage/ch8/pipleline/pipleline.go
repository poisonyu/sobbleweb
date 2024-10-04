package main

import (
	"fmt"
)

func pipleline1() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; ; i++ {
			naturals <- i
		}
	}()

	go func() {
		for {
			n, ok := <-naturals
			if !ok {
				break
			}
			squares <- n * n
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
	}
}

func pipleline2() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	for x := range squares {
		fmt.Println(x)
	}
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}
func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
func pipleline3() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
func main() {
	pipleline3()
}
