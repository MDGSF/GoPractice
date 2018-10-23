package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println(TarNameCheck("lane_20180609_370.tar.gz"))
	fmt.Println(TarNameCheck("lane_huangping.tar.gz"))
}

func TarNameCheck(zipName string) bool {
	unzipName := strings.TrimSuffix(zipName, ".tar.gz")

	cmd := exec.Command("bash", "-c", "tar -tf "+zipName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	arr := strings.Split(string(output), "\n")
	if len(arr) == 0 {
		return false
	}

	firstline := arr[0]

	return strings.HasPrefix(firstline, unzipName)
}
