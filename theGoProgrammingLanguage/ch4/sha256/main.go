package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var sha = flag.String("sha", "256", "default use sha256 to encode, you also can use 512 or 384")

func encode(data []byte) (res string) {
	if *sha == "256" {
		c1 := sha256.Sum256(data)
		// fmt.Printf("%s: %x\n", *s, c1)
		res = fmt.Sprintf("%s: %x\n", *sha, c1)
	} else if *sha == "512" {
		c2 := sha512.Sum512(data)
		// fmt.Printf("%s: %x\n", *s, c2)
		res = fmt.Sprintf("%s: %x\n", *sha, c2)
	} else if *sha == "384" {
		c3 := sha512.Sum384(data)
		// fmt.Printf("%s: %x\n", *s, c3)
		res = fmt.Sprintf("%s: %x\n", *sha, c3)
	} else {
		// fmt.Println("code format error")
		// os.Exit(1)
		log.Fatal("code format err, please use sha256,sha512 or sha384")
	}
	return
}

func main() {
	// c1 := sha256.Sum256([]byte{'x', 'a', 'n'})
	// c2 :=
	// 	fmt.Printf("% x\n%v\n", c1, c1)
	flag.Parse()
	// if *s == "sha256" {
	// 	sha256.Sum256()
	// }
	for {
		input := bufio.NewReader(os.Stdin)
		// str, _ := input.ReadString('\n')
		data, _ := input.ReadBytes('\n')
		// encode(data)
		res := encode(data)
		// fmt.Println(res)
		if len(res) == 0 {
			continue
		} else {
			fmt.Printf("%s", res)
		}
	}
}
