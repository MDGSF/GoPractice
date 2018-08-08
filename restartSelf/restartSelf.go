package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("vim-go")
	time.Sleep(time.Second)

	v := os.Getenv("restartSelf")
	if v == "1" {
		fmt.Println("restart self success")
		return
	}

	restart()
}

func restart() {
	fmt.Println("start to restart self")
	os.Setenv("restartSelf", "1")
	exe := os.Args[0]
	args := os.Args[1:]
	cmd := exec.Command(exe, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
