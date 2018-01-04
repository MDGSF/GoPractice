//dup2 prints the count and text of lines that appear more than once
//in the input. It reads from stdin or from a list of named files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Println("k = ", k)
			fmt.Println("v = ", v)
		}
	}
}
