package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "no arguments\n")
		os.Exit(1)
	}

	pid := syscall.Getpid()
	pgid, _ := syscall.Getpgid(pid)
	fmt.Printf("pid = %v, ppid = %v, pgid = %v\n", pid, syscall.Getppid(), pgid)

	// if pid == pgid, Setsid will failed.
	_, err := syscall.Setsid()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Setsid failed, err = %v\n", err)
		os.Exit(2)
	}

	cmd := exec.Command(args[0], args[1:]...)
	//cmd.Dir = "/"
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "exec Command failed, err = %v\n", err)
		os.Exit(3)
	}

	if cmd.Process == nil {
		fmt.Fprintf(os.Stderr, "Command process is nil\n")
		os.Exit(4)
	}

	fmt.Printf("child process id = %v\n", cmd.Process.Pid)
}
