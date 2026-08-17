package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
)

func TestPathFrom(t *testing.T) {
	const raw string = "foo/bar"

	var path *fayl.Path = fayl.PathFrom(raw)

	if path.String() != raw {
		t.Fatalf("expected %q, got %q", raw, path.String())
	}
}

func TestPathFromParts(t *testing.T) {
	var path *fayl.Path = fayl.PathFromParts("src", "main", fayl.Go)

	var expected string = filepath.Join("src", "main.go")

	if path.String() != expected {
		t.Fatalf("expected %q, got %q", expected, path.String())
	}
}

func TestExists(t *testing.T) {
	var dir string = t.TempDir()

	if !fayl.PathFrom(dir).Exists() {
		t.Fatal("expected temporary directory to exist")
	} else if fayl.PathFrom(filepath.Join(dir, "does_not_exist")).Exists() {
		t.Fatal("expected nonexistent path not to exist")
	}
}

func TestAbsolute(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom(".")
	var absolute *fayl.Path
	var exception *gopolutils.Exception
	absolute, exception = path.Absolute()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	}
	var expected string
	var absoluteError error
	expected, absoluteError = filepath.Abs(".")
	if absoluteError != nil {
		t.Fatal(absoluteError)
	} else if absolute.String() != expected {
		t.Fatalf("expected %q, got %q", expected, absolute.String())
	}
}

func TestAppend(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom("foo")

	path.Append(fayl.PathFrom("bar"))

	var expected string = filepath.Join("foo", "bar")

	if path.String() != expected {
		t.Fatalf("expected %q, got %q", expected, path.String())
	}
}

func TestAppendAs(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom("foo")

	path.AppendAs("bar")

	var expected string = filepath.Join("foo", "bar")

	if path.String() != expected {
		t.Fatalf("expected %q, got %q", expected, path.String())
	}
}

func TestJoin(t *testing.T) {
	var original *fayl.Path = fayl.PathFrom("foo")

	var result *fayl.Path = original.Join(*fayl.PathFrom("bar"))

	var expected string = filepath.Join("foo", "bar")

	if result.String() != expected {
		t.Fatalf("expected %q, got %q", expected, result.String())
	} else if original.String() != "foo" {
		t.Fatal("Join modified the original path")
	}
}

func TestJoinAs(t *testing.T) {
	var original *fayl.Path = fayl.PathFrom("foo")

	var result *fayl.Path = original.JoinAs("bar")

	var expected string = filepath.Join("foo", "bar")

	if result.String() != expected {
		t.Fatalf("expected %q, got %q", expected, result.String())
	} else if original.String() != "foo" {
		t.Fatal("JoinAs modified the original path")
	}
}

type pathSuffixTest struct {
	name     string
	path     string
	expected fayl.Suffix
	hasError bool
}

func TestSuffix(t *testing.T) {
	var tests []pathSuffixTest = []pathSuffixTest{
		{
			name:     "Go",
			path:     "main.go",
			expected: fayl.Go,
		},
		{
			name:     "Header",
			path:     "file.hpp",
			expected: fayl.Header,
		},
		{
			name:     "Yaml",
			path:     "config.yml",
			expected: fayl.Yaml,
		},
		{
			name:     "NoSuffix",
			path:     "README",
			expected: fayl.None,
		},
		{
			name:     "Unknown",
			path:     "file.unknown",
			expected: fayl.None,
			hasError: true,
		},
	}

	var i int
	for i = range tests {
		var test pathSuffixTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			var suffix fayl.Suffix
			var exception *gopolutils.Exception
			suffix, exception = fayl.PathFrom(test.path).Suffix()

			if test.hasError {
				if exception == nil {
					t.Fatal("expected exception")
				}
				return
			} else if exception != nil {
				t.Fatalf("unexpected exception: %v", exception)
			} else if suffix != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, suffix)
			}
		})
	}
}

func TestParent(t *testing.T) {
	var dir string = t.TempDir()

	var child string = filepath.Join(dir, "child")
	var parent *fayl.Path
	var exception *gopolutils.Exception
	parent, exception = fayl.PathFrom(child).Parent()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	} else if parent.String() != dir {
		t.Fatalf("expected %q, got %q", dir, parent.String())
	}
}

func TestPathString(t *testing.T) {
	const raw string = "abc/def"

	var path *fayl.Path = fayl.PathFrom(raw)

	if path.String() != raw {
		t.Fatalf("expected %q, got %q", raw, path.String())
	}
}

func TestNewPath(t *testing.T) {
	var path *fayl.Path = fayl.NewPath()
	var expected string
	var directoryError error
	expected, directoryError = os.Getwd()
	if directoryError != nil {
		t.Fatal(directoryError)
	} else if path.String() != expected {
		t.Fatalf("expected %q, got %q", expected, path.String())
	}
}

func TestConfiguration(t *testing.T) {
	var path *fayl.Path = fayl.Configuration()
	var expected string
	var directoryError error
	expected, directoryError = os.UserConfigDir()
	if directoryError != nil {
		t.Fatal(directoryError)
	} else if path.String() != expected {
		t.Fatalf("expected %q, got %q", expected, path.String())
	}
}

func TestHome(t *testing.T) {
	var path *fayl.Path = fayl.Home()

	expected, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	} else if path.String() != expected {
		t.Fatalf("expected %q, got %q", expected, path.String())
	}
}

