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
