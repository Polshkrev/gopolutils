package tests

import (
	"sync"
	"testing"

	"github.com/Polshkrev/gopolutils/events"
)

func TestNewEventManager(t *testing.T) {
	var manager events.EventManager = events.NewEventManager()

	if manager == nil {
		t.Fatal("expected non-nil event manager")
	}
}

func TestSubscribeAndPost(t *testing.T) {
	const eventType events.EventType = "test"

	var called bool
	var received any

	events.Subscribe(eventType, func(data any) {
		called = true
		received = data
	})

	events.Post(eventType, "hello")

	if !called {
		t.Fatal("event was not called")
	} else if received != "hello" {
		t.Fatalf("expected \"hello\", got %v", received)
	}
}

func TestPostWithoutSubscribers(t *testing.T) {
	const eventType events.EventType = "does_not_exist"

	// Should not panic.
	events.Post(eventType, "ignored")
}

func TestMultipleSubscribers(t *testing.T) {
	const eventType events.EventType = "multiple"

	var first bool
	var second bool
	var third bool

	events.Subscribe(eventType, func(any) {
		first = true
	})

	events.Subscribe(eventType, func(any) {
		second = true
	})

	events.Subscribe(eventType, func(any) {
		third = true
	})

	events.Post(eventType, nil)

	if !first || !second || !third {
		t.Fatal("expected every subscriber to be called")
	}
}

func TestSeparateEventTypes(t *testing.T) {
	const (
		firstEvent  events.EventType = "first"
		secondEvent events.EventType = "second"
	)

	var firstCalled bool
	var secondCalled bool

	events.Subscribe(firstEvent, func(any) {
		firstCalled = true
	})

	events.Subscribe(secondEvent, func(any) {
		secondCalled = true
	})

	events.Post(firstEvent, nil)

	if !firstCalled {
		t.Fatal("first event not called")
	} else if secondCalled {
		t.Fatal("second event should not have been called")
	}
}

func TestNilPayload(t *testing.T) {
	const eventType events.EventType = "nil"

	var received any

	events.Subscribe(eventType, func(data any) {
		received = data
	})

	events.Post(eventType, nil)

	if received != nil {
		t.Fatal("expected nil payload")
	}
}

func TestSubscriberOrder(t *testing.T) {
	const eventType events.EventType = "order"

	var order []int

	events.Subscribe(eventType, func(any) {
		order = append(order, 1)
	})

	events.Subscribe(eventType, func(any) {
		order = append(order, 2)
	})

	events.Subscribe(eventType, func(any) {
		order = append(order, 3)
	})

	events.Post(eventType, nil)

	if len(order) != 3 {
		t.Fatal("incorrect number of callbacks")
	}

	var expected []int = []int{1, 2, 3}
	var i int
	for i = range expected {
		if order[i] != expected[i] {
			t.Fatalf("callbacks executed out of order")
		}
	}
}

func TestPostManyTimes(t *testing.T) {
	const eventType events.EventType = "repeat"

	var count int

	events.Subscribe(eventType, func(any) {
		count++
	})

	for range 10 {
		events.Post(eventType, nil)
	}

	if count != 10 {
		t.Fatalf("expected 10 calls, got %d", count)
	}
}

func TestConcurrentPost(t *testing.T) {
	const eventType events.EventType = "concurrent"

	var mutex sync.Mutex
	var count int

	events.Subscribe(eventType, func(any) {
		mutex.Lock()
		count++
		mutex.Unlock()
	})

	const workers = 32

	var wait sync.WaitGroup
	wait.Add(workers)
	for range workers {
		go func() {
			defer wait.Done()
			events.Post(eventType, nil)
		}()
	}

	wait.Wait()

	if count != workers {
		t.Fatalf("expected %d calls, got %d", workers, count)
	}
}
