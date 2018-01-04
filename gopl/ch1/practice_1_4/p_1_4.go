//dup2 prints the count and text of lines that appear more than once
//in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineInfo struct {
	count    int
	filename map[string]int
}

func NewLineInfo() *lineInfo {
	lInfo := &lineInfo{}
	lInfo.count = 0
	lInfo.filename = make(map[string]int)
	return lInfo
}

func main() {
	counts := make(map[string]*lineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, v := range files {
			f, err := os.Open(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for k, v := range counts {
		if v.count > 1 {
			fmt.Println(k, v)
		}
	}
}

func countLines(f *os.File, counts map[string]*lineInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if v, ok := counts[input.Text()]; ok {
			v.count++
			v.filename[f.Name()] = 1
		} else {
			counts[input.Text()] = NewLineInfo()
			curLineInfo := counts[input.Text()]
			curLineInfo.count++
			curLineInfo.filename[f.Name()] = 1
		}
	}
}
