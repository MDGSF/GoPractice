package main

import (
	"encoding/hex"
	"log"

	"golang.org/x/crypto/sha3"
)

func main() {
	input := "huangjian"
	output := make([]byte, 64)
	sha3.ShakeSum256(output, []byte(input))
	log.Println(input)
	outputStr := hex.EncodeToString(output)
	log.Println(outputStr)
}
