package gopolutils

import "runtime/debug"

// Python-like concurrent future.
type Future[Type any] struct {
	await func() (Type, *Exception)
}

// Await a [Future].
// Returns the result of the awaited function within the [Future].
func (future *Future[Type]) Await() (Type, *Exception) {
	return future.await()
}

// Construct a [Future] based on a given callback.
// Returns a new [Future] to be awaited based on a given callback.
func Async[Type any](callback func() (Type, *Exception)) *Future[Type] {
	var result Type
	var except *Exception
	var done chan struct{} = make(chan struct{})
	go func() {
		defer func() {
			var result any
			if result = recover(); result != nil {
				switch x := result.(type) {
				case *Exception:
					except = NewNamedException(ChildProcessError, "Panic in asynchronous worker: %w\n%s", x, string(debug.Stack()))
				default:
					except = NewNamedException(ChildProcessError, "Panic in asynchronous worker: %v\n%s", x, string(debug.Stack()))
				}
			}
			close(done)
		}()
		result, except = callback()
	}()
	var future *Future[Type] = new(Future[Type])
	future.await = func() (Type, *Exception) {
		<-done
		return result, except
	}
	return future
}
