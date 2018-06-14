package pool

type Object struct {
	id int
}

type Pool chan *Object

func New(total int) Pool {
	p := make(Pool, total)
	for i := 0; i < total; i++ {
		p <- new(Object)
	}
	return p
}
