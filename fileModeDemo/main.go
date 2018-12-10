package main

import (
	"log"
	"os"
)

func main() {
	f666Info, err := os.Stat("f666")
	if err != nil {
		log.Fatal(err)
	}

	// file perm 这个是以 8 进制来存储的
	log.Printf("%v %d %o", f666Info.Mode(), f666Info.Mode(), f666Info.Mode())
	log.Printf("%v %d %o", f666Info.Mode()&0400, f666Info.Mode()&0400, f666Info.Mode()&0400)

	f777Info, _ := os.Stat("f777")
	log.Printf("%v %d %o", f777Info.Mode(), f777Info.Mode(), f777Info.Mode())
}
