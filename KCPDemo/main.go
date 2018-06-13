package main

import (
	"fmt"

	"github.com/MDGSF/utils/log"
	kcp "github.com/xtaci/kcp-go"
)

func main() {
	ln, err := kcp.ListenWithOptions("127.0.0.1:12580", nil, 10, 3)
	if err != nil {
		panic(err)
	}

	go func() {
		sess, err := ln.AcceptKCP()
		if err != nil {
			panic(err)
		}

		go handleConnection(sess)
	}()

	kcpconn, err := kcp.DialWithOptions("127.0.0.1:12580", nil, 10, 3)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 10)
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("hello%v", i)
		kcpconn.Write([]byte(msg))
		n, err := kcpconn.Read(buf)
		if err != nil {
			panic(err)
		}
		log.Info("%v", string(buf[:n]))
	}
}

func handleConnection(sess *kcp.UDPSession) {
	buf := make([]byte, 4096)
	for {
		n, err := sess.Read(buf)
		if err != nil {
			panic(err)
		}
		sess.Write(buf[:n])
	}
}