func TestCache(t *testing.T) {
	var path *fayl.Path = fayl.Cache()
	var expected string
	var directoryError error
	expected, directoryError = os.UserCacheDir()
	if directoryError != nil {
		t.Fatal(directoryError)
	} else if path.String() != expected {
		t.Fatalf("expected %q, got %q", expected, path.String())
	}
}

func TestRoot(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom(".")
	var root *fayl.Path
	var exception *gopolutils.Exception
	root, exception = path.Root()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	} else if root.String() == "" {
		t.Fatal("expected non-empty root")
	}
}

func TestParentOfRoot(t *testing.T) {
	var root *fayl.Path
	var exception *gopolutils.Exception
	root, exception = fayl.NewPath().Root()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	}
	var parent *fayl.Path
	parent, exception = root.Parent()
	if exception == nil {
		t.Fatal("expected exception")
	} else if parent != nil {
		t.Fatal("expected nil parent")
	}
}

func TestSuffixHiddenFile(t *testing.T) {
	var exception *gopolutils.Exception
	_, exception = fayl.PathFrom(".gitignore").Suffix()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestSuffixTrailingDot(t *testing.T) {
	var suffix fayl.Suffix
	var exception *gopolutils.Exception
	suffix, exception = fayl.PathFrom("file.").Suffix()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	} else if suffix != fayl.None {
		t.Fatalf("expected %v, got %v", fayl.None, suffix)
	}
}

func TestSuffixMultipleDots(t *testing.T) {
	var suffix fayl.Suffix
	var exception *gopolutils.Exception
	suffix, exception = fayl.PathFrom("archive.tar.gz").Suffix()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	} else if suffix != fayl.Gz {
		t.Fatalf("expected %v, got %v", fayl.Gz, suffix)
	}
}

func TestSuffixUppercase(t *testing.T) {
	var exception *gopolutils.Exception
	_, exception = fayl.PathFrom("main.GO").Suffix()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestPathFromPartsNoneSuffix(t *testing.T) {
	var path *fayl.Path = fayl.PathFromParts("folder", "file", fayl.None)

	var expected string = filepath.Join("folder", "file.")
	if path.String() != expected {
		t.Fatalf("expected %q, got %q", expected, path.String())
	}
}

func TestJoinDoesNotModifyOriginal(t *testing.T) {
	var original *fayl.Path = fayl.PathFrom("a")

	var joined *fayl.Path = original.JoinAs("b")

	if original.String() != "a" {
		t.Fatal("original path was modified")
	}

	var expected string = filepath.Join("a", "b")
	if joined.String() != expected {
		t.Fatalf("expected %q, got %q", expected, joined.String())
	}
}

func TestAppendMultiple(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom("a")

	path.AppendAs("b")
	path.AppendAs("c")
	path.AppendAs("d")

	var expected string = filepath.Join("a", "b", "c", "d")
	if path.String() != expected {
		t.Fatalf("expected %q, got %q", expected, path.String())
	}
}

func TestJoinMultiple(t *testing.T) {
	var path *fayl.Path = fayl.PathFrom("root")

	var joined *fayl.Path = path.JoinAs("a").JoinAs("b").JoinAs("c")

	var expected string = filepath.Join("root", "a", "b", "c")
	if joined.String() != expected {
		t.Fatalf("expected %q, got %q", expected, joined.String())
	}
}

func TestAbsoluteAlreadyAbsolute(t *testing.T) {
	var dir string = t.TempDir()
	var absolute *fayl.Path
	var exception *gopolutils.Exception
	absolute, exception = fayl.PathFrom(dir).Absolute()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	} else if absolute.String() != dir {
		t.Fatalf("expected %q, got %q", dir, absolute.String())
	}
}

func TestExistsFile(t *testing.T) {
	var dir string = t.TempDir()
	var file string = filepath.Join(dir, "test.txt")

	var writeError error = os.WriteFile(file, []byte("hello"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	} else if !fayl.PathFrom(file).Exists() {
		t.Fatal("expected file to exist")
	}
}

func TestExistsEmptyPath(t *testing.T) {
	if !fayl.PathFrom("").Exists() {
		t.Fatal("expected empty path to exist because it resolves to current directory")
	}
}

func TestParentRelativePath(t *testing.T) {
	var parent *fayl.Path
	var exception *gopolutils.Exception
	parent, exception = fayl.PathFrom(filepath.Join("a", "b", "c")).Parent()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	}

	var expected string = filepath.Join("a", "b")
	if parent.String() != expected {
		t.Fatalf("expected %q, got %q", expected, parent.String())
	}
}

func TestRootIsStable(t *testing.T) {
	var root1 *fayl.Path
	var exception *gopolutils.Exception
	root1, exception = fayl.NewPath().Root()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	}
	var root2 *fayl.Path
	root2, exception = fayl.NewPath().Root()
	if exception != nil {
		t.Fatalf("unexpected exception: %v", exception)
	} else if root1.String() != root2.String() {
		t.Fatal("root should be deterministic")
	}
}

func TestPathFromEmptyString(t *testing.T) {
	var expected *fayl.Path = fayl.NewPath()

	var actual *fayl.Path = fayl.PathFrom("")

	if actual.String() != expected.String() {
		t.Fatalf("expected %q, got %q", expected.String(), actual.String())
	}
}
