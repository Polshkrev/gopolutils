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

// Obtain the byte size of the directory.
// Returns a [Size] of each of the entries in the directory.
func (directory Directory) ByteSize() Size {
	var size gopolutils.Size = 0
	var i int
	for i = range directory.Collect() {
		size += directory.Collect()[i].ByteSize().Size()
	}
	return *SizeFromBytes(size)
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
	var root string = directory.Root().String()
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
		var entry *Entry = NewEntry(PathFrom(path))
		entry.SetType(gopolutils.Must(assignType(entry.Path().String())))
		directory.Append(entry)
	}
	return nil
}

// Obtain the relative [Path] of the given target based on the given root.
// Returns the relative [Path] of the given target based on the given root.
func getRelativePath(root, target *Path) *Path {
	var path string
	var relativeError error
	path, relativeError = filepath.Rel(root.String(), target.String())
	if relativeError != nil {
		panic(gopolutils.NewNamedException(gopolutils.IOError, "%s", relativeError.Error()))
	}
	return PathFrom(path)
}

// Append a given relative [Entry] path to the given destination [Directory].
func assignRelative(destination *Directory, entries collections.View[*Entry]) {
	var i gopolutils.Size
	for i = range collections.Enumerate(entries) {
		var entry *Entry = entries.Collect()[i]
		var path *Path = destination.Root().Join(*entry.Path())
		var absolute *Path = gopolutils.Must(path.Absolute())
		var item *Entry = NewEntry(absolute)
		item.SetType(entry.Type())
		destination.Append(item)
	}
}

// Obtain the relative [Path] of each of the items within the given base [Directory].
// Returns a [collections.View] of [Entry]s based on the relative [Path] from the directory's root.
func relativeCopyDirectory(base Directory) collections.View[*Entry] {
	var result collections.Collection[*Entry] = collections.NewArray[*Entry]()
	var i gopolutils.Size
	for i = range collections.Enumerate(base.entries) {
		var entry *Entry = base.Collect()[i]
		var path *Path = getRelativePath(base.Root(), entry.Path())
		var item *Entry = NewEntry(path)
		item.SetType(entry.Type())
		result.Append(item)
	}
	return result
}

// Concurrently copy a given [Directory] to a destination [Directory].
func copyDirectoryConcurrent(directory, destination Directory, exceptChannel chan<- *gopolutils.Exception) {
	defer close(exceptChannel)
	var i int
	for i = range directory.Collect() {
		var item *Entry = directory.Collect()[i]
		var destinationEntry *Entry = destination.Collect()[i]
		var except *gopolutils.Exception = item.Copy(destinationEntry)
		if except != nil {
			exceptChannel <- except
			return
		}
	}
}

// Copy each of the entries in the directory to a given destination directory.
// If the given destination is determined to be empty, the given directory is extended with the entries from the source directory.
// If the destination entry does not initially exist and subsequently can not be created, an [gopolutils.IOError] is returned.
func (directory Directory) Copy(destination *Directory) *gopolutils.Exception {
	if directory.IsEmpty() {
		return nil
	} else if !destination.root.Exists() {
		var except *gopolutils.Exception = destination.Create()
		if except != nil {
			return except
		}
	}

	var empty bool = destination.IsEmpty()
	if empty {
		var entries collections.View[*Entry] = relativeCopyDirectory(directory)
		assignRelative(destination, entries)
	}

	var exceptChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception, 1)
	go copyDirectoryConcurrent(directory, *destination, exceptChannel)
	var except *gopolutils.Exception = <-exceptChannel
	return except
}

// Represent the directory as a string.
// Returns a representation of the directory as a string.
func (directory Directory) String() string {
	var buffer *strings.Builder = &strings.Builder{}
	var i int
	for i = range directory.Collect() {
		var item *Entry = directory.Collect()[i]
		fmt.Fprintf(buffer, "%s%c%s - %s\n", directory.Root(), filepath.Separator, item.Path(), item.Type())
	}
	return buffer.String()
}
