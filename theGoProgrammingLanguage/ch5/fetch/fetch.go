package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	// 没有对f.close采用defer机制，因为这会产生一些微妙的错误。
	// 许多文件系统，尤其是NFS，写入文件时发生的错误会被延迟到文件关闭时反馈。
	//如果没有检查文件关闭时的反馈信息，可能会导致数据丢失，而我们还误以为写入操作成功。
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	// 优先返回io.Copy的错误信息，它先于f.Close发生
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err

}
