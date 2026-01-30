//go:build defer
// +build defer

package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func cc() (i int) {
	defer func() { i++ }()
	return i

}

func main() {
	//a()
	//b()
	cc()
}
