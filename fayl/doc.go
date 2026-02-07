/*
fayl provides numerious utilities pertaining to the creation, destruction, serialization, and manipulation of files and their corresponding paths on the filesystem.
The example below describes a marshalling of a single object given a constructed [Path] from its given folder, filename, and [Suffix].

Example:

	import (

		"github.com/Polshkrev/gopolutils"
		"github.com/Polshkrev/gopolutils/fayl"

	)

	func main() {
		var settings *settings.Settings // Just an example object; no settings package exists within this package.
		var except *gopolutils.Except
		settings, except = fayl.ReadObject(fayl.PathFromParts("./settings", "settings", fayl.Json))
		if except != nil {
			panic(except)
		}
	}

Or as a more succinct example:

	import (

		"github.com/Polshkrev/gopolutils"
		"github.com/Polshkrev/gopolutils/fayl"

	)

	func main() {
		var settings *settings.Settings = gopolutils.Must(fayl.ReadObject(fayl.PathFromParts("./settings", "settings", fayl.Json)))
	}

Each of the file operations are concurrent, although not explicitly concurrent-safe.
*/
package fayl
