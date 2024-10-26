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

