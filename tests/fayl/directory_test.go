package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/fayl"
)

func TestDirectoryAppend(t *testing.T) {
	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(t.TempDir()))

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("a.txt"))

	directory.Append(entry)

	if directory.Size() != 1 {
		t.Fatalf("expected 1 entry, got %d", directory.Size())
	}
}

func TestDirectoryExtend(t *testing.T) {
	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(t.TempDir()))

	var collection collections.Collection[*fayl.Entry] = collections.NewArray[*fayl.Entry]()

	var first *fayl.Entry = fayl.NewEntry(fayl.PathFrom("a.txt"))
	var second *fayl.Entry = fayl.NewEntry(fayl.PathFrom("b.txt"))

	collection.Append(first)
	collection.Append(second)

	directory.Extend(collection)

	if directory.Size() != 2 {
		t.Fatalf("expected 2 entries, got %d", directory.Size())
	}
}

func TestDirectoryAt(t *testing.T) {
	var dir *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(t.TempDir()))

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("a.txt"))

	dir.Append(entry)
	var actual **fayl.Entry
	var exception *gopolutils.Exception
	actual, exception = dir.At(0)
	if exception != nil {
		t.Fatal(exception)
	} else if *actual != entry {
		t.Fatal("incorrect entry")
	}
}

func TestDirectoryAtOutOfRange(t *testing.T) {
	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(t.TempDir()))
	var exception *gopolutils.Exception
	_, exception = directory.At(0)

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestDirectoryRemove(t *testing.T) {
	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(t.TempDir()))

	directory.Append(fayl.NewEntry(fayl.PathFrom("a.txt")))
	var exception *gopolutils.Exception = directory.Remove(0)
	if exception != nil {
		t.Fatal(exception)
	} else if !directory.IsEmpty() {
		t.Fatal("directory should be empty")
	}
}

func TestDirectoryRemoveOutOfRange(t *testing.T) {
	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(t.TempDir()))
	var exception *gopolutils.Exception = directory.Remove(0)
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestDirectoryCollect(t *testing.T) {
	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(t.TempDir()))

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("a.txt"))

	directory.Append(entry)

	var items []*fayl.Entry = directory.Collect()

	if len(items) != 1 {
		t.Fatal("incorrect length")
	} else if items[0] != entry {
		t.Fatal("incorrect item")
	}
}

func TestDirectoryByteSize(t *testing.T) {
	var root string = t.TempDir()

	var file string = filepath.Join(root, "file.txt")
	var writeError error = os.WriteFile(file, []byte("hello"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(root))
	directory.Append(fayl.NewEntry(fayl.PathFrom(file)))

	if directory.ByteSize().IsEmpty() {
		t.Fatal("expected non-zero size")
	}
}

func TestDirectoryCreate(t *testing.T) {
	var root string = t.TempDir()

	var path string = filepath.Join(root, "file.txt")

	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(root))
	directory.Append(fayl.NewEntry(fayl.PathFrom(path)))
	var exception *gopolutils.Exception = directory.Create()
	var statError error
	if exception != nil {
		t.Fatal(exception)
	} else if _, statError = os.Stat(path); statError != nil {
		t.Fatal(statError)
	}
}

func TestDirectoryDelete(t *testing.T) {
	var root string = t.TempDir()

	var path string = filepath.Join(root, "file.txt")
	var writeError error = os.WriteFile(path, nil, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(root))
	directory.Append(fayl.NewEntry(fayl.PathFrom(path)))
	var exception *gopolutils.Exception = directory.Delete()
	var statError error
	if exception != nil {
		t.Fatal(exception)
	} else if _, statError = os.Stat(path); !os.IsNotExist(statError) {
		t.Fatal("file still exists")
	}
}

func TestDirectoryRead(t *testing.T) {
	var root string = t.TempDir()
	var writeError error = os.WriteFile(filepath.Join(root, "a.txt"), nil, 0644)
	var makeError error
	if writeError != nil {
		t.Fatal(writeError)
	} else if makeError = os.Mkdir(filepath.Join(root, "child"), 0755); makeError != nil {
		t.Fatal(makeError)
	}

	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(root))
	var exception *gopolutils.Exception = directory.Read()
	if exception != nil {
		t.Fatal(exception)
	} else if directory.Size() != 2 {
		t.Fatalf("expected 2 entries, got %d", directory.Size())
	}
}

