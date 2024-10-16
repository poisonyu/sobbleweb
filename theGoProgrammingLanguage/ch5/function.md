##

# 5.4 错误 
五种错误处理策略
1. 传播错误
函数中某个子程序的失败，会变成该函数的失败
```
func findLinks(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    // ...
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        // 给返回的错误加入了发生错误时的解析器，发生错误的url
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
    // ...
}
```
>由于错误信息经常是以链式组合在一起的，所以错误信息中应避免大写和换行符
>一般而言，被调用函数f(x)会将调用信息和参数信息作为发生错误时的上下文放在错误信息中并返回给调用者，调用者需要添加一些错误信息中不包含的信息，比如添加url到html.Parse返回的错误中。

2. 重新尝试失败的操作
如果错误的发生时偶然性的，或由不可预知的问题导致的，可以重试，但要限制重试的时间间隔或重试次数
```
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s);retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
```

3. 输出错误信息并结束程序(应用在main中)
错误发生后，程序无法继续运行
```
func main() {
    if err := WaitForServer(url); err != nil {
        fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
        os.Exit(1)
        // or 
        log.Fatalf("Site is down: %v\n", err)
    }
}
```

4. 有时
只需要输出错误信息，不需要中断程序运行
```
if err := Ping(); err != nil {
    // log包中的所有函数会为没有换行符的字符串增加换行符
    log.Printf("ping failed: %v; networking disabled", err)
    // or
    fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
}
```

5. 直接忽略错误
```
	dir, err := os.MkdirTemp("", "scratch")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}
    // 操作系统会定期清理临时目录
	os.RemoveAll(dir)
```

EOF
```
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read failed:%v", err)
		}
		// ...use r...
	}
```
