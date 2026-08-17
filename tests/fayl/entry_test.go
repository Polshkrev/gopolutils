package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections/safe"
	"github.com/Polshkrev/gopolutils/fayl"
)

func TestNewEntry(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("test.txt"))

	if entry.Path().String() != "test.txt" {
		t.Fatal("unexpected path")
	} else if entry.Type() != fayl.FileType {
		t.Fatal("expected FileType")
	} else if !entry.Content().IsEmpty() {
		t.Fatal("expected empty content")
	}
}

func TestSetters(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("a"))

	var newPath *fayl.Path = fayl.PathFrom("b")
	entry.SetPath(newPath)
	entry.SetType(fayl.DirectoryType)

	var content *safe.Array[byte] = safe.NewArray[byte]()
	content.Append('a')
	content.Append('b')

	entry.SetContent(content)

	if entry.Path() != newPath {
		t.Fatal("SetPath failed")
	} else if entry.Type() != fayl.DirectoryType {
		t.Fatal("SetType failed")
	} else if entry.Content().Size() != 2 {
		t.Fatal("SetContent failed")
	}
}

func TestIs(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("file"))

	if !entry.Is(fayl.FileType) {
		t.Fatal("expected FileType")
	} else if entry.Is(fayl.DirectoryType) {
		t.Fatal("unexpected DirectoryType")
	}
}

func TestTouch(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))

	var exception *gopolutils.Exception = entry.Touch()
	var statError error
	_, statError = os.Stat(path)

	if exception != nil {
		t.Fatal(exception)
	} else if statError != nil {
		t.Fatal(statError)
	}
}

