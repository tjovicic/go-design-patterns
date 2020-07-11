package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	Event struct {
		data int
	}

	Observer interface {
		ID() int
		NotifyCallback(Event)
	}

	Subject interface {
		AddListener(Observer)
		RemoveListener(Observer)
		Notify(Event)
	}

	eventObserver struct {
		id   int
		time time.Time
	}

	eventSubject struct {
		observers sync.Map
	}
)

func (e *eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Observer: %d Received: %d after %v\n", e.id, event.data, time.Since(e.time))
}

func (e *eventObserver) ID() int {
	return e.id
}

func (s *eventSubject) AddListener(o Observer) {
	s.observers.Store(o, struct{}{})
}

func (s *eventSubject) RemoveListener(o Observer) {
	s.observers.Delete(o)
}

func (s *eventSubject) Notify(e Event) {
	s.observers.Range(func(key interface{}, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}

		key.(Observer).NotifyCallback(e)
		return true
	})
}

func fib(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()

	return out
}

func main() {
	n := eventSubject{
		observers: sync.Map{},
	}

	obs1 := eventObserver{id: 1, time: time.Now()}
	obs2 := eventObserver{id: 2, time: time.Now()}
	n.AddListener(&obs1)
	n.AddListener(&obs2)

	go func() {
		select {
		case <-time.After(time.Millisecond * 10):
			n.RemoveListener(&obs1)
		}
	}()

	for x := range fib(10000) {
		n.Notify(Event{data: x})
	}
}
