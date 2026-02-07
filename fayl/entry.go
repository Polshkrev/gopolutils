package fayl

import (
	"fmt"
	"os"

	"github.com/Polshkrev/gopolutils"
)

// Representation of the different finite types of files.
type EntryType = gopolutils.StringEnum

const (
	Directory EntryType = "Directory"
	File      EntryType = "File"
)

// Representation of a file on the filesystem.
type Entry struct {
	path    *Path
	kind    EntryType
	content []byte
}

// Construct a new file entry based on a given [Path]. The path and therefore corresponding entry may be ephemeral.
// Returns a pointer to a new entry with a given path.
func NewEntry(path *Path) *Entry {
	var entry *Entry = new(Entry)
	entry.path = path
	entry.kind = File
	entry.content = make([]byte, 0)
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
func (entry Entry) Read() []byte {
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
func (entry *Entry) SetContent(content []byte) {
	entry.content = content
}

// Determine if the entry is of a given type.
// Returns true if the entry is of the given type, else false.
func (entry Entry) Is(kind EntryType) bool {
	return entry.kind == kind
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
		return Directory, nil
	}
	return File, nil
}

// Generic dispatch creation method.
// If the entry does not exist on the file system, a [gopolutils.FileNotFoundError] is returned.
// If the [os.FileInfo] of the entry can not be obtained, an [gopolutils.IOError] is returned.
// If the entry already exists on the filesystem, a [gopolutils.FileExistsError] is returned.
// If the entry is a file and the file can not be created, an [gopolutils.IOError] is returned.
// If the entry is a directory and the directory can not be created, an [gopolutils.IOError] is returned.
func (entry Entry) Create() *gopolutils.Exception {
	if !entry.Path().Exists() {
		return gopolutils.NewNamedException(gopolutils.FileNotFoundError, fmt.Sprintf("File '%s' can not be found.", entry.Path().ToString()))
	}
	var entryType EntryType
	var except *gopolutils.Exception
	entryType, except = assignType(entry.Path().ToString())
	if except != nil {
		return except
	}
	switch entryType {
	case Directory:
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
		return gopolutils.NewNamedException(gopolutils.FileExistsError, fmt.Sprintf("File '%s' already exists.", entry.Path().ToString()))
	} else if !entry.Is(File) {
		return gopolutils.NewNamedException(gopolutils.IsADirectoryError, fmt.Sprintf("Entry '%s' is not a file.", entry.Path().ToString()))
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
		return gopolutils.NewNamedException(gopolutils.FileExistsError, fmt.Sprintf("Directory '%s' already exists.", entry.Path().ToString()))
	} else if !entry.Is(Directory) {
		return gopolutils.NewNamedException(gopolutils.NotADirectoryError, fmt.Sprintf("Entry '%s' is not a directory.", entry.Path().ToString()))
	}
	var errorChannel chan error = make(chan error, 1)
	go concurrentMakeDirectory(entry.Path().ToString(), errorChannel)
	var makeDirectoryError error = <-errorChannel
	if makeDirectoryError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, makeDirectoryError.Error())
	}
	return nil
}

// Concurrently write content to a path.
func concurrentWrite(path string, content []byte, errorChannel chan<- error) {
	errorChannel <- os.WriteFile(path, content, 0644)
	defer close(errorChannel)
}

// Write the content of the file entry into a persistent file.
// If the file already exists, the content of the file is overridden.
// If the file can not be written to, an [gopolutils.IOError] is returned.
func (entry *Entry) Write(content []byte) *gopolutils.Exception {
	var errorChannel chan error = make(chan error, 1)
	go concurrentWrite(entry.Path().ToString(), content, errorChannel)
	var writeError error = <-errorChannel
	if writeError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, writeError.Error())
	}
	entry.SetContent(content)
	return nil
}

// Copy an entry into a given destination.
// After the copy has been completed on the filesystem, the given internal content of the destination entry is set to the internal content of the original entry.
// If the destination entry does not initially exist and subsequently can not be created, an [gopolutils.IOError] is returned.
func (entry Entry) Copy(destination *Entry) *gopolutils.Exception {
	if !destination.Path().Exists() {
		var except *gopolutils.Exception = destination.Touch()
		if except != nil {
			return except
		}
	} else if !entry.Path().Exists() {
		var except *gopolutils.Exception = entry.Touch()
		if except != nil {
			return except
		}
	}
	return destination.Write(entry.Read())
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
		return gopolutils.NewNamedException(gopolutils.FileNotFoundError, fmt.Sprintf("File '%s' can not be found.", entry.Path().ToString()))
	} else if !entry.Is(File) {
		return gopolutils.NewNamedException(gopolutils.IsADirectoryError, fmt.Sprintf("Entry '%s' is not a file.", entry.Path().ToString()))
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
		return gopolutils.NewNamedException(gopolutils.FileNotFoundError, fmt.Sprintf("Directory '%s' does not exist.", entry.Path().ToString()))
	} else if !entry.Is(Directory) {
		return gopolutils.NewNamedException(gopolutils.NotADirectoryError, fmt.Sprintf("Entry '%s' is not a directory.", entry.Path().ToString()))
	}
	var errorChannel chan error = make(chan error, 1)
	go concurrentRemoveDirectory(entry.Path().ToString(), errorChannel)
	var removeDirectoryError error = <-errorChannel
	if removeDirectoryError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, removeDirectoryError.Error())
	}
	return nil
}

// Generic dispatch removal method.
// If the entry does not exist on the filesystem, a [gopolutils.FileNotFoundError] is returned.
// If the [os.FileInfo] of the entry can not be obtained, an [gopolutils.IOError] is returned.
// If the entry is a file and the file can not be removed, an [gopolutils.IOError] is returned.
// If the entry is a directory and the directory can not be removed, an [gopolutils.IOError] is returned.
func (entry Entry) Remove() *gopolutils.Exception {
	if !entry.Path().Exists() {
		return gopolutils.NewNamedException(gopolutils.FileNotFoundError, fmt.Sprintf("Directory '%s' can not be found.", entry.Path().ToString()))
	}
	var entryType EntryType
	var except *gopolutils.Exception
	entryType, except = assignType(entry.Path().ToString())
	if except != nil {
		return except
	}
	switch entryType {
	case Directory:
		return entry.RemoveDirectory()
	default:
		return entry.RemoveFile()
	}
}
