package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Polshkrev/gopolutils"
)

func TestNewLogger(t *testing.T) {
	var name string = "Test"
	var level gopolutils.LoggingLevel = gopolutils.Debug
	var logger *gopolutils.Logger = gopolutils.NewLogger(name, level)
	var nameResult string = logger.GetName()
	var levelResult gopolutils.LoggingLevel = logger.GetLevel()
	if logger == nil {
		t.Fatal("NewLogger returned nil")
	} else if !strings.EqualFold(nameResult, name) {
		t.Fatalf("Expected name %q, got %q", name, nameResult)
	} else if levelResult != level {
		t.Fatalf("Expected level %v, got %v", level, levelResult)
	}
}

func TestSetName(t *testing.T) {
	var oldName string = "old"
	var newName string = "new"
	var logger *gopolutils.Logger = gopolutils.NewLogger(oldName, gopolutils.Debug)
	logger.SetName(newName)
	var nameResult string = logger.GetName()
	if nameResult != newName {
		t.Fatalf("Expected name %q, got %q", newName, nameResult)
	}
}

func TestSetLevel(t *testing.T) {
	logger := gopolutils.NewLogger("logger", gopolutils.Debug)

	logger.SetLevel(gopolutils.Error)

	if logger.GetLevel() != gopolutils.Error {
		t.Fatalf("expected level %v, got %v", gopolutils.Error, logger.GetLevel())
	}
}

type levelTest struct {
	name     string
	level    gopolutils.LoggingLevel
	expected string
}

func TestLoggingLevelString(t *testing.T) {
	var tests []levelTest = []levelTest{
		{"Debug", gopolutils.Debug, "DEBUG"},
		{"Info", gopolutils.Info, "INFO"},
		{"Warning", gopolutils.Warning, "WARNING"},
		{"Error", gopolutils.Error, "ERROR"},
		{"Critical", gopolutils.Critical, "CRITICAL"},
	}
	var i int
	for i = range tests {
		var test levelTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			var actual string = test.level.String()
			if actual != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, actual)
			}
		})
	}
}

func TestAddConsoleFailsWhenOutputLimitExceeded(t *testing.T) {
	var loggerName string = "logger"
	var logger *gopolutils.Logger = gopolutils.NewLogger(loggerName, gopolutils.Debug)

	var except *gopolutils.Exception = logger.ConsoleOnly()
	if except != nil {
		t.Fatalf("Unexpected error: %v", except)
	}
	except = logger.AddConsole()
	if except == nil {
		t.Fatal("Expected AddConsole() to fail at capacity.")
	}
}

func TestAddFileFailsWhenOutputLimitExceeded(t *testing.T) {
	var temporaryDirectory string = t.TempDir()
	var path string = filepath.Join(temporaryDirectory, "test.log")

	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	var except *gopolutils.Exception = logger.ConsoleOnly()
	if except != nil {
		t.Fatalf("Unexpected error: %v", except)
	}
	except = logger.AddFile(path)

	if except == nil {
		t.Fatal("Expected AddFile() to fail at capacity.")
	}
	defer logger.Close()
}

func TestAddFileFailsForInvalidPath(t *testing.T) {
	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)

	var except *gopolutils.Exception = logger.AddFile(filepath.Join("", "\x00"))

	if except == nil {
		t.Fatal("Expected AddFile to return an error")
	}
}

func TestAddFileCreatesFile(t *testing.T) {
	var temporaryDirectory string = t.TempDir()
	var path string = filepath.Join(temporaryDirectory, "test.log")
	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	var except *gopolutils.Exception = logger.AddFile(path)
	if except != nil {
		t.Fatalf("AddFile() returned %v", except)
	}
	var statError error
	_, statError = os.Stat(path)
	if statError != nil {
		t.Fatalf("expected log file to exist: %v", statError)
	}
	logger.Close()
}

func TestFullSetup(t *testing.T) {
	var filename string = "test.log"
	var loggerName string = "logger"
	var tempDirectory string = t.TempDir()
	var path string = filepath.Join(tempDirectory, filename)
	var logger *gopolutils.Logger = gopolutils.NewLogger(loggerName, gopolutils.Debug)
	var except *gopolutils.Exception = logger.FullSetup(path)
	if except != nil {
		t.Fatalf("FullSetup() returned %v", except)
	}
	logger.Close()
}

func TestLogIgnoresNilOutputs(t *testing.T) {
	var temporaryDirectory string = t.TempDir()
	var path string = filepath.Join(temporaryDirectory, "test.log")

	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	var except *gopolutils.Exception = logger.AddFile(path)
	if except != nil {
		t.Fatalf("AddFile() returned %v", except)
	}
	defer logger.Close()

	defer func() {
		var recovery any = recover()
		if recovery != nil {
			t.Fatalf("Log() panicked with a nil output: %v", recovery)
		}
	}()

	logger.Log("Hello, World", gopolutils.Info)
}

func TestLogfDoesNotPanic(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.log")

	logger := gopolutils.NewLogger("logger", gopolutils.Debug)

	if err := logger.AddFile(path); err != nil {
		t.Fatal(err)
	}
	defer logger.Close()

	defer func() {
		var recovery any = recover()
		if recovery != nil {
			t.Fatalf("Log() panicked: %v\n", recovery)
		}
	}()

	logger.Logf(gopolutils.Info, "Value = %d", 42)
}

