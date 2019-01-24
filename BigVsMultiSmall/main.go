package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("vim-go")
	//testSmall()
	testBig()
	fmt.Println("main end")
}

func testBig() {
	start := time.Now()
	dir := "hjtestBig2/images/jsw28355tuheipictures_8_from24816to28355"
	entries, _ := ioutil.ReadDir(dir)
	big, _ := os.Create("/home/huangjian/a/big")

	for _, v := range entries {
		path := filepath.Join(dir, v.Name())
		f, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}

		io.Copy(big, f)
		f.Close()
	}

	big.Close()
	duration := time.Since(start)
	fmt.Println(duration)
}

func testSmall() {
	start := time.Now()
	dir := "hjtestBig2/images/jsw28355tuheipictures_8_from24816to28355"
	entries, _ := ioutil.ReadDir(dir)
	smalldir := "/home/huangjian/a/small"

	for _, v := range entries {
		path := filepath.Join(dir, v.Name())
		f, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}

		small := filepath.Join(smalldir, v.Name())
		smallFile, _ := os.Create(small)
		io.Copy(smallFile, f)

		f.Close()
		smallFile.Close()
	}

	duration := time.Since(start)
	fmt.Println(duration)
}
