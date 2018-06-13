package main

import (
	"fmt"
	"io"
	"net"

	"github.com/MDGSF/utils/log"
	"github.com/xtaci/smux"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}()

	addr := ln.Addr().String()
	cliconn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	session, _ := smux.Client(cliconn, nil)
	defer session.Close()
	stream, _ := session.OpenStream()

	buf := make([]byte, 10)
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("hello%v", i)
		stream.Write([]byte(msg))
		n, err := stream.Read(buf)
		if err != nil {
			panic(err)
		}
		log.Info("%v", string(buf[:n]))
	}
}

func handleConnection(conn net.Conn) {
	session, _ := smux.Server(conn, nil)
	for {
		stream, err := session.AcceptStream()
		if err != nil {
			return
		}

		go func(s io.ReadWriteCloser) {
			buf := make([]byte, 65536)
			for {
				n, err := s.Read(buf)
				if err != nil {
					return
				}
				s.Write(buf[:n])
			}
		}(stream)
	}
}
