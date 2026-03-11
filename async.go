package gopolutils

import "runtime/debug"

// Python-like concurrent future.
type Future[Type any] struct {
	await func() (Type, *Exception)
}

// Create an async [Future] based on a given callback.
// If the child process fails, a [ChildProcessError] is returned.
func Async[Type any](caller func() (Type, *Exception)) *Future[Type] {
	var result Type
	var except *Exception
	var done chan struct{} = make(chan struct{})
	go func() {
		defer func() {
			var result any
			if result = recover(); result != nil {
				switch x := result.(type) { // I don't like this syntax.
				case error:
					except = NewNamedException(ChildProcessError, "Panic in async worker: %w\n%s", x, string(debug.Stack()))
				default:
					except = NewNamedException(ChildProcessError, "Panic in async worker: %v\n%s", x, string(debug.Stack()))
				}
			}
			close(done)
		}()
		result, except = caller()
	}()
	var future *Future[Type] = new(Future[Type])
	future.await = func() (Type, *Exception) {
		<-done
		return result, except
	}
	return future
}

// Await a [Future].
// The exception state is dependant on the [Future]'s child process.
func (future *Future[Type]) Await() (Type, *Exception) {
	return future.await()
}
