package fayl

import (
	"fmt"
	"os"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections/safe"
)

// Representation of the different finite types of files.
type EntryType = gopolutils.StringEnum

const (
	DirectoryType EntryType = "Directory"
	FileType      EntryType = "File"
)

// Representation of a file on the filesystem.
type Entry struct {
	path    *Path
	kind    EntryType
	content safe.Collection[byte]
}

// Construct a new file entry based on a given [Path]. The path and therefore corresponding entry may be ephemeral.
// Returns a pointer to a new entry with a given path.
func NewEntry(path *Path) *Entry {
	var entry *Entry = new(Entry)
	entry.path = path
	entry.kind = FileType
	entry.content = safe.NewArray[byte]()
	return entry
}

// Obtain the [Path] location of the file entry.
// Returns a mutable pointer to the internal path location of the file entry.
func (entry Entry) Path() *Path {
	return entry.path
}

// Obtain the type of the file entry.
// Returns the [EntryType] of the file entry.
func (entry Entry) Type() EntryType {
	return entry.kind
}

// Obtain the contents of the file.
// This method does not open a file.
// Returns the byte content stored in the file entry.
func (entry Entry) Content() safe.Collection[byte] {
	return entry.content
}

// Set the path of the file entry.
func (entry *Entry) SetPath(path *Path) {
	entry.path = path
}

// Set the type of the file entry.
func (entry *Entry) SetType(kind EntryType) {
	entry.kind = kind
}

// Set the content of the file entry.
func (entry *Entry) SetContent(content safe.Collection[byte]) {
	entry.content = content
}

// Determine if the entry is of a given type.
// Returns true if the entry is of the given type, else false.
func (entry Entry) Is(kind EntryType) bool {
	return entry.kind == kind
}

// Generic dispatch creation method.
// If the entry already exists on the filesystem, a [gopolutils.FileExistsError] is returned.
// If the entry can not be created, an [gopolutils.IOError] is returned.
func (entry Entry) Create() *gopolutils.Exception {
	if entry.Path().Exists() {
		return gopolutils.NewNamedException(gopolutils.FileExistsError, "File '%s' already exists.", entry.Path().ToString())
	}
	var entryType Suffix
	var except *gopolutils.Exception
	entryType, except = SuffixFromString(entry.Path().ToString())
	if except != nil {
		return except
	}
	switch entryType {
	case None:
		return entry.MakeDirectory()
	default:
		return entry.Touch()
	}
}

// Concurrently create a file on the filesystem.
func concurrentTouch(path string, resultChannel chan<- *os.File, errorChannel chan<- error) {
	var result *os.File
	var createError error
	result, createError = os.Create(path)
	resultChannel <- result
	errorChannel <- createError
	defer close(resultChannel)
	defer close(errorChannel)
}

// Create a file on the filesystem.
// If the entry already exists, a [gopolutils.FileExistsError] is returned.
// If the entry is not a file, an [gopolutils.IsADirectoryError] is returned.
// If the file can not be created, an [gopolutils.IOError] is returned.
func (entry Entry) Touch() *gopolutils.Exception {
	if entry.Path().Exists() {
		return gopolutils.NewNamedException(gopolutils.FileExistsError, "File '%s' already exists.", entry.Path().ToString())
	} else if !entry.Is(FileType) {
		return gopolutils.NewNamedException(gopolutils.IsADirectoryError, "Entry '%s' is not a file.", entry.Path().ToString())
	}
	var resultChannel chan *os.File = make(chan *os.File, 1)
	var errorChannel chan error = make(chan error, 1)
	go concurrentTouch(entry.Path().ToString(), resultChannel, errorChannel)
	var result *os.File = <-resultChannel
	var touchError error = <-errorChannel
	defer result.Close()
	if touchError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, touchError.Error())
	}
	return nil
}

// Concurrently create a directory on the filesystem.
func concurrentMakeDirectory(path string, errorChannel chan<- error) {
	errorChannel <- os.MkdirAll(path, 0755)
	defer close(errorChannel)
}

// Create a directory on the filesystem.
// If the entry already exists on the filesystem, a [gopolutils.FileExistsError] is returned.
// If the directory can not be created, an [gopolutils.IOError] is returned.
func (entry Entry) MakeDirectory() *gopolutils.Exception {
	if entry.Path().Exists() {
		return gopolutils.NewNamedException(gopolutils.FileExistsError, "Directory '%s' already exists.", entry.Path().ToString())
	} else if !entry.Is(DirectoryType) {
		return gopolutils.NewNamedException(gopolutils.NotADirectoryError, "Entry '%s' is not a directory.", entry.Path().ToString())
	}
	var errorChannel chan error = make(chan error, 1)
	go concurrentMakeDirectory(entry.Path().ToString(), errorChannel)
	var makeDirectoryError error = <-errorChannel
	if makeDirectoryError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, makeDirectoryError.Error())
	}
	return nil
}

