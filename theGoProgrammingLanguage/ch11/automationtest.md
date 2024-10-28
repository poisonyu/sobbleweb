## The Go Programming Language

# go test
功能测试函数、基准测试函数、示例函数
1. 功能测试函数是以Test前缀命名的函数，用来检测程序逻辑正确性
2. 基准测试函数的名称以Benchmark开头，用来测试性能
3. 示例函数以Example开头，用来提供机器检查过的文档

# Test函数
参数t提供了汇报测试失败和日志记录的功能
t.Fatal or t.Fatalf终止测试
测试错误消息的一般格式是"f(x)=y, want z"
-v 输出包中每个测试用例的名称和执行时间
go test -v
-run 使go test只运行匹配的测试函数
go test -v -run="French|Canal"

4. 外部测试包
通过将测试函数定义在外部测试包中来解决包循环引用

产品代码文件
`go list -f={{.GoFiles}} fmt`

包内测试文件
`go list -f={{.TestGoFiles}} fmt`

包外测试文件
`go list -f={{.XTestGoFiles}} fmt`

包内测试文件export_test.go中添加一些函数声明，将包内部的功能暴露给外部测试
```
// export_test.go
package fmt 

var IsSpace = isSpace 
```

5. 编写有效的测试
一个好的测试
输出该问题一个简洁、清晰的现象描述，以及其他上下文相关的信息
在一次运行中尝试报告多个错误
```
func TestSplit(t *testing.T) {
    s, sep := "a:b:c", ":"
    words := strings.Split(s, sep)
    if got, want := len(words), 3; got != want {
        t.Errorf("Split(%q, %q) returned %d words, want %d", s, sep, got, want)
    }
    // ...
}
```
测试函数报告调用函数的名称、输入以及输出表示的含义，显示区别实际值和期望值，并且即使在测试失败的情况下，也可以继续执行。下一步是在循环中执行这个测试。

6. 避免脆弱的测试
仅检查你关心的属性
首先测试程序中越来越简单和稳定的接口，然后是内部函数，寻找在程序进化过程中不会发生改变的子串，写一个稳定的函数从复杂的输出中提取核心的内容

# 11.3 覆盖率
ch7/eval/eval_test.go
`go test -v -run=Coverage`
`go test -run=Coverage -coverprofile=c.out`
`go tool cover -html=c.out`

# 11.4 Benchmark函数
ch11/word1/word_test.go
-bench的参数是一个匹配Benchmark函数名的一个正则表达式，~~“.”匹配包中所有的基准测试函数~~
~~`go test -bench=.`~~
`go test -bench=IsPalindrome`
最快的程序通常是那些内存分配次数最少的程序
-benchmen在报告中包含内存分配和数据统计
选择最小的缓冲区并带来最佳的性能表现
测试出每个算法的优缺点

# 11.5 性能剖析
发现关键代码
CPU性能剖析识别出执行过程中需要CPU最多的函数
堆性能剖析识别出负责分配最多内存的语句
阻塞性能剖析识别出那些阻塞协程最久的操作

获取性能剖析报告就，一次只能使用一个标记，如果使用多个标记，就会覆盖其他报告
`go test -cpuprofile=cpu.out`
`go test -blockprofile=block.out`
`go test -memprofile=mem.out`