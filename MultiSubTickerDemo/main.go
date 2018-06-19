package main

import (
	"fmt"
	"time"

	"github.com/MDGSF/utils/log"
)

type TSub struct {
	Name     string
	Timeout  int
	F        func()
	MainChan chan *TMsg2Main
}

func (sub *TSub) Vacuum() {
	fmt.Println("vacuum")
}

func (sub *TSub) Start() {
	go sub.run()
}

func (sub *TSub) run() {
	t := time.NewTicker(time.Second * time.Duration(sub.Timeout))
	for {
		select {
		case <-t.C:
			log.Info("%v triggered", sub.Name)
			msg := &TMsg2Main{F: sub.Vacuum}
			sub.MainChan <- msg
		}
	}
}

type TMsg2Main struct {
	F func()
}

func main() {

	var mainChan chan *TMsg2Main
	mainChan = make(chan *TMsg2Main)

	a := make([]*TSub, 0)
	sub1 := &TSub{}
	sub1.Name = "sub1"
	sub1.Timeout = 1
	sub1.MainChan = mainChan
	a = append(a, sub1)

	sub2 := &TSub{}
	sub2.Name = "sub2"
	sub2.Timeout = 2
	sub2.MainChan = mainChan
	a = append(a, sub2)

	sub1.Start()
	sub2.Start()

	for {
		select {
		case data, ok := <-mainChan:
			if !ok {
				log.Error("no ok")
				return
			}
			log.Info("main receiver data = %v", data)
			data.F()
		}
	}
}
