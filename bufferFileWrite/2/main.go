package main

import (
	"os"
	"time"

	"github.com/MDGSF/utils/log"
)

func main() {

	fileName := "test.txt"
	//f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	defer f.Close()

	//wg := &sync.WaitGroup{}
	//wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go func(f *os.File) {
			for {
				f.Write([]byte("{\"name\": \"huangjian\", \"key\": 123456789, \"age\": 10000}{\"name\": \"huangjian\", \"key\": 123456789, \"age\": 10000}{\"name\": \"huangjian\", \"key\": 123456789, \"age\": 10000}{\"name\": \"huangjian\", \"key\": 123456789, \"age\": 10000}{\"name\": \"huangjian\", \"key\": 123456789, \"age\": 10000}{\"name\": \"huangjian\", \"key\": 123456789, \"age\": 10000}{\"name\": \"huangjian\", \"key\": 123456789, \"age\": 10000}\n"))
				time.Sleep(time.Millisecond * 100)
			}
			//wg.Done()
		}(f)
	}

	//wg.Wait()

	start := time.Now()
	for {
		cur := time.Now()
		if cur.Sub(start)/1000000000 > 10 {
			break
		}
	}
}
