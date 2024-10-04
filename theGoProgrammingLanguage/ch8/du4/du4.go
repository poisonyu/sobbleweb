package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var semaphore = make(chan struct{}, 20)
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func walkDir(root string, filesizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(root) {
		if entry.IsDir() {
			subpath := filepath.Join(root, entry.Name())
			wg.Add(1)
			go walkDir(subpath, filesizes, wg)
		} else {
			info, _ := entry.Info()
			filesizes <- info.Size()
		}
	}

}

func dirents(dir string) []fs.DirEntry {
	//fmt.Println(dir)
	select {
	case semaphore <- struct{}{}:
	case <-done:
		return nil
	}
	// semaphore <- struct{}{}
	defer func() {
		<-semaphore
	}()
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Println("[ReadDir]", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles int64, sizes int64) {
	fmt.Printf("%d files %.1fGB\n", nfiles, float64(sizes)/1e6)
}
func main() {
	flag.Parse()
	filesizes := make(chan int64)
	var tick <-chan time.Time
	var wg sync.WaitGroup
	var nfiles int64
	var sizes int64

	if *verbose {
		tick = time.Tick(500 * time.Microsecond)
	}
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, filesizes, &wg)
	}
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	go func() {
		wg.Wait()
		close(filesizes)
	}()

loop:
	for {
		select {
		case <-done:
			for size := range filesizes {
				nfiles++
				sizes += size
			}
			return
		case <-tick:
			printDiskUsage(nfiles, sizes)
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			sizes += size
		}
	}
}
