package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
)

func TestNewException(t *testing.T) {
	t.Parallel()

	var exception *gopolutils.Exception = gopolutils.NewException("hello %s", "world")

	var expected string = "Exception: hello world"
	if exception == nil {
		t.Fatal("Expected exception, got nil")
	} else if exception.Name() != gopolutils.BaseException {
		t.Fatalf("Expected name %q, got %q", gopolutils.BaseException, exception.Name())
	} else if exception.Message() != "hello world" {
		t.Fatalf("Expected message %q, got %q", "hello world", exception.Message())
	} else if exception.Error() != expected {
		t.Fatalf("Expected error %q, got %q", expected, exception.Error())
	}
}

func TestNewNamedException(t *testing.T) {
	t.Parallel()

	var name gopolutils.ExceptionName = "MyException"

	var exception *gopolutils.Exception = gopolutils.NewNamedException(name, "Value %d", 42)

	var expected string = "MyException: Value 42"

	if exception == nil {
		t.Fatal("Expected exception, got nil")
	} else if exception.Name() != name {
		t.Fatalf("Expected name %q, got %q", name, exception.Name())
	} else if exception.Message() != "Value 42" {
		t.Fatalf("Expected message %q, got %q", "Value 42", exception.Message())
	} else if exception.Error() != expected {
		t.Fatalf("Expected error %q, got %q", expected, exception.Error())
	}
}

type isTest struct {
	name     string
	except   *gopolutils.Exception
	match    gopolutils.ExceptionName
	expected bool
}

func TestExceptionIs(t *testing.T) {
	t.Parallel()

	var tests []isTest = []isTest{
		{
			name:     "matching",
			except:   gopolutils.NewNamedException("Custom", "message"),
			match:    "Custom",
			expected: true,
		},
		{
			name:     "not matching",
			except:   gopolutils.NewNamedException("Custom", "message"),
			match:    "Other",
			expected: false,
		},
		{
			name:     "Exception",
			except:   gopolutils.NewException("message"),
			match:    gopolutils.BaseException,
			expected: true,
		},
	}
	var i int
	for i = range tests {
		var test isTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			var result bool = test.except.Is(test.match)
			if result != test.expected {
				t.Fatalf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestMustReturnsResult(t *testing.T) {
	t.Parallel()

	var value int = gopolutils.Must(123, nil)

	if value != 123 {
		t.Fatalf("Expected 123, got %d", value)
	}
}

func TestMustPanics(t *testing.T) {
	t.Parallel()

	var exception *gopolutils.Exception = gopolutils.NewException("Failure")

	defer func() {
		var recovered any = recover()

		var message string
		var ok bool
		message, ok = recovered.(string)

		if recovered == nil {
			t.Fatal("expected panic")
		} else if !ok {
			t.Fatalf("expected panic string, got %T", recovered)
		} else if message != exception.Error() {
			t.Fatalf("expected panic %q, got %q", exception.Error(), message)
		}
	}()

	var _ int = gopolutils.Must(0, exception)
}

type errorTest struct {
	name     string
	except   *gopolutils.Exception
	expected string
}

func TestExceptionError(t *testing.T) {
	t.Parallel()
	var runtimeError gopolutils.ExceptionName = gopolutils.RuntimeError
	var tests []errorTest = []errorTest{
		{
			name:     "base",
			except:   gopolutils.NewException("abc"),
			expected: "Exception: abc",
		},
		{
			name:     "named",
			except:   gopolutils.NewNamedException(runtimeError, "xyz"),
			expected: "RuntimeError: xyz",
		},
	}

	var i int
	for i = range tests {
		var test errorTest = tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			var result string = test.except.Error()
			if result != test.expected {
				t.Fatalf("Expected %q, got %q", test.expected, result)
			}
		})
	}
}

func TestMessage(t *testing.T) {
	t.Parallel()

	var exception *gopolutils.Exception = gopolutils.NewException("message")

	if exception.Message() != "message" {
		t.Fatalf("Expected %q, got %q", "message", exception.Message())
	}
}

func TestName(t *testing.T) {
	t.Parallel()

	var name gopolutils.ExceptionName = gopolutils.ValueError

	var exception *gopolutils.Exception = gopolutils.NewNamedException(name, "bad value")

	if exception.Name() != name {
		t.Fatalf("Expected %q, got %q", name, exception.Name())
	}
}
