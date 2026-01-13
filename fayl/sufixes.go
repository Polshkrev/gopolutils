package fayl

import (
	"fmt"
	"sync"

	"github.com/Polshkrev/gopolutils"
)

// Representation of a file suffix i.e. ".go", ".c", ".h", et cetera.
type Suffix = gopolutils.Enum

const (
	A Suffix = iota
	Asm
	Bat
	C
	Class
	Cpp
	Cs
	Css
	Csv
	Db
	Dll
	Exe
	Go
	Header
	Html
	Jar
	Java
	Js
	Json
	Lib
	Log
	Md
	None
	O
	Py
	Rs
	Sh
	So
	Sql
	Tex
	Toml
	Txt
	Xml
	Yaml
	Zip
)

const (
	// Total amount of supported file suffixes.
	SuffixCount uint8 = Zip
)

var (
	stringLock sync.Mutex
	// Mapping of the string representation of the file suffix to its corresponding enum value. This mapping is not concurrent safe.
	stringToSuffix map[string]Suffix = map[string]Suffix{
		"7z":    Zip,
		"a":     A,
		"asm":   Asm,
		"bat":   Bat,
		"c":     C,
		"class": Class,
		"cpp":   Cpp,
		"cs":    Cs,
		"css":   Css,
		"csv":   Csv,
		"db":    Db,
		"dll":   Dll,
		"exe":   Exe,
		"go":    Go,
		"h":     Header,
		"hpp":   Header,
		"html":  Html,
		"jar":   Jar,
		"java":  Java,
		"js":    Js,
		"json":  Json,
		"lib":   Lib,
		"log":   Log,
		"md":    Md,
		"o":     O,
		"py":    Py,
		"rar":   Zip,
		"rs":    Rs,
		"sh":    Sh,
		"so":    So,
		"sql":   Sql,
		"tex":   Tex,
		"toml":  Toml,
		"txt":   Txt,
		"xml":   Xml,
		"yaml":  Yaml,
		"yml":   Yaml,
		"zip":   Zip,
	}
)

var (
	// Concurrent safe lock from mapping [suffixToString]
	suffixLock sync.Mutex
	// Mapping of [Suffix] enum values to their corresponding string value.
	suffixToString map[Suffix]string = map[Suffix]string{
		A:      "a",
		Asm:    "asm",
		Bat:    "bat",
		C:      "c",
		Class:  "class",
		Cpp:    "cpp",
		Cs:     "cs",
		Css:    "css",
		Csv:    "csv",
		Db:     "db",
		Dll:    "dll",
		Exe:    "exe",
		Go:     "go",
		Header: "h",
		Html:   "html",
		Jar:    "jar",
		Java:   "java",
		Js:     "js",
		Json:   "json",
		Lib:    "lib",
		Log:    "log",
		Md:     "md",
		None:   "",
		O:      "o",
		Py:     "py",
		Rs:     "rs",
		Sh:     "sh",
		So:     "so",
		Sql:    "sql",
		Tex:    "tex",
		Toml:   "toml",
		Txt:    "txt",
		Xml:    "xml",
		Yaml:   "yaml",
		Zip:    "zip",
	}
)

// Obtain a [Suffix] from a raw string.
// Returns a path suffix obtained from [stringToSuffix].
// If the path suffix is not defined in stringToSuffix, a `KeyError` is returned with the `None` suffix value.
func SuffixFromString(suffix string) (Suffix, *gopolutils.Exception) {
	stringLock.Lock()
	defer stringLock.Unlock()
	var ok bool
	var item Suffix
	item, ok = stringToSuffix[suffix]
	if !ok {
		return None, gopolutils.NewNamedException(gopolutils.KeyError, fmt.Sprintf("'%s' is not defined in mapping.", suffix))
	}
	return item, nil
}

// Obtain a string from a [Suffix].
// Returns a string obtained from [suffixToString].
// If the path suffix is not defined in suffixToString, a `KeyError` is returned with an empty string.
func StringFromSuffix(suffix Suffix) (string, *gopolutils.Exception) {
	suffixLock.Lock()
	defer suffixLock.Unlock()
	var ok bool
	var item string
	item, ok = suffixToString[suffix]
	if !ok {
		return "", gopolutils.NewNamedException(gopolutils.KeyError, fmt.Sprintf("'%s' is not defined in mapping.", suffix))
	}
	return item, nil
}
