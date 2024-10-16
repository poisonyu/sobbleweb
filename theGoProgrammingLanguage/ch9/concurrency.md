# [the go programming language](https://golang-china.github.io/gopl-zh/index.html title="Go语言圣经")

## 9.1 基于共享变量的并发

### 竞争条件
当我们没有办法自信地确认一个事件是在另一个事件的前面或者后面发生的话，就说明x和y这两个事件是并发的
一个函数在线性程序中可以正确地工作。如果在并发的情况下，这个函数依然可以正确地工作的话，那么我们就说这个函数是并发安全的
对于某个类型来说，如果其所有可访问的方法和操作都是并发安全的话，那么该类型便是并发安全的

为了避免并发访问大多数的类型，可以将变量局限在单一的一个goroutine内，或用互斥条件维持更高级别的不变性。

包级别的导出函数一般情况下都是并发安全的，包级别的变量没法被限制在单一的goroutine中，修改这些变量必须使用互斥条件。

**竞争条件**是指程序在多个goroutine交叉执行操作时，没有给出正确结果。
**数据竞争**是一个特定的竞争条件。任何时候，只要有两个以上的goroutine并发访问同一变量，且至少其中的一个是写操作的时候就会发生数据竞争。
**数据竞态**是指有两个或更多的操作访问同一块内存，并且至少有一个操作是写入，而这些操作在不同的线程中进行，不受同步机制的保护/
```
// Package bank implements a bank with only one account.
package bank
var balance int
func Deposit(amount int) { balance = balance + amount }
func Balance() int { return balance }

// Alice:
go func() {
    bank.Deposit(200)                // A1
    fmt.Println("=", bank.Balance()) // A2
}()

// Bob:
go bank.Deposit(100)                 // B
```

```
var x []int
go func() { x = make([]int, 10) }()
go func() { x = make([]int, 1000000) }()
x[999999] = 1 // NOTE: undefined behavior; memory corruption possible!
```

避免数据竞争
* 不去写变量
    在程序初始化阶段，初始化变量并且再也不去修改它们，那么任意数量的goroutine并发访问Icon都是安全的，因为每一个goroutine都只是去读取而已

    ```
    var icons = map[string]image.Image{
    "spades.png":   loadIcon("spades.png"),
    "hearts.png":   loadIcon("hearts.png"),
    "diamonds.png": loadIcon("diamonds.png"),
    "clubs.png":    loadIcon("clubs.png"),
    }

    // Concurrency-safe.
    func Icon(name string) image.Image { return icons[name] }
    ```

* 避免从多个goroutine访问变量
    * 把变量都限定在一个单独的goroutine中
        其他goroutine不能够直接访问变量，只能通过channel给指定的goroutine发送请求来查询更新变量，也就是使用通信来共享数据。
        一个提供对一个指定的变量通过channel来请求的goroutine叫这个变量的monitor goroutine
        ```
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
        ```
    * 串行绑定
        ```
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
        ```
* 互斥
    允许很多goroutine去并发访问变量，但在同一时刻最多只有一个goroutine访问变量

## 9.2 sync.Mutex互斥锁
一个只有1和0的信号量叫做二元信号量
```
var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
```
```
package bank

import "sync"

var (
	// sema    = make(chan struct{}, 1)
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	// sema <- struct{}{}
	mu.Lock()
	balance += amount
	// <-sema
	mu.Unlock()
}

func Balance() int {
	// sema <- struct{}{}
	// b := balance
	// <-sema
	// return b
	mu.Lock()
	defer mu.Unlock()
	return balance
}
```

Lock和Unlock之间的代码段叫做临界区
每一个函数在一开始就获取互斥锁并在最后释放锁，从而保证共享变量不会被并发访问，这种函数、互斥锁和变量的编排叫做监控monitor
deferred Unlock即使在临界区发生panic时依然会执行，这对于recover来恢复的程序是很重要的。
```
var (
	mu      sync.Mutex
	balance int
)

// 一个不导出函数，这个函数假设锁总是会被保持并去做实际的操作
func deposit(amount int) {
	balance += amount
}
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
// 在调用deposit函数前会先获取锁
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
```

## 9.3 sync.RWMutex读写锁

多读单写锁(multiple readers,single writer lock)，允许多个只读操作并行执行，但写操作会完全互斥。
```
var mu sync.RWMutex
var balance int

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
} 
```

## 9.4 内存同步

为什么Balance方法需要用到互斥条件
1. Balance不会在其他操作比如Withdraw"中间"执行
2. "同步"不仅仅是一堆goroutine执行顺序的问题，同样也会涉及到内存问题

所有并发的问题都可以用一致的、简单的既定的模式来规避。
将变量限定在goroutine内部，如果是多个goroutine都要访问的变量，使用互斥条件来访问

## 9.5 sync.Once惰性初始化

```
var mu sync.RWMutex
var icons map[string]image.Image

func Icon(name string) image.Image {
	mu.RLock()
	if icons != nil {
		defer mu.RUnlock()
		return icons[name]
	}
	mu.RUnlock()
	mu.Lock()
	// 再次判断icons,确保此时其他goroutine没有初始化icons
	if icons == nil {
		loadIcons()
	}
	defer mu.Unlock()
	return icons[name]
}
```

```
var loadIconsOnce sync.Once
var icons map[string]image.Image

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
```
这样可以避免在变量被初始化之前和其他goroutine共享变量


## 9.6 竞争条件检测
竞争检测器the race detector 
go build *.go -race
go run main.go -race

## 9.7 
memo1 memo5

## 9.8 Goroutine和线程
1. 动态栈
OS线程有一个固定大小的内存块(2MB)来做栈，栈会存储当前正在被调用或挂起的函数的内部变量。
goroutine会以一个很小的栈开始其生命周期(2KB),这个栈会保存其活跃或挂起的函数调用的本地变量，栈的大小不固定，会根据需要动态的伸缩。栈的最大值有1GB
2. Goroutine调度
操作系统内核调度OS线程，调用一个叫做scheduler的内核函数。因为OS线程是被内核调度的，所以从一个线程向另一个线程移动需要完整的上下文切换(保存一个用户线程的状态到内存，恢复另一个线程到寄存器，然后更新调度器的数据结构)。这几步操作很慢，因为其局部性很差需要几次内存访问，并且会增加运行cpu周期。
Go有自己的调度器，m:n调度(在n个操作系统线程上多工调度m个goroutine)。Go调度器只关注单独的Go程序中的goroutine。Go调度器被Go语言"建筑"本身进行调度。例如当一个goroutine调用了time.Sleep，或者被channel调用或者mutex操作阻塞时，调度器会使其进入休眠并开始执行另一个goroutine，直到时机到了再去唤醒第一个goroutine。这种调度方式不需要进入内核的上下文，所以重新调度一个goroutine比调度一个线程代价要低得多
3. GOMAXPROCS
GOMAXPROCS的变量决定了Go调度器会有多少个操作系统线程同时执行Go代码
4. Goroutine没有ID号
Go鼓励更为简单的模式，这种模式下参数对函数的影响都是显式的。这样不仅使程序变得更易读，而且会让我们自由地向一些给定的函数分配子任务时不用担心其身份信息影响行为。


