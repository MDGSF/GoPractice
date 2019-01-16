package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8080"
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("host = %v, port = %v\n", host, port)

	addr2 := net.JoinHostPort(host, port)
	fmt.Printf("addr2 = %v\n", addr2)

	ip := net.ParseIP(host).To4()
	b := []byte(ip)
	fmt.Printf("ip array = %v\n", b)

	ip2 := net.IPv4(b[0], b[1], b[2], b[3]).To4()
	fmt.Printf("ip2 string = %v\n", ip2.String())
}
