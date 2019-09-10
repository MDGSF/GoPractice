package main

import (
	"github.com/MDGSF/utils"
)

func main() {
	EventLogFileChan := make(chan []byte, 16)
	p := CreateEventLogProcessor(".", EventLogFileChan, 4*1024*1024)
	p.Start()

	for {
		out := utils.GetRandomBytes(100)
		out = append(out, '\n')
		EventLogFileChan <- out
	}
}
