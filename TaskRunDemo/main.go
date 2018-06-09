package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/MDGSF/utils/log"
)

func main() {
	t := NewTask()
	for {
		t.Start()
		time.Sleep(time.Second * 5)
		t.Stop()
	}
	//shutdown := make(chan bool)
	//<-shutdown
}

type Task struct {
	loopClosed chan bool
	closed     chan bool
	wg         *sync.WaitGroup
}

func NewTask() *Task {
	t := &Task{}
	t.loopClosed = make(chan bool)
	t.closed = make(chan bool)
	t.wg = &sync.WaitGroup{}
	return t
}

func (t *Task) Stop() {
	close(t.loopClosed)
	t.stopSubRoutine()
}

func (t *Task) stopSubRoutine() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("stop sub goroutine panic recover, err = %v", err)
		}
	}()

	close(t.closed)
	t.wg.Wait()
}

func (t *Task) Start() {
	fmt.Println("task start......................")
	go t.run()
}

func (t *Task) run() {
	t.loopClosed = make(chan bool)
	for {
		select {
		case <-t.loopClosed:
			return
		default:
		}

		t.loop()
	}
}

func (t *Task) loop() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("loop panic recover, err = %v", err)
		}
	}()

	t.wg.Add(3)
	t.closed = make(chan bool)
	readExit := t.StartGoroutine(t.readLoop)
	writeExit := t.StartGoroutine(t.writeLoop)
	pingExit := t.StartGoroutine(t.pingLoop)
	select {
	case <-t.closed:
	case <-readExit:
		t.stopSubRoutine()
	case <-writeExit:
		t.stopSubRoutine()
	case <-pingExit:
		t.stopSubRoutine()
	}
	panic("loop panic")
}

func (t *Task) StartGoroutine(f func(chan bool)) chan bool {
	exit := make(chan bool)
	go f(exit)
	return exit
}

func (t *Task) readLoop(exit chan bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("read panic recover, err = %v", err)
		}
		t.wg.Done()
		close(exit)
		fmt.Println("read exit")
	}()

	for {
		select {
		case <-t.closed:
			return
		default:
			for i := 0; i < 3; i++ {
				fmt.Println("do reading")
				time.Sleep(time.Second)
			}
			panic("read panic")
		}
	}
}

func (t *Task) writeLoop(exit chan bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("write panic recover, err = %v", err)
		}
		t.wg.Done()
		close(exit)
		fmt.Println("write exit")
	}()

	for {
		select {
		case <-t.closed:
			return
		default:
			fmt.Println("do writing")
			time.Sleep(time.Second)
		}
	}
}

func (t *Task) pingLoop(exit chan bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("ping panic recover, err = %v", err)
		}
		t.wg.Done()
		close(exit)
		fmt.Println("ping exit")
	}()

	for {
		select {
		case <-t.closed:
			return
		default:
			fmt.Println("do ping")
			time.Sleep(time.Second)
		}
	}
}
