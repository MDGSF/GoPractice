package main

import (
	"fmt"
	"time"
)

type (
	Event struct {
		Data int64
	}

	Observer interface {
		OnNotify(Event)
	}

	Notifier interface {
		Register(Observer)

		Deregister(Observer)

		Notify(Event)
	}
)

type eventObserver struct {
	id int
}

type eventNotifier struct {
	observers map[Observer]struct{}
}

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("Observer %d received: %d\n", o.id, e.Data)
}

func (n *eventNotifier) Register(o Observer) {
	n.observers[o] = struct{}{}
}

func (n *eventNotifier) Deregister(o Observer) {
	delete(n.observers, o)
}

func (n *eventNotifier) Notify(e Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}

func main() {
	n := eventNotifier{observers: map[Observer]struct{}{}}
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}
