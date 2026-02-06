package gopolutils

// Finite list of named exceptions. Inspired by python's built-in exceptions.
type ExceptionName = StringEnum

const (
	BaseException          ExceptionName = "Exception"
	ArithmeticError        ExceptionName = "ArithmeticError"        // For any miscellaneous arithmatic exception.
	OverflowError          ExceptionName = "OverflowError"          // For any overflow.
	UnderflowError         ExceptionName = "UnderflowError"         // For any underflow.
	ZeroDivisionError      ExceptionName = "ZeroDivisionError"      // When dividing by zero.
	AssertionError         ExceptionName = "AssertionError"         // For any failed assertion.
	EOFError               ExceptionName = "EOFError"               // When reaching the end of a file or buffer.
	LookupError            ExceptionName = "LookupError"            // When any miscellaneous lookup fails.
	OutOfRangeError        ExceptionName = "OutOfRangeError"        // For any indexed access outside of the allotted range.
	IndexError             ExceptionName = "IndexError"             // For any indexed access.
	KeyError               ExceptionName = "KeyError"               // For any keyed access.
	OSError                ExceptionName = "OSError"                // For any miscellaneous operating system exception.
	IOError                ExceptionName = "IOError"                // For any miscellaneous io exceptions.
	BlockingIOError        ExceptionName = "BlockingIOError"        // For any miscellaneous blocking io exception.
	ChildProcessError      ExceptionName = "ChildProcessError"      // For any miscellaneous child process exception.
	ConnectionError        ExceptionName = "ConnectionError"        // For any miscellaneous connection exception.
	BrokenPipeError        ExceptionName = "BrokenPipeError"        // For any exception relating to a broken pipe.
	ConnectionAbortedError ExceptionName = "ConnectionAbortedError" // For an aborted connection.
	ConnectionRefusedError ExceptionName = "ConnectionRefusedError" // For a refused connection.
	ConnectionResetError   ExceptionName = "ConnectionResetError"   // For a reset connection.
	FileExistsError        ExceptionName = "FileExistsError"        // When a file exists already in the file system.
	FileNotFoundError      ExceptionName = "FileNotFoundError"      // When a file cannot be found.
	IsADirectoryError      ExceptionName = "IsADirectoryError"      // When an accessed file is a directory.
	NotADirectoryError     ExceptionName = "NotADirectoryError"     // When an accessed directory is not a directory.
	PermissionError        ExceptionName = "PermissionError"        // When an accessed file or endpoint does not have the correct permissions
	ProcessLookupError     ExceptionName = "ProcessLookupError"     // For a porcess lookup fail.
	TimeoutError           ExceptionName = "TimeoutError"           // For a timeout occurrence.
	RuntimeError           ExceptionName = "RuntimeError"           // For any miscellaneous runtime exception.
	NotImplementedError    ExceptionName = "NotImplementedError"    // For any unimplemented methods or functions.
	ValueError             ExceptionName = "ValueError"             // For any miscellaneous value exception.
	UnreachableError       ExceptionName = "UnreachableError"       // For an unreachable case.
)
