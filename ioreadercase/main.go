package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type upperCaseReader struct {
	reader io.Reader
}

func (r *upperCaseReader) Read(b []byte) (int, error) {
	n, err := r.reader.Read(b)
	if n > 0 {
		newStr := strings.ToUpper(string(b))
		copy(b, []byte(newStr))
	}
	return n, err
}

func main() {
	//io.Reader接口提供了一系列的实现，包括文件，内存，网络，压缩加密等
	//Read方法返回的是字节切片的字节数和错误值，当读取到结尾会返回一个io.EOF错误

	//2使用不同的reader实现对reader进行包装从而读取不同效果,转大写

	r := strings.NewReader("Hello World")
	upper := upperCaseReader{r}
	b := make([]byte, 1024)
	for {
		n, err := upper.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		fmt.Printf("value = %v, n = %d \n", string(b), n)
	}

	//1文件读取。使用字节数组切片拼接读取保证buffer空间应对未知大小的文件

	f, err := os.Open("./ioreadercase/test.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	fd := make([][]byte, 0)
	var len int64 = 0
	for {
		buffer := make([]byte, 8)
		n, err := f.ReadAt(buffer, len)
		if n > 0 {
			len = len + int64(n)
			fd = append(fd, buffer)
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
	}
	data := bytes.Join(fd, []byte(""))
	fmt.Printf("file value = %s, n = %d \n", string(data), len)

}
