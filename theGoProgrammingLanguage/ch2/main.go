package main

/*
对于局部变量，尽量使用短小的名字
如果一个名字的作用域比较大，生命周期也比较长，那么用长的名字将会更有意义
像ASCII和HTML这样的缩略词则避免使用大小写混合的写法，它们可能被称为htmlEscape、HTMLEscape或escapeHTML，但不会是escapeHtml
在标准库有QuoteRuneToASCII和parseRequestLine这样的函数命名

接口或引用类型（包括slice、指针、map、chan和函数）变量对应的零值是nil
数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值
零值初始化机制可以确保每个声明的变量总是有一个良好定义的值，因此在Go语言中不存在未初始化的变量

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
    flag.Parse() // 用于更新每个标志参数对应变量的值（之前是默认值）
	flag.Args() // 访问非标志参数的普通命令行参数
    fmt.Print(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()
    }
}

new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T

由于new只是一个预定义的函数，它并不是一个关键字，因此我们可以将new名字重新定义为别的类型
func delta(old, new int) int { return new - old }
由于new被定义为int类型的变量名，因此在delta函数内部是无法使用内置的new函数的

img.SetColorIndex(
	size+int(x*size+0.5), size+int(y*size+0.5),
	blackIndex, // 最后插入的逗号不会导致编译错误，这是Go编译器的一个特性
)               // 小括弧另起一行缩进，和大括弧的风格保存一致

逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响
如果将指向短生命周期对象的指针保存到具有长生命周期的对象中，特别是保存到全局变量时，会阻止对短生命周期对象的垃圾回收（从而可能影响程序的性能）。
var global *int

func f() {
    var x int
    x = 1
    global = &x
}

某些表格数据初始化并不是一个简单的赋值过程。在这种情况下，我们可以用一个特殊的init初始化函数来简化初始化工作。每个文件都可以包含多个init初始化函数


func init() { }


2.7 作用域
声明语句的作用域是指源代码中可以有效使用这个名字的范围。

声明语句的作用域对应的是一个源代码的文本区域；它是一个编译时的属性。


一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个运行时的概念。

Go语言的习惯是在if中处理错误然后直接返回，这样可以确保正常执行的语句不需要代码缩进


var cwd string

func init() {
    var err error
    cwd, err = os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}





*/
