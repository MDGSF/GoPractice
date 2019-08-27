package main

import (
	"fmt"
	"time"
)

func main() {
	curTime := time.Now()
	curTime.Unix()
	year, month, day := curTime.Date()
	fmt.Println(fmt.Sprintf("%04d-%02d-%02d", year, month, day))
}
