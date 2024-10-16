package main

import (
	"fmt"
	"sync"
	"time"
)

var loadpcOnce sync.Once
var pc [256]byte

// func init() {
// 	for i := range pc {
// 		pc[i] = pc[i/2] + byte(i&1)
// 	}
// }

func loadpc() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Popcount(x uint64) int {
	loadpcOnce.Do(loadpc)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func Popcountrange(x uint64) int {
	var b byte
	for i := 0; i < 8; i++ {
		b += pc[byte(x>>(i*8))]
	}

	return int(b)
}

func Test(f func(x uint64) int, x uint64) {
	var total float64
	for i := 1; i < 11; i++ {
		start := time.Now()
		res := f(x)
		fmt.Printf("%d. popcount %d\n", i, res)
		elapsed := time.Since(start)
		fmt.Println("elapsed time ", elapsed)
		total += elapsed.Seconds()
	}
	fmt.Println("average elapsed time ", total/10.0)
}
func main() {
	// start := time.Now()
	// fmt.Println("1 -> ", Popcount(300))
	// fmt.Println("elapsed -> ", time.Since(start))
	fmt.Println("Popcount:")
	Test(Popcount, 30000)
	fmt.Printf("********\nPopcountrange\n")
	Test(Popcountrange, 30000)

}

//	func WaitForServer(url string) error {
//		const timeout = 1 * time.Minute
//		deadline := time.Now().Add(timeout)
//		for tries := 0; time.Now().Before(deadline); tries++ {
//			_, err := http.Head(url)
//			if err == nil {
//				return nil
//			}
//			log.Printf("server not responding (%s);retrying...", err)
//			time.Sleep(time.Second << uint(tries))
//		}
//		return fmt.Errorf("server %s failed to respond after %s", url, timeout)
//	}
// func ffsf() error {
// 	// ioutil.TempDir("", "scratch")
// 	dir, err := os.MkdirTemp("", "scratch")
// 	if err != nil {
// 		return fmt.Errorf("failed to create temp dir: %v", err)
// 	}
// 	os.Getwd()
// 	os.ReadDir()
// 	os.ReadFile()
// 	os.RemoveAll(dir)
// }

// func sgsg() error {
// 	in := bufio.NewReader(os.Stdin)
// 	for {
// 		r, _, err := in.ReadRune()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return fmt.Errorf("read failed:%v", err)
// 		}
// 		// ...use r...
// 	}
// }
