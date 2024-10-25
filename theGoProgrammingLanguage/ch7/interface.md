## 

# 7.6 sort.Interface接口
```
package sort
type Interface interface{
    Len() int 
    Less(i, j int) bool
    Swap(i, j int)
}
```

给[]*Track排序见sorting/sorting.go

```
a := []int{1, 6, 2, 4, 3, 5}
// 给切片排序
// sort.IntSlice类型实现了sort.Interface接口，
// 把a转换为sort.IntSlice类型，再调用sort.Sort()方法排序
**sort.Sort(sort.IntSlice(a))**

// 反向排序
// 先把类型变为sort.IntSlice,这样就实现了sort.Interface接口
// sort.Reverse覆盖(修改)Less方法，使排序反向
sort.Sort(sort.Reverse(sort.IntSlice(a)))

// 判断一个序列是否排序
sort.IsSorted(sort.IntSlice(a))

// []int封装函数
sort.Ints(a)
slices.Sort(a) // go1.22
sort.IntsAreSorted(a)


// sort.Stable() todo 
```

# 7.7 http.Handler接口
见http2,http3
```
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
func ListenAndServe(address string, h handler) error
```
# 7.8 error接口
```
type error interface {
    Error() string
}

fmt.Errorf()
```


# 7.10 类型断言
x.(T) x是接口值(操作数)，T是断言类型
类型断言会检查x的动态类型是否满足指定的断言类型
1. 断言类型T是一个具体的类型
判断接口值x的动态类型是否是断言类型T，如果是,断言类型的结果就是接口值x的动态值,也就是从操作数中取出动态值
```
var x io.Writer
x = os.Stdout
f := x.(*os.File) // 取出动态值，赋值给f
```

2. 断言类型T是一个接口类型
判断接口值x的动态类型是否实现了接口类型T，如果是，返回一个接口类型是T，动态类型和动态值不变的接口值。
```
var x io.Writer
x = os.Stdout
rw := x.(io.ReadWriter) // 返回一个接口类型是T，动态类型和动态值不变的接口值
```

3. 检测一个接口值的动态类型
```
var w io.Writer
w = os.Stdout
// ok为false时，f为断言类型的零值
if f, ok := w.(*os.File); ok {
    // ...use f...
}
```
# 7.11 类型断言识别错误类型
接口值的比较？
```
var ErrNotExist = errors.New("file does not exist")
func IsNotExist(err error) bool {
    if pe, ok := err.(*PathError); ok {
        err = pe.Err
    }
    return err == syscall.ENOENT || err == ErrNotExist
}
```

# 7.12 通过接口类型断言来查询特性
对一个接口值用一个接口类型断言，如果成功，就可以调用断言类型(接口)的方法
也就是断言类型T是一个接口类型这种情况
```
func writeString(w io.Writer, s string) (n int, err error) {
    // 类型断言，将io.Writer接口转换成io.StringWriter接口
	if sw, ok := w.(io.StringWriter); ok {
		return sw.WriteString(s) // 避免了内存复制
	}
	return w.Write([]byte(s)) // 分配了临时内存
}

func writeHeader(w io.Writer, contentType string) error {
    if _, err := writeString(w, "Content-Type"); err != nil {
        return err
    }
}
```

# 7.13 类型分支
接口能够容纳各种具体类型
```
import "database/sql"

func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	result, err := db.Exec(
		"SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?", 
		artist, minYear, maxYear
	)
	// ...
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x的类型为interface{}
	case bool:
		if x { // bool
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(s) // string
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
```

# 7.14 示例：基于标记的XML解析

