# the go programming language

## 8.4 Channels
* goroutine 并发体
* channels 通信机制

在发送语句中 '<-'分割channel和要发送的值
chan<- int表示一个只发送int的channel
>因为关闭操作只用于断言不再向channel发送新的数据，所以只有在发送者所在的goroutine才会调用close函数

在接受语句中，'<-'写在channel对象之前
<-chan int表示一个只接收int的channel

close用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常
对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据；如果channel中已经没有数据的话将产生一个零值的数据

1. 无缓存channels(同步channels)
无缓存channels的发送和接收操作将导致两个goroutine做一次同步操作。
> 当我们说x事件既不是在y事件之前发生也不是在y事件之后发生，我们就说x事件和y事件是并发的。这并不是意味着x事件和y事件就一定是同时发生的，我们只是不能确定这两个事件发生的先后顺序。当两个goroutine并发访问了相同的变量时，我们有必要保证某些事件的执行顺序，以避免出现某些并发问题

channels发送消息
* 每个消息都有一个值
* 通讯的事实和发生的时刻(消息事件)

2. 带缓存的channels(buffered channel)


3. 串联的Channels(Pipeline)
```
go func() {
    for {
        x, ok := <-naturals
        if !ok {
            break
        }
        squares <- x * x
    }
    close(squares)
}()
```
>其实你并不需要关闭每一个channel。只有当需要告诉接收者goroutine，所有的数据已经全部发送时才需要关闭channel。不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收。

## 8.5 并发的循环

```
<!-- func makeThumbnails2(filenames []string) {
    res := make(chan string)
    for _, f := range filenames {
        go func(infile string) {
            s, _ := thumbnail.ImageFile(infile)
            res <- s
        }
    }
    for range res {}
} -->
// 等待goroutine完成后退出
func makeThumbnails2(filenames []string) {
    ch := make(chan struct{})
    for _, f := range filenames {
        go func(f string) {
            thumbnail.ImageFile(f)
            ch <- struct{}{}
        }(f)
    }
    for range filenames {
        <-ch
    }
}

func makeThumbnails2(filenames []string) error {
    // errors := make(chan error, len(filenames))
    errors := make(chan error) 
    for _, f := range filenames {
        go func(f string) {
            _, err := thumbnail.ImageFile(f)
            errors <- err
        }(f)
    }
    for range filenames {
        if err := <-errors; err != nil {
            // 返回err可能会产生goroutine泄漏，需要给通道设置合适大小的缓存或者创建一个goroutine去排空channel
            //go func() {
            //    for range errors {}
            //}()
            return err
        }
    }
    return nil 
}

func makeThumbnails2(filenames []string) (thumbfiles []stirng, err error) {
    type item struct {
        thumbfile string
        err error
    }
    ch := make(chan item, len(filenames))
    for _, f := range filenames {
        go func(f string) {
            var it item
            it.thumbfile, it.err := thumbnail.ImageFile(f)
            ch <- it
        }(f)
    }
    for range filenames {
        it := <-ch
        if it.err != nil {
            return thumbfiles, it.err
        }
        thumbfiles = append(thumbfiles, it.thumbfile)
    }
    return thumbfiles, nil 
}

```
sync.WaitGroup
thumbnail/thumbnail.go

## 8.6 
```
func crawl(url string) []string {
    fmt.Println(url)
    list, err := links.Exttract(url)
    if err != nil {
        log.Print(err)
    }
    return list
}

func main() {
    worklist := make(chan []string)
    go func() {
        worklist <-os.Args()[1:]
    }()
    seen := make(map[string]bool)
    // 在main goroutine中的worklist是接收channel
    for list := range worklist {
        for _, link := range list {
            if !seen[link] {
                seen[link] = true
                go func(link string) {
                    // 在crawl goroutine中worklist是发生channel
                    worklist <- crawl(link)
                }(link)
            }
        }
    }
}

// 使用buffered channel控制并发
// 解决程序不会终止的问题
func main() {
    worklist := make(chan []string)
    var n int 
    n++
    go func() {
        worklist <-os.Args()[1:]
    }()
    seen := make(map[string]bool)
    sema := make(chan struct{}, 20)
    // 在main goroutine中的worklist是接收channel
    for ;n>0;n-- {
        list := <- worklist
        for _, link := range list {
            if !seen[link] {
                seen[link] = true
                n++ 
                go func(link string) {
                    sema <- struct{}{}
                    // 在crawl goroutine中worklist是发生channel
                    worklist <- crawl(link)
                    <-sema
                }(link)
            }
        }
    }
}

// 另一种控制并发的思路
func main() {
    worklist := make(chan []string)
    unseenLinks := make(chan string)

    for i:=0; i<20; i++ {
        go func() {
            for link := range unseenLinks {
                foundLinks := crawl(link)
                go func() {
                    worklist <- foundLinks
                }()
            }
        }()
    }
    seen := make(map[string]bool)
    // 这里会让程序永远都不终止吗？
    for list := range worklist {
        for _, link := list {
            if !seen[link] {
                seen[link] = true
                unseenLinks <- link
            }
        }
    }
}
```
