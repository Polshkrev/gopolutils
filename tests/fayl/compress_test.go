package tests

import (
	"bytes"
	"sync"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
)

func TestGzip(t *testing.T) {
	var input []byte = []byte("Hello, world!")
	var compressed []byte
	var exception *gopolutils.Exception
	compressed, exception = fayl.Gzip(input)
	if exception != nil {
		t.Fatal(exception)
	} else if len(compressed) == 0 {
		t.Fatal("expected compressed data")
	} else if bytes.Equal(input, compressed) {
		t.Fatal("expected compressed data to differ from input")
	}
}

func TestGunzip(t *testing.T) {
	var input []byte = []byte("Hello, world!")
	var compressed []byte
	var exception *gopolutils.Exception
	compressed, exception = fayl.Gzip(input)
	if exception != nil {
		t.Fatal(exception)
	}
	var output []byte
	output, exception = fayl.Gunzip(compressed)
	if exception != nil {
		t.Fatal(exception)
	} else if !bytes.Equal(input, output) {
		t.Fatalf("expected %q, got %q", input, output)
	}
}

func TestGzipGunzipRoundTrip(t *testing.T) {
	var input []byte = []byte("The quick brown fox jumps over the lazy dog")
	var compressed []byte
	var exception *gopolutils.Exception
	compressed, exception = fayl.Gzip(input)
	if exception != nil {
		t.Fatal(exception)
	}
	var output []byte
	output, exception = fayl.Gunzip(compressed)
	if exception != nil {
		t.Fatal(exception)
	} else if !bytes.Equal(input, output) {
		t.Fatal("round-trip failed")
	}
}

func TestGzipEmpty(t *testing.T) {
	var compressed []byte
	var exception *gopolutils.Exception
	compressed, exception = fayl.Gzip(nil)
	if exception != nil {
		t.Fatal(exception)
	}
	var output []byte
	output, exception = fayl.Gunzip(compressed)
	if exception != nil {
		t.Fatal(exception)
	} else if len(output) != 0 {
		t.Fatalf("expected empty output, got %d bytes", len(output))
	}
}

func TestGzipLarge(t *testing.T) {
	var input []byte = bytes.Repeat([]byte("abcdefghijklmnopqrstuvwxyz"), 10000)
	var compressed []byte
	var exception *gopolutils.Exception
	compressed, exception = fayl.Gzip(input)
	if exception != nil {
		t.Fatal(exception)
	}
	var output []byte
	output, exception = fayl.Gunzip(compressed)
	if exception != nil {
		t.Fatal(exception)
	} else if !bytes.Equal(input, output) {
		t.Fatal("large round-trip failed")
	}
}

func TestGzipBinaryData(t *testing.T) {
	var input []byte = make([]byte, 256)
	var i int
	for i = range input {
		input[i] = byte(i)
	}
	var compressed []byte
	var exception *gopolutils.Exception
	compressed, exception = fayl.Gzip(input)
	if exception != nil {
		t.Fatal(exception)
	}
	var output []byte
	output, exception = fayl.Gunzip(compressed)
	if exception != nil {
		t.Fatal(exception)
	} else if !bytes.Equal(input, output) {
		t.Fatal("binary round-trip failed")
	}
}

func TestGunzipInvalidData(t *testing.T) {
	var exception *gopolutils.Exception
	_, exception = fayl.Gunzip([]byte("this is not gzip"))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestGunzipNilInput(t *testing.T) {
	var exception *gopolutils.Exception
	_, exception = fayl.Gunzip(nil)

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestGunzipTruncatedData(t *testing.T) {
	var input []byte = []byte("Hello, world!")
	var compressed []byte
	var exception *gopolutils.Exception
	compressed, exception = fayl.Gzip(input)
	if exception != nil {
		t.Fatal(exception)
	}

	compressed = compressed[:len(compressed)/2]

	_, exception = fayl.Gunzip(compressed)

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestGzipRepeated(t *testing.T) {
	var input []byte = []byte("Repeated compression test")
	var i int
	for i = range 100 {
		var compressed []byte
		var exception *gopolutils.Exception
		compressed, exception = fayl.Gzip(input)
		if exception != nil {
			t.Fatal(exception)
		}
		var output []byte
		output, exception = fayl.Gunzip(compressed)
		if exception != nil {
			t.Fatal(exception)
		} else if !bytes.Equal(input, output) {
			t.Fatalf("iteration %d failed", i)
		}
	}
}

func TestConcurrentGzip(t *testing.T) {
	const workers int = 50

	var input []byte = []byte("Concurrent gzip test")

	var wait sync.WaitGroup
	wait.Add(workers)
	for range workers {
		go func() {
			defer wait.Done()
			var compressed []byte
			var exception *gopolutils.Exception
			compressed, exception = fayl.Gzip(input)
			if exception != nil {
				t.Error(exception)
				return
			}
			var output []byte
			output, exception = fayl.Gunzip(compressed)
			if exception != nil {
				t.Error(exception)
				return
			} else if !bytes.Equal(input, output) {
				t.Error("round-trip failed")
			}
		}()
	}

	wait.Wait()
}
