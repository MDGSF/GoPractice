package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MDGSF/utils/log"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		p := make([]byte, 1024)
		readed, err := reader.Read(p)
		if err != nil {
			log.Error("err = %v", err)
			continue
		}
		// fmt.Println(string(p[:readed]))
		fmt.Println(p[:readed], string(p[:readed]))
	}
}
