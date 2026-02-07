/*
gopolutils provides numerious utilities standardizing logging, enum creation, exception handling, and internal version control.
The example below describes a "Hello World" programme using a [Logger] and [Exception].

Example:

	import "github.com/Polshkrev/gopolutils"

	func main() {
		var logger *gopolutils.Logger = gopolutils.NewLogger("main", gopolutils.Debug)
		var except *gopolutils.Exception = logger.ConsoleOnly()
		if except != nil {
			panic(except)
		}
		logger.Log("Hello World", gopolutils.Debug)

	}

A simple use case for the [Version] is release flags.

As an example:

	import "github.com/Polshkrev/gopolutils"

	var version *gopolutils.Version = gopolutils.VersionConvert(0, 1, 0)

	func main() {
		if !version.IsPublic() {
			// code/functionality to expose when the internal version becomes public.
		}
		panic("This is not the code you're looking for.")
	}

Or in a more verbose way:

	import (
		"fmt"
		"os"

		"github.com/Polshkrev/gopolutils"
	)

	var version *gopolutils.Version = gopolutils.VersionConvert(0, 1, 0)

	func main() {
		if !version.IsPublic() {
			// code/functionality to expose when the internal version becomes public.
		}
		fmt.Fprintln(os.Stderr, gopolutils.NewNamedException(gopolutils.ValueError, "All your code is belong to us."))
		os.Exit(1)
	}
*/
package gopolutils
