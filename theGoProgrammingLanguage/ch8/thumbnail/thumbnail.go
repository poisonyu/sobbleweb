package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func makeThumbnails(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		// if _, err := ImageFile(f); err != nil {
		// 	log.Println(err)
		// }
		// go ImageFile(f)
		go func(filename string) {
			ImageFile(filename)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

func makeThumbnails2(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		// if _, err := ImageFile(f); err != nil {
		// 	log.Println(err)
		// }
		// go ImageFile(f)
		go func(filename string) {
			_, err := ImageFile(filename)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

type item struct {
	thumbfile string
	err       error
}

func makeThumbnails3(filenames []string) ([]string, error) {
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		// if _, err := ImageFile(f); err != nil {
		// 	log.Println(err)
		// }
		// go ImageFile(f)
		go func(filename string) {
			var it item
			it.thumbfile, it.err = ImageFile(filename)
			ch <- it
			// thumbfile, err := ImageFile(filename)
			// ch <- item{thumbfile: thumbfile, err: err}
		}(f)
	}
	var thumbfiles []string
	for range filenames {
		it := <-ch
		if it.err != nil {
			log.Println(it.err)
			continue
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil

}
func makeThumbnails4(in <-chan string) (total int64) {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for filename := range in {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(filename)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(filename)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()
	for size := range sizes {
		total += size
	}
	return
}
func ImageFile(infile string) (string, error) {
	fmt.Println(infile, "filename ")
	return infile, nil
}