func TestTouchAlreadyExists(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")
	var writeError error = os.WriteFile(path, nil, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	var exception *gopolutils.Exception = entry.Touch()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestMakeDirectory(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "dir")

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	entry.SetType(fayl.DirectoryType)
	var exception *gopolutils.Exception = entry.MakeDirectory()
	if exception != nil {
		t.Fatal(exception)
	}
	var info os.FileInfo
	var statError error
	info, statError = os.Stat(path)
	if statError != nil {
		t.Fatal(statError)
	} else if !info.IsDir() {
		t.Fatal("expected directory")
	}
}

func TestMakeDirectoryAlreadyExists(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "dir")
	var makeError error = os.Mkdir(path, 0755)
	if makeError != nil {
		t.Fatal(makeError)
	}
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	entry.SetType(fayl.DirectoryType)
	var exception *gopolutils.Exception = entry.MakeDirectory()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestRemoveFile(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")
	var writeError error = os.WriteFile(path, []byte("abc"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	var exception *gopolutils.Exception = entry.RemoveFile()
	var statError error
	_, statError = os.Stat(path)
	if exception != nil {
		t.Fatal(exception)
	} else if !os.IsNotExist(statError) {
		t.Fatal("file still exists")
	}
}

func TestRemoveDirectory(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "dir")
	var makeError error = os.Mkdir(path, 0755)
	if makeError != nil {
		t.Fatal(makeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	entry.SetType(fayl.DirectoryType)
	var exception *gopolutils.Exception = entry.RemoveDirectory()
	var statError error
	_, statError = os.Stat(path)
	if exception != nil {
		t.Fatal(exception)
	} else if !os.IsNotExist(statError) {
		t.Fatal("directory still exists")
	}
}

func TestHandle(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")
	var writeError error = os.WriteFile(path, []byte("abc"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	var handle *os.File
	var exception *gopolutils.Exception
	handle, exception = entry.Handle()
	if exception != nil {
		t.Fatal(exception)
	}
	defer handle.Close()

	if handle == nil {
		t.Fatal("expected handle")
	}
}

func TestHandleMissingFile(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(filepath.Join(t.TempDir(), "missing.txt")))
	var handle *os.File
	var exception *gopolutils.Exception
	handle, exception = entry.Handle()

	if exception == nil {
		t.Fatal("expected exception")
	} else if handle != nil {
		t.Fatal("expected nil handle")
	}
}

func TestByteSizeMemory(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("missing"))

	entry.Content().Append('a')
	entry.Content().Append('b')
	entry.Content().Append('c')

	var size fayl.Size = entry.ByteSize()

	if size.Size() != 3 {
		t.Fatalf("expected 3, got %d", size.Size())
	}
}

func TestByteSizeFilesystem(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")
	var writeError error = os.WriteFile(path, []byte("hello"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))

	var size fayl.Size = entry.ByteSize()

	if size.Size() != 5 {
		t.Fatalf("expected 5, got %d", size.Size())
	}
}

func TestEntryIsEmpty(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("a"))

	if !entry.IsEmpty() {
		t.Fatal("expected empty")
	}

	entry.Content().Append(1)

	if entry.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

func TestEntrySize(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("a"))

	entry.Content().Append(1)
	entry.Content().Append(2)

	if entry.Size() != 2 {
		t.Fatal("unexpected size")
	}
}

func TestEntryString(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom("test.txt"))

	var expected string = "test.txt - File"

	if entry.String() != expected {
		t.Fatalf("expected %q, got %q", expected, entry.String())
	}
}

func TestRead(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")

	var expected []byte = []byte("hello world")
	var writeError error = os.WriteFile(path, expected, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	var exception *gopolutils.Exception = entry.Read()
	if exception != nil {
		t.Fatal(exception)
	} else if entry.Size() != gopolutils.Size(len(expected)) {
		t.Fatalf("expected %d bytes, got %d", len(expected), entry.Size())
	}

	var actual []byte = entry.Content().Collect()

	if !bytes.Equal(actual, expected) {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestReadMissingFile(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(filepath.Join(t.TempDir(), "missing.txt")))
	var exception *gopolutils.Exception = entry.Read()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestCreateFile(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	var exception *gopolutils.Exception = entry.Create()
	var statError error
	_, statError = os.Stat(path)
	if exception != nil {
		t.Fatal(exception)
	} else if statError != nil {
		t.Fatal(statError)
	}
}

func TestCreateDirectory(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "directory")

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	entry.SetType(fayl.DirectoryType)
	var exception *gopolutils.Exception = entry.Create()
	if exception != nil {
		t.Fatal(exception)
	}
	var info os.FileInfo
	var statError error
	info, statError = os.Stat(path)
	if statError != nil {
		t.Fatal(statError)
	} else if !info.IsDir() {
		t.Fatal("expected directory")
	}
}

func TestCreateAlreadyExists(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")
	var writeError error = os.WriteFile(path, nil, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	var exception *gopolutils.Exception = entry.Create()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestEntryCopyFile(t *testing.T) {
	var directory string = t.TempDir()

	sourcePath := filepath.Join(directory, "source.txt")
	destinationPath := filepath.Join(directory, "destination.txt")

	var content []byte = []byte("copy me")
	var writeError error = os.WriteFile(sourcePath, content, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var source *fayl.Entry = fayl.NewEntry(fayl.PathFrom(sourcePath))
	var exception *gopolutils.Exception = source.Read()
	if exception != nil {
		t.Fatal(exception)
	}

	var destination *fayl.Entry = fayl.NewEntry(fayl.PathFrom(destinationPath))
	exception = source.Copy(destination)
	if exception != nil {
		t.Fatal(exception)
	}
	var actual []byte
	var readError error
	actual, readError = os.ReadFile(destinationPath)
	if readError != nil {
		t.Fatal(readError)
	} else if !bytes.Equal(actual, content) {
		t.Fatalf("expected %q, got %q", content, actual)
	} else if destination.Size() != source.Size() {
		t.Fatal("destination content not updated")
	}
}

func TestEntryCopyCreatesDestination(t *testing.T) {
	var directory string = t.TempDir()

	var sourcePath string = filepath.Join(directory, "source.txt")
	var destinationPath string = filepath.Join(directory, "destination.txt")
	var writeError error = os.WriteFile(sourcePath, []byte("abc"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var source *fayl.Entry = fayl.NewEntry(fayl.PathFrom(sourcePath))
	var exception *gopolutils.Exception = source.Read()
	if exception != nil {
		t.Fatal(exception)
	}

	var destination *fayl.Entry = fayl.NewEntry(fayl.PathFrom(destinationPath))
	exception = source.Copy(destination)
	if exception != nil {
		t.Fatal(exception)
	} else if !destination.Path().Exists() {
		t.Fatal("destination file not created")
	}
}

func TestCopyMissingSource(t *testing.T) {
	var directory string = t.TempDir()

	var source *fayl.Entry = fayl.NewEntry(fayl.PathFrom(filepath.Join(directory, "missing.txt")))
	var destination *fayl.Entry = fayl.NewEntry(fayl.PathFrom(filepath.Join(directory, "destination.txt")))
	var exception *gopolutils.Exception = source.Copy(destination)
	if exception != nil {
		t.Fatal(exception)
	} else if !source.Path().Exists() {
		t.Fatal("expected source to be created")
	}
}

func TestRemoveDispatchFile(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")
	var writeError error = os.WriteFile(path, nil, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	var exception *gopolutils.Exception = entry.Remove()
	if exception != nil {
		t.Fatal(exception)
	} else if entry.Path().Exists() {
		t.Fatal("expected file to be removed")
	}
}

func TestRemoveDispatchDirectory(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "directory")
	var makeError error = os.Mkdir(path, 0755)
	if makeError != nil {
		t.Fatal(makeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	entry.SetType(fayl.DirectoryType)
	var exception *gopolutils.Exception = entry.Remove()
	if exception != nil {
		t.Fatal(exception)
	} else if entry.Path().Exists() {
		t.Fatal("expected directory to be removed")
	}
}

func TestRemoveMissing(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(filepath.Join(t.TempDir(), "missing")))
	var exception *gopolutils.Exception = entry.Remove()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestRemoveFileOnDirectory(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "directory")
	var makeError error = os.Mkdir(path, 0755)
	if makeError != nil {
		t.Fatal(makeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	entry.SetType(fayl.DirectoryType)
	var exception *gopolutils.Exception = entry.RemoveFile()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestRemoveDirectoryOnFile(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")
	var writeError error = os.WriteFile(path, nil, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}

	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(path))
	var exception *gopolutils.Exception = entry.RemoveDirectory()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestTouchDirectoryEntry(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(filepath.Join(t.TempDir(), "directory")))
	entry.SetType(fayl.DirectoryType)
	var exception *gopolutils.Exception = entry.Touch()
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestMakeDirectoryFileEntry(t *testing.T) {
	var entry *fayl.Entry = fayl.NewEntry(fayl.PathFrom(filepath.Join(t.TempDir(), "file")))
	var exception *gopolutils.Exception = entry.MakeDirectory()
	if exception == nil {
		t.Fatal("expected exception")
	}
}
