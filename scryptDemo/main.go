package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/scrypt"
)

func main() {
	fmt.Println("vim-go")

	password := "123456"
	salt := []byte{0x12, 0x31, 0x11, 0x11, 0x12, 0x31, 0x11, 0x11}

	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(dk))
}
