// go的两种设计模式
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
)

// 函数选项模式

type Server struct {
	Addr         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Timeout      time.Duration
}

type Option func(*Server)

func WithAddr(address string) Option {
	return func(server *Server) {
		server.Addr = address
	}
}
func WithPort(port string) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeout = timeout
	}
}

func NewServer(options ...Option) *Server {
	server := &Server{
		Addr:         "localhost",
		Port:         "4060",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		Timeout:      2 * time.Second,
	}
	for _, option := range options {
		option(server)
	}
	return server
}

func TestFunctionOption() {
	srv := NewServer()
	fmt.Println(srv)
	srv2 := NewServer(WithAddr("192.168.0.1"), WithPort("8080"), WithReadTimeout(2*time.Minute))
	fmt.Println(srv2)
}

// 装饰器模式
type Hf func(http.ResponseWriter, *http.Request)

func Logger1(hf Hf) Hf {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		hf(w, r)
		log.Printf("elapse time: %v\n", time.Since(now))
	}
}

// Logger中间件
func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("url: %s, elapse: %v\n", r.URL, time.Since(now))
	}
	return http.HandlerFunc(fn)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

func HowAreYou(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I am fine"))
}

func TestDecoration() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", HelloWorld)
	mux.HandleFunc("GET /how", HowAreYou)

	srv := http.Server{
		Addr:    ":4060",
		Handler: Logger(mux),
	}

	fmt.Println("listen at:", srv.Addr)
	srv.ListenAndServe()
}

func T(c *gin.Context) {
	c.GetHeader("Accept-Language")
	c.Header("cache-control", "public, max-age=31536000")
	c.Writer.Header().Set("key", "value")
	c.Writer.Header().Del("key")

	strings.HasPrefix(c.Request.RequestURI, "/static/")

	// strings.ParseInt("")
	// strings.Atoi("dog")
	// path := c.Request.URL.Path

}

func test2() {
	path := filepath.Join("upload", "pic.jpg")

	fmt.Println(path) // upload/pic.jpg
	path = filepath.Dir(path)
	fmt.Println("path:", path)      // upload/
	fmt.Println(filepath.Dir(path)) // .

	tpath := reflect.TypeOf(path)
	fmt.Println(tpath) // string
	tpath.Kind()       // string
	// for i := 0; i < tpath.NumField(); i++ {
	// }
	var a *int
	ta := reflect.TypeOf(a)
	fmt.Println(ta) // *int
	ta.Kind()       // ptr
	ta.Elem()       // int

	// type MyInt int
	// var myint MyInt = 2
	// value = reflect.ValueOf(myint)
	// ty := value.TypeOf()
}

func L() {
	// l := list.New()
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(cap(slice))
	s := slice[1:3:4]
	fmt.Println(cap(s))

}

func GenerateVerificationCode(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(10))
	}
	fmt.Println(b)
	return bytes2string(b)

}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ReadString() {

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("*********")
	for input.Scan() {
		fmt.Println(input.Text())
	}

}
func main() {
	// TestDecoration()

	// gin.DebugMode

	// gin.SetMode(gin.ReleaseMode)
	// router := gin.Default()
	// userApi := router.Group("user")
	// router.SetHTMLTemplate()
	// router.NoRoute()
	// router.Handle("GET", "/")
	// router.GET()
	// router.StaticFS()

	// ctx := context.Background()

	// fmt.Println(ctx.Done())
	// ctx.Deadline()
	// ctx.Err()
	// ctx.Value()

	// test2()

	// s := make([]int, 10)
	// a := &s
	// b := uintptr(8)
	// Len := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + b))
	// fmt.Println(Len)
	// unsafe.Sizeof()
	// mp := make(map[string]int)

	b := []byte{'c', 'f', 'd'}
	s := bytes2string(b)
	fmt.Println(s, string(b))
	ReadString()

}