func TestDirectoryReadMissing(t *testing.T) {
	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(filepath.Join(t.TempDir(), "missing")))
	var exception *gopolutils.Exception = directory.Read()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestDirectoryString(t *testing.T) {
	var root string = t.TempDir()

	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(root))

	directory.Append(fayl.NewEntry(fayl.PathFrom("a.txt")))

	if directory.String() == "" {
		t.Fatal("expected non-empty string")
	}
}

func TestDirectoryMultipleRead(t *testing.T) {
	var root string = t.TempDir()
	var writeError error = os.WriteFile(filepath.Join(root, "a.txt"), []byte("a"), 0644)
	var makeError error
	if writeError != nil {
		t.Fatal(writeError)
	} else if makeError = os.Mkdir(filepath.Join(root, "child"), 0755); makeError != nil {
		t.Fatal(makeError)
	} else if writeError = os.WriteFile(filepath.Join(root, "child", "b.txt"), []byte("b"), 0644); writeError != nil {
		t.Fatal(writeError)
	}

	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(root))
	var exception *gopolutils.Exception = directory.Read()
	if exception != nil {
		t.Fatal(exception)
	} else if directory.Size() != 3 {
		t.Fatalf("expected 3 entries, got %d", directory.Size())
	}
}

func TestDirectoryReadMissingRoot(t *testing.T) {
	var root string = filepath.Join(t.TempDir(), "missing")

	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(root))
	var exception *gopolutils.Exception = directory.Read()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestDirectoryCopy(t *testing.T) {
	var sourceRoot string = t.TempDir()
	var destRoot string = t.TempDir()

	var sourceFile string = filepath.Join(sourceRoot, "test.txt")
	var writeError error = os.WriteFile(sourceFile, []byte("hello"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var source *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(sourceRoot))
	var exception *gopolutils.Exception = source.Read()
	if exception != nil {
		t.Fatal(exception)
	}

	var destination *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(destRoot))
	exception = destination.Read()
	if exception != nil {
		t.Fatal(exception)
	} else if exception := source.Copy(destination); exception != nil {
		t.Fatal(exception)
	}

	var copied string = filepath.Join(destRoot, "test.txt")
	var data []byte
	var readError error
	data, readError = os.ReadFile(copied)
	if readError != nil {
		t.Fatal(readError)
	} else if string(data) != "hello" {
		t.Fatal("file contents not copied")
	}
}

func TestDirectoryCopyIntoExistingDestination(t *testing.T) {
	var sourceRoot string = t.TempDir()
	var destRoot string = t.TempDir()
	var writeError error = os.WriteFile(filepath.Join(sourceRoot, "a.txt"), []byte("abc"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	} else if writeError = os.WriteFile(filepath.Join(destRoot, "a.txt"), nil, 0644); writeError != nil {
		t.Fatal(writeError)
	}

	var source *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(sourceRoot))
	source.Read()

	var destination *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(destRoot))
	destination.Read()
	var exception *gopolutils.Exception = source.Copy(destination)
	if exception != nil {
		t.Fatal(exception)
	}
	var data []byte
	var readError error
	data, readError = os.ReadFile(filepath.Join(destRoot, "a.txt"))
	if readError != nil {
		t.Fatal(readError)
	} else if string(data) != "abc" {
		t.Fatal("destination not overwritten")
	}
}

func TestDirectoryCopyEmptyDirectory(t *testing.T) {
	var source *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(t.TempDir()))
	var destination *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(t.TempDir()))
	var exception *gopolutils.Exception = source.Copy(destination)
	if exception != nil {
		t.Fatal(exception)
	} else if !destination.IsEmpty() {
		t.Fatal("destination should be empty")
	}
}

func TestDirectoryByteSizeSuccess(t *testing.T) {
	var root string = t.TempDir()

	var a string = filepath.Join(root, "a")
	var b string = filepath.Join(root, "b")

	os.WriteFile(a, []byte("12345"), 0644)
	os.WriteFile(b, []byte("1234567890"), 0644)

	var directory *fayl.Directory = fayl.NewDirectory(fayl.PathFrom(root))
	directory.Read()

	if directory.ByteSize().Size() != 15 {
		t.Fatalf("expected 15 bytes, got %d", directory.ByteSize().Size())
	}
}
