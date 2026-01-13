package fayl

import "github.com/Polshkrev/gopolutils"

// Representation of a file suffix i.e. ".go", ".c", ".h", et cetera.
type Suffix = gopolutils.Enum

const (
	ASM Suffix = iota
	BAT
	C
	CLASS
	CPP
	CS
	CSS
	CSV
	DB
	DLL
	EXE
	GO
	HEADER
	HTML
	JAR
	JAVA
	JS
	JSON
	LIB
	LOG
	MD
	O
	PY
	RS
	SH
	SO
	SQL
	TEX
	TOML
	TXT
	XML
	YAML
	ZIP
)

const (
	// Total amount of supported file suffixes.
	SuffixCount uint8 = ZIP
)

var (
	// Mapping of the string representation of the file suffix to its corresponding enum value. This mapping is not concurrent safe.
	StringToSuffix map[string]Suffix = map[string]Suffix{
		"7z":    ZIP,
		"asm":   ASM,
		"bat":   BAT,
		"c":     C,
		"class": CLASS,
		"cpp":   CPP,
		"cs":    CS,
		"css":   CSS,
		"csv":   CSV,
		"db":    DB,
		"dll":   DLL,
		"exe":   EXE,
		"go":    GO,
		"h":     HEADER,
		"hpp":   HEADER,
		"html":  HTML,
		"jar":   JAR,
		"java":  JAVA,
		"js":    JS,
		"json":  JSON,
		"lib":   LIB,
		"log":   LOG,
		"md":    MD,
		"o":     O,
		"py":    PY,
		"rar":   ZIP,
		"rs":    RS,
		"sh":    SH,
		"so":    SO,
		"sql":   SQL,
		"tex":   TEX,
		"toml":  TOML,
		"txt":   TXT,
		"xml":   XML,
		"yaml":  YAML,
		"zip":   ZIP,
	}
)

var (
	// Mapping of [Suffix] enum values to their corresponding string value.
	SuffixToString map[Suffix]string = map[Suffix]string{
		ASM:    "asm",
		BAT:    "bat",
		C:      "c",
		CLASS:  "class",
		CPP:    "cpp",
		CS:     "cs",
		CSS:    "css",
		CSV:    "csv",
		DB:     "db",
		DLL:    "dll",
		EXE:    "exe",
		GO:     "go",
		HEADER: "h",
		HTML:   "html",
		JAR:    "jar",
		JAVA:   "java",
		JS:     "js",
		JSON:   "json",
		LIB:    "lib",
		LOG:    "log",
		MD:     "md",
		O:      "o",
		PY:     "py",
		RS:     "rs",
		SH:     "sh",
		SO:     "so",
		SQL:    "sql",
		TEX:    "tex",
		TOML:   "toml",
		TXT:    "txt",
		XML:    "xml",
		YAML:   "yaml",
		ZIP:    "zip",
	}
)
