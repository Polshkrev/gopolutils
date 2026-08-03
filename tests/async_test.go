package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/Polshkrev/gopolutils"
)

func TestAsyncReturnsValue(t *testing.T) {
	var future *gopolutils.Future[int] = gopolutils.Async(func() (int, *gopolutils.Exception) {
		return 42, nil
	})

	var value int
	var except *gopolutils.Exception
	value, except = future.Await()

	if except != nil {
		t.Fatalf("Expected nil exception, got %v", except)
	} else if value != 42 {
		t.Fatalf("Expected 42, got %d", value)
	}
}

func TestAsyncReturnsException(t *testing.T) {
	var expected *gopolutils.Exception = gopolutils.NewException("failure")

	var future *gopolutils.Future[int] = gopolutils.Async(func() (int, *gopolutils.Exception) {
		return 0, expected
	})

	var value int
	var except *gopolutils.Exception
	value, except = future.Await()

	if value != 0 {
		t.Fatalf("Expected zero value, got %d", value)
	} else if except != expected {
		t.Fatalf("Expected returned exception")
	}
}

func TestAwaitBlocksUntilFinished(t *testing.T) {
	var future *gopolutils.Future[string] = gopolutils.Async(func() (string, *gopolutils.Exception) {
		time.Sleep(100 * time.Millisecond)
		return "done", nil
	})

	var start time.Time = time.Now()

	var value string
	var except *gopolutils.Exception
	value, except = future.Await()

	var elapsed time.Duration = time.Since(start)

	if except != nil {
		t.Fatalf("unexpected exception: %v", except)
	} else if value != "done" {
		t.Fatalf("expected done, got %q", value)
	} else if elapsed < 100*time.Millisecond {
		t.Fatalf("Await returned too early")
	}
}

func TestFutureCanBeAwaitedMultipleTimes(t *testing.T) {
	var future *gopolutils.Future[int] = gopolutils.Async(func() (int, *gopolutils.Exception) {
		return 123, nil
	})

	var valueOne int
	var exceptOne *gopolutils.Exception
	valueOne, exceptOne = future.Await()
	var valueTwo int
	var exceptTwo *gopolutils.Exception
	valueTwo, exceptTwo = future.Await()

	if exceptOne != nil || exceptTwo != nil {
		t.Fatalf("Unexpected exception")
	} else if valueOne != 123 || valueTwo != 123 {
		t.Fatalf("Expected repeated result")
	}
}

func TestAsyncPanicWithException(t *testing.T) {
	expected := gopolutils.NewException("boom")

	var future *gopolutils.Future[int] = gopolutils.Async(func() (int, *gopolutils.Exception) {
		panic(expected)
	})

	var except *gopolutils.Exception
	_, except = future.Await()

	if except == nil {
		t.Fatal("expected exception")
	}

	if except.Name() != gopolutils.ChildProcessError {
		t.Fatalf("expected %q, got %q", gopolutils.ChildProcessError, except.Name())
	}
}

func TestAsyncPanicWithString(t *testing.T) {
	var future *gopolutils.Future[int] = gopolutils.Async(func() (int, *gopolutils.Exception) {
		panic("boom")
	})

	var except *gopolutils.Exception
	_, except = future.Await()

	if except == nil {
		t.Fatal("expected exception")
	} else if except.Name() != gopolutils.ChildProcessError {
		t.Fatalf("expected %q, got %q", gopolutils.ChildProcessError, except.Name())
	}
}

func TestAsyncPanicWithError(t *testing.T) {
	var expected error = errors.New("failure")

	var future *gopolutils.Future[int] = gopolutils.Async(func() (int, *gopolutils.Exception) {
		panic(expected)
	})

	var except *gopolutils.Exception
	_, except = future.Await()

	if except == nil {
		t.Fatal("Expected exception")
	} else if except.Name() != gopolutils.ChildProcessError {
		t.Fatalf("Expected %q, got %q", gopolutils.ChildProcessError, except.Name())
	}
}

func TestAsyncReferenceType(t *testing.T) {
	type data struct {
		value int
	}

	var future *gopolutils.Future[*data] = gopolutils.Async(func() (*data, *gopolutils.Exception) {
		return &data{value: 99}, nil
	})

	var value *data
	var except *gopolutils.Exception
	value, except = future.Await()

	if except != nil {
		t.Fatalf("Unexpected exception: %v", except)
	} else if value == nil {
		t.Fatal("Expected non-nil pointer")
	} else if value.value != 99 {
		t.Fatalf("Expected 99, got %d", value.value)
	}
}

func TestAsyncZeroValueOnPanic(t *testing.T) {
	var future *gopolutils.Future[int] = gopolutils.Async(func() (int, *gopolutils.Exception) {
		panic("failure")
	})

	var value int
	var except *gopolutils.Exception
	value, except = future.Await()

	if except == nil {
		t.Fatal("Expected exception")
	} else if value != 0 {
		t.Fatalf("Expected zero value, got %d", value)
	}
}

func TestConcurrentAwait(t *testing.T) {
	var future *gopolutils.Future[int] = gopolutils.Async(func() (int, *gopolutils.Exception) {
		time.Sleep(50 * time.Millisecond)
		return 7, nil
	})
	var n int = 10
	var done chan gopolutils.None = make(chan gopolutils.None)
	for range n {
		go func() {
			var value int
			var except *gopolutils.Exception
			value, except = future.Await()

			if except != nil {
				t.Errorf("unexpected exception: %v", except)
			} else if value != 7 {
				t.Errorf("expected 7, got %d", value)
			}

			done <- gopolutils.None{}
		}()
	}

	for range n {
		<-done
	}
}
