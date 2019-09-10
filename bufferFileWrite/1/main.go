package main

import (
	"bufio"
	"os"
	"time"

	"github.com/MDGSF/utils/log"
)

func main() {
	fileName := "test.txt"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	defer f.Close()

	fw := bufio.NewWriter(f)
	defer fw.Flush()

	nn, err := fw.Write([]byte("a"))
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	log.Info("nn = %v", nn)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
	}
}