func TestConsoleOnly(t *testing.T) {
	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	var except *gopolutils.Exception = logger.ConsoleOnly()
	if except != nil {
		t.Fatalf("ConsoleOnly() returned %v", except)
	}
}

func TestFileOnly(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "test.log")

	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	var except *gopolutils.Exception = logger.FileOnly(path)
	if except != nil {
		t.Fatalf("FileOnly() returned %v", except)
	}

	logger.Close()
}

func TestLoggerName(t *testing.T) {
	var firstName string = "one"
	var secondName string = "two"
	var logger *gopolutils.Logger = gopolutils.NewLogger(firstName, gopolutils.Debug)

	logger.SetName(secondName)
	var result string = logger.GetName()
	if result != secondName {
		t.Fatalf("Expected %q, got %q", secondName, result)
	}
}

func TestLoggerLevel(t *testing.T) {
	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	logger.SetLevel(gopolutils.Error)
	var result gopolutils.LoggingLevel = logger.GetLevel()
	if result != gopolutils.Error {
		t.Fatalf("expected %v, got %v", gopolutils.Error, result)
	}
}

func TestMultipleLoggersAreIndependent(t *testing.T) {
	var firstLogger *gopolutils.Logger = gopolutils.NewLogger("a", gopolutils.Debug)
	var secondLogger *gopolutils.Logger = gopolutils.NewLogger("b", gopolutils.Debug)

	var except *gopolutils.Exception = firstLogger.AddConsole()
	if except != nil {
		t.Fatal(except)
	}
	except = secondLogger.AddConsole()
	if except != nil {
		t.Fatal(except)
	}
	except = secondLogger.AddConsole()
	if except != nil {
		t.Fatal(except)
	}
}

func readFile(path string) string {
	var data []byte
	var readError error
	data, readError = os.ReadFile(path)
	if readError != nil {
		panic(gopolutils.NewNamedException(gopolutils.IOError, readError.Error()))
	}
	return string(data)
}

func TestLogWritesMessage(t *testing.T) {
	var temporaryDirectory string = t.TempDir()
	var path string = filepath.Join(temporaryDirectory, "test.log")

	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	var except *gopolutils.Exception = logger.AddFile(path)
	if except != nil {
		t.Fatal(except)
	}

	logger.Log("Hello, world!", gopolutils.Info)
	defer logger.Close()

	var output string = readFile(path)

	if !strings.Contains(output, "Hello, world!") {
		t.Fatalf("Expected output to contain message, got %q", output)
	} else if !strings.Contains(output, gopolutils.Info.String()) {
		t.Fatalf("Expected output to contain logging level, got %q", output)
	} else if !strings.Contains(output, "logger") {
		t.Fatalf("Expected output to contain logger name, got %q", output)
	}
}

func TestLogfWritesFormattedMessage(t *testing.T) {
	var temporaryDirectory string = t.TempDir()
	var path string = filepath.Join(temporaryDirectory, "test.log")

	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	var except *gopolutils.Exception = logger.AddFile(path)
	if except != nil {
		t.Fatal(except)
	}

	logger.Logf(gopolutils.Info, "value = %d", 42)

	defer logger.Close()

	var data string = readFile(path)

	if !strings.Contains(data, "value = 42") {
		t.Fatalf("unexpected log output: %q", data)
	}
}

func TestLogRespectsLoggingLevel(t *testing.T) {
	var temporaryDirectory string = t.TempDir()
	var path string = filepath.Join(temporaryDirectory, "test.log")

	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Warning)
	var except *gopolutils.Exception = logger.AddFile(path)
	if except != nil {
		t.Fatal(except)
	}

	logger.Log("debug", gopolutils.Debug)
	logger.Log("info", gopolutils.Info)
	logger.Log("warning", gopolutils.Warning)
	defer logger.Close()

	var output string = readFile(path)

	if strings.Contains(output, "debug") {
		t.Fatal("debug message should not have been logged")
	} else if strings.Contains(output, "info") {
		t.Fatal("info message should not have been logged")
	} else if !strings.Contains(output, "warning") {
		t.Fatal("warning message should have been logged")
	}
}

func TestLogEmptyMessage(t *testing.T) {
	var temporaryDirectory string = t.TempDir()
	var path string = filepath.Join(temporaryDirectory, "test.log")

	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	var except *gopolutils.Exception = logger.AddFile(path)
	if except != nil {
		t.Fatal(except)
	}

	logger.Log("", gopolutils.Info)
	defer logger.Close()

	if len(readFile(path)) == 0 {
		t.Fatal("Expected log output.")
	}
}

func TestLogMultipleMessages(t *testing.T) {
	var messages []string = []string{"one", "two", "three"}
	var temporaryDirectory string = t.TempDir()
	var path string = filepath.Join(temporaryDirectory, "test.log")

	var logger *gopolutils.Logger = gopolutils.NewLogger("logger", gopolutils.Debug)
	var except *gopolutils.Exception = logger.AddFile(path)
	if except != nil {
		t.Fatal(except)
	}

	logger.Log(messages[0], gopolutils.Info)
	logger.Log(messages[1], gopolutils.Warning)
	logger.Log(messages[2], gopolutils.Error)
	defer logger.Close()

	var output string = readFile(path)
	var i int
	for i = range messages {
		var message string = messages[i]
		if !strings.Contains(output, message) {
			t.Fatalf("Expected output to contain %q.", message)
		}
	}
}
