// 10.5 空导入
// 对包级别的变量执行初始化表达式求值，并执行init函数

package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	// _ "image/png" // 注册PNG解码器
	"io"
	"log"
	"os"
	// "github.com/chai2010/webp"
	// "golang.org/x/image/bmp"
)

const inputImage = "input.png"

// const outputImge = "output"

var format = flag.String("f", "png", "the image output format")

func main() {
	// f, err := os.Open("pngencode.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// wf, err := os.Create("jpegencode.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer wf.Close()
	// if err := toJPEG(f, wf); err != nil {
	// 	fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
	// 	os.Exit(1)
	// }

	flag.Parse()
	input, err := os.Open(inputImage)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	img, _, err := CheckInputFormat(input)
	if err != nil {
		log.Fatal(err)
	}
	switch *format {
	case "png":
		w, err := createOutputFile(*format)
		if err != nil {
			log.Fatal(err)
		}
		err = png.Encode(w, img)
		w.Close()
		if err != nil {
			log.Fatal(err)
		}
	case "jpg":
		w, err := createOutputFile(*format)
		if err != nil {
			log.Fatal(err)
		}
		err = jpeg.Encode(w, img, &jpeg.Options{Quality: 95})
		w.Close()
		if err != nil {
			log.Fatal(err)
		}
	case "gif":
		w, err := createOutputFile(*format)
		if err != nil {
			log.Fatal(err)
		}
		err = gif.Encode(w, img, &gif.Options{})
		w.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

}

func createOutputFile(kind string) (*os.File, error) {
	w, err := os.Create("output." + kind)
	if err != nil {
		return nil, err
	}
	return w, err
}
func CheckInputFormat(in io.Reader) (image.Image, string, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, "", err
	}
	fmt.Println("Input image format = ", kind)
	return img, kind, nil
}
func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})

}

// package main

// import (
// 	"database/sql"

// 	_ "github.com/go-sql-driver/mysql" // 添加MySQL支持
// 	_ "github.com/lib/pq"              // 添加Postgres支持
// )

// func main() {
// 	db, err := sql.Open("postgres", "dbname")
// 	db, err = sql.Open("mysql", "dbname")
// 	db, err = sql.Open("sqlite3", "dbname")
// }
