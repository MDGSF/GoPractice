package main

import (
	"fmt"
	"sync"
	"time"
)

type TService struct {
	stopChan  chan bool
	waitGroup *sync.WaitGroup
}

func NewService() *TService {
	s := &TService{}
	s.stopChan = make(chan bool)
	s.waitGroup = &sync.WaitGroup{}
	return s
}

func (s *TService) Start() {
	go s.run()
}

func (s *TService) Stop() {
	close(s.stopChan)
	s.waitGroup.Wait()
}

func (s *TService) run() {
	s.waitGroup.Add(1)
	defer s.waitGroup.Done()
	fmt.Println("run start")

	go s.doRead()
	go s.doWrite()

	for {
		select {
		case <-s.stopChan:
			fmt.Println("run stop")
			return
		default:
		}

		fmt.Println("I'm running...")
		time.Sleep(time.Second)
	}
}

func (s *TService) doRead() {
	s.waitGroup.Add(1)
	defer s.waitGroup.Done()
	fmt.Println("read start")

	for {
		select {
		case <-s.stopChan:
			fmt.Println("read stop")
			return
		default:
		}

		fmt.Println("I'm reading...")
		time.Sleep(time.Second)
	}
}

func (s *TService) doWrite() {
	s.waitGroup.Add(1)
	defer s.waitGroup.Done()
	fmt.Println("write start")

	for {
		select {
		case <-s.stopChan:
			fmt.Println("write stop")
			return
		default:
		}

		fmt.Println("I'm writing...")
		time.Sleep(time.Second)
	}
}

func main() {
	s := NewService()
	s.Start()

	fmt.Println("main start sleep 5 second...")
	time.Sleep(5 * time.Second)
	fmt.Println("main wake up")

	s.Stop()
}
