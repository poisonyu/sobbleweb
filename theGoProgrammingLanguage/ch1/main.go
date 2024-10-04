package main

/*
1.6 并发获取多个URL goroutin channel
main函数本身也运行在一个goroutine中

time.Since(start).Seconds()

nbytes, err := io.Copy(ioutil.Discard, src)

resp, err := http.Get(url)
//...
resp.Body.Close() // don't leak resources

1.7 Web服务
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }
}

1.8 本章要点 switch
tagless switch
switch不带操作对象时默认用true值代替，然后将每个case的表达式和true值进行比较
func Signum(x int) int {
    switch {
    case x > 0:
        return +1
    default:
        return 0
    case x < 0:
        return -1
    }
}
&操作符可以返回一个变量的内存地址，
*操作符可以获取指针指向的变量内容
*/
