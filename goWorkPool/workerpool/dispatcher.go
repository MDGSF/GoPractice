package workerpool

import "fmt"

type Dispatcher struct {
	WorkerPool chan chan Job
	Len        int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, maxWorkers),
		Len:        maxWorkers,
	}
}

func (d *Dispatcher) Start() {
	fmt.Println("length of workerpool:", d.Len)
	for i := 0; i < d.Len; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.run()
}

func (d *Dispatcher) run() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}
