package workerpool

var (
	MaxJobQueueNum = 1000000
)

var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, MaxJobQueueNum)
}
