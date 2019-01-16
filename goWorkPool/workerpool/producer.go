package workerpool

type Producer struct {
	maxJob int
}

func NewProducer(maxJob int) *Producer {
	return &Producer{
		maxJob: maxJob,
	}
}

func (p *Producer) Run() {
	for i := 0; i < p.maxJob; i++ {
		//fmt.Println("produce job", i)
		job := Job{Payload: Payload(i)}
		JobQueue <- job
	}
}