// Copy an entry into a given destination.
// After the copy has been completed on the filesystem, the given internal content of the destination entry is set to the internal content of the original entry.
// If the destination entry does not initially exist and subsequently can not be created, an [gopolutils.IOError] is returned.
func (entry Entry) Copy(destination *Entry) *gopolutils.Exception {
	if !entry.Path().Exists() {
		var except *gopolutils.Exception = entry.Create()
		if except != nil {
			return except
		}
	} else if !destination.Path().Exists() {
		var except *gopolutils.Exception = destination.Create()
		if except != nil {
			return except
		}
	} else if destination.Is(FileType) {
		var except *gopolutils.Exception = Write(destination.Path(), entry.Content().Collect())
		if except != nil {
			return except
		}
		destination.SetContent(entry.Content())
	}
	return nil
}

// Concurrently remove a file on the filesystem.
func concurrentRemoveFile(path string, errorChannel chan<- error) {
	errorChannel <- os.Remove(path)
	defer close(errorChannel)
}

// Remove a file on the filesystem.
// If the entry does not exist, a [gopolutils.FileNotFoundError] is returned.
// If the entry is not a file, an [gopolutils.IsADirectoryError] is returned.
// If the file can not be removed, an [gopolutils.IOError] is returned.
func (entry Entry) RemoveFile() *gopolutils.Exception {
	if !entry.Path().Exists() {
		return gopolutils.NewNamedException(gopolutils.FileNotFoundError, "File '%s' can not be found.", entry.Path().ToString())
	} else if !entry.Is(FileType) {
		return gopolutils.NewNamedException(gopolutils.IsADirectoryError, "Entry '%s' is not a file.", entry.Path().ToString())
	}
	var errorChannel chan error = make(chan error, 1)
	go concurrentRemoveFile(entry.Path().ToString(), errorChannel)
	var removeFileError error = <-errorChannel
	if removeFileError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, removeFileError.Error())
	}
	return nil
}

// Concurrently remove a directory.
func concurrentRemoveDirectory(path string, errorChannel chan<- error) {
	errorChannel <- os.RemoveAll(path)
	defer close(errorChannel)
}

// Remove a directory on the filesystem.
// If the entry does not exist on the filesystem, a [gopolutils.FileNotFoundError] is returned.
// If the directory is not a directory, a [gopolutils.NotADirectoryError] is returned.
// If the directory can not be removed, an [gopolutils.IOError] is returned.
func (entry Entry) RemoveDirectory() *gopolutils.Exception {
	if !entry.Path().Exists() {
		return gopolutils.NewNamedException(gopolutils.FileNotFoundError, "Directory '%s' does not exist.", entry.Path().ToString())
	} else if !entry.Is(DirectoryType) {
		return gopolutils.NewNamedException(gopolutils.NotADirectoryError, "Entry '%s' is not a directory.", entry.Path().ToString())
	}
	var errorChannel chan error = make(chan error, 1)
	go concurrentRemoveDirectory(entry.Path().ToString(), errorChannel)
	var removeDirectoryError error = <-errorChannel
	if removeDirectoryError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, removeDirectoryError.Error())
	}
	return nil
}

// Assign the type of a file on the filesystem.
// Returns an [EntryType] obtained through [os.Stat].
// If the [os.FileInfo] of the given path can not be obtained, an [gopolutils.IOError] is returned with a nil result.
func assignType(path string) (EntryType, *gopolutils.Exception) {
	var info os.FileInfo
	var infoError error
	info, infoError = os.Stat(path)
	if infoError != nil {
		return "", gopolutils.NewNamedException(gopolutils.IOError, infoError.Error())
	}
	if info.IsDir() {
		return DirectoryType, nil
	}
	return FileType, nil
}

// Generic dispatch removal method.
// If the entry does not exist on the filesystem, a [gopolutils.FileNotFoundError] is returned.
// If the [os.FileInfo] of the entry can not be obtained, an [gopolutils.IOError] is returned.
// If the entry is a file and the file can not be removed, an [gopolutils.IOError] is returned.
// If the entry is a directory and the directory can not be removed, an [gopolutils.IOError] is returned.
func (entry Entry) Remove() *gopolutils.Exception {
	if !entry.Path().Exists() {
		return gopolutils.NewNamedException(gopolutils.FileNotFoundError, "Directory '%s' can not be found.", entry.Path().ToString())
	}
	var entryType EntryType
	var except *gopolutils.Exception
	entryType, except = assignType(entry.Path().ToString())
	if except != nil {
		return except
	}
	switch entryType {
	case DirectoryType:
		return entry.RemoveDirectory()
	default:
		return entry.RemoveFile()
	}
}

func (entry Entry) ToString() string {
	return fmt.Sprintf("%s - %s", entry.Path().ToString(), entry.Type())
}
