package fayl

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/collections/safe"
)

// Representation of a directory containing file entries.
type Directory struct {
	root    *Path
	entries safe.Collection[*Entry]
}

// Construct a new directory from its given root [Path].
// Returns a new [Directory] from its given root [Path].
func NewDirectory(root *Path) *Directory {
	var directory *Directory = new(Directory)
	directory.root = root
	directory.entries = safe.NewArray[*Entry]()
	return directory
}

// Append a directory entry to the directory.
func (directory *Directory) Append(entry *Entry) {
	directory.entries.Append(entry)
}

// Append a [collections.View] of [Entry] to the directory.
func (directory *Directory) Extend(entries collections.View[*Entry]) {
	directory.entries.Extend(entries)
}

// Obtain a pointer to the [Entry] at a given [gopolutils.Size] index.
// Returns a pointer to the [Entry] at the given index.
// If the given index is greater than the size of the collection, an [gopolutils.OutOfRangeError] is returned with a nil data pointer.
func (directory Directory) At(index gopolutils.Size) (**Entry, *gopolutils.Exception) {
	return directory.entries.At(index)
}

// Remove an [Entry] at the given [gopolutils.Size] index.
// If the given index is greater than the size of the collection, an [gopolutils.OutOfRangeError] is returned.
func (directory *Directory) Remove(index gopolutils.Size) *gopolutils.Exception {
	return directory.entries.Remove(index)
}

// Access a pointer to a slice of the entries within the directory.
// Returns a mutable pointer to the underlying entries within the directory.
func (directory Directory) Items() *[]*Entry {
	return directory.entries.Items()
}

// Collect the directory's entries into a slice.
// Returns a slice of [Entry].
func (directory Directory) Collect() []*Entry {
	return directory.entries.Collect()
}

// Obtain the size of the directory.
// Returns a [gopolutils.Size] of the directory.
func (directory Directory) Size() gopolutils.Size {
	return directory.entries.Size()
}

// Determine if the directory is empty.
// Returns true if the directory's size is equal to zero or if the underlying data is nil, else false.
func (directory Directory) IsEmpty() bool {
	return directory.entries.IsEmpty()
}

// Obtain the root of the directory.
// Returns the root of the directory.
func (directory Directory) Root() *Path {
	return directory.root
}

// Persist each of the entries on the filesystem.
// If the entry already exists on the filesystem, a [gopolutils.FileExistsError] is returned.
// If the entry can not be created, an [gopolutils.IOError] is returned.
func (directory Directory) Create() *gopolutils.Exception {
	var i int
	for i = range directory.Collect() {
		var item *Entry = directory.Collect()[i]
		var except *gopolutils.Exception = item.Create()
		if except != nil {
			return except
		}
	}
	return nil
}

// Delete each of the entries within the directory.
// If the entry does not exist on the filesystem, a [gopolutils.FileNotFoundError] is returned.
// If the [os.FileInfo] of the entry can not be obtained, an [gopolutils.IOError] is returned.
// If the entry is a file and the file can not be removed, an [gopolutils.IOError] is returned.
// If the entry is a directory and the directory can not be removed, an [gopolutils.IOError] is returned.
func (directory Directory) Delete() *gopolutils.Exception {
	var i int
	for i = range directory.Collect() {
		var item *Entry = directory.Collect()[i]
		var except *gopolutils.Exception = item.Remove()
		if except != nil {
			return except
		}
	}
	return nil
}

// Concurrently walk a directory with a given root.
func walkConcurrent(root string, paths chan<- []string, errorChannel chan<- error) {
	defer close(paths)
	defer close(errorChannel)
	var result []string = make([]string, 0)
	var walkError error = filepath.WalkDir(root, func(path string, _ fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if path == root {
			return nil
		}
		result = append(result, path)
		return nil
	})
	paths <- result
	errorChannel <- walkError
}

// Recursively append each of the child entry paths to the directory.
// If the entries can not be obtained, an [gopolutils.OSError] is returned.
func (directory *Directory) Read() *gopolutils.Exception {
	var root string = directory.Root().ToString()
	var pathsChannel chan []string = make(chan []string, 1)
	var errorChannel chan error = make(chan error, 1)
	go walkConcurrent(root, pathsChannel, errorChannel)
	var paths []string = <-pathsChannel
	var walkError error = <-errorChannel
	if walkError != nil {
		return gopolutils.NewNamedException(gopolutils.OSError, walkError.Error())
	}
	var i int
	for i = range paths {
		var path string = paths[i]
		directory.Append(NewEntry(PathFrom(path)))
	}
	return nil
}

// Copy each of the entries in the directory to a given destination directory.
// If the given destination is determined to be empty, the given directory is extended with the entries from the source directory.
// If the destination entry does not initially exist and subsequently can not be created, an [gopolutils.IOError] is returned.
func (directory Directory) Copy(destination *Directory) *gopolutils.Exception {
	if destination.IsEmpty() {
		destination.Extend(directory)
	}
	var i int
	for i = range directory.Collect() {
		var item *Entry = directory.Collect()[i]
		var destinationEntry *Entry = destination.Collect()[i]
		var except *gopolutils.Exception = item.Copy(destinationEntry)
		if except != nil {
			return except
		}
	}
	return nil
}

// Represent the directory as a string.
// Returns a representation of the directory as a string.
func (directory Directory) ToString() string {
	var buffer *strings.Builder = &strings.Builder{}
	var i int
	for i = range directory.Collect() {
		var item *Entry = directory.Collect()[i]
		buffer.WriteString(fmt.Sprintf("%s%c%s - %s", directory.Root().ToString(), filepath.Separator, item.Path().ToString(), item.Type()))
	}
	return buffer.String()
}
