package fayl

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Polshkrev/gopolutils"
)

// Representation of a filesystem path.
type Path struct {
	raw string
}

// Construct a new filesystem path.
// Returns a pointer to a new path containing the current working directory.
// If the current working directory can not be obtained, an OSError is printed to standard error and the programme exits.
func NewPath() *Path {
	var path *Path = new(Path)
	var workingDirectory string
	var err error
	workingDirectory, err = os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, gopolutils.NewNamedException("OSError", err.Error()))
		os.Exit(1)
	}
	path.raw = workingDirectory
	return path
}

// Construct a new filesystem path from a given path string.
// Returns a new filesystem path containing the given path.
func PathFrom(path string) *Path {
	var result *Path = new(Path)
	result.raw = path
	return result
}

// Construct a new filesystem path from its given parts.
// Returns a new filesystem path containing the absolute path composed of the given parts.
// If the absolute path of the given parts can not be obtained, an OSError is printed to standard error and the programme exits.
func PathFromParts(folderName, fileName, fileType string) *Path {
	var buffer strings.Builder = strings.Builder{}
	buffer.WriteString(filepath.Join(folderName, fileName))
	buffer.WriteByte('.')
	buffer.WriteString(fileType)
	var absolute string
	var absoluteError error
	absolute, absoluteError = filepath.Abs(buffer.String())
	if absoluteError != nil {
		fmt.Fprintln(os.Stderr, gopolutils.NewNamedException("OSError", absoluteError.Error()))
		os.Exit(1)
	}
	return PathFrom(absolute)
}

// Determine if the filesystem path exists.
// Returns true if the filesystem path is evaluated to exist, else false.
func (path Path) Exists() bool {
	var err error
	_, err = os.Stat(path.raw)
	return !errors.Is(err, os.ErrNotExist)
}

// Obtain the absolute path.
// Returns a pointer to a new path containing the absolute filesystem path.
// If the absolute path can not be obtained, an OSError is returned with a nil data pointer.
func (path Path) Absolute() (*Path, *gopolutils.Exception) {
	var absolute string
	var absoluteError error
	absolute, absoluteError = filepath.Abs(path.raw)
	if absoluteError != nil {
		return nil, gopolutils.NewNamedException("OSError", absoluteError.Error())
	}
	return PathFrom(absolute), nil
}

// Append a filesystem path to another.
func (path *Path) Append(item Path) {
	path.raw = filepath.Join(path.raw, item.raw)
}

// Append a filesystem path as a string to a path object.
// If the absolute path can not be obtained, an OSError is printed to standard error and the programme exits.
func (path *Path) AppendAs(item string) {
	var absolute *Path
	var except *gopolutils.Exception
	absolute, except = PathFrom(item).Absolute()
	if except != nil {
		fmt.Fprintln(os.Stderr, except.Error())
		os.Exit(1)
	}
	path.Append(*absolute)
}

// Obtain the suffix of the filesystem path.
// If the suffix can not be obtained, an OSError is returned with an empty string.
func (path Path) Suffix() (string, *gopolutils.Exception) {
	var index int = strings.LastIndexByte(path.raw, '.')
	if index < 0 {
		return "", gopolutils.NewNamedException("OSError", "Path does not have an associated suffix.")
	}
	return path.raw[index+1:], nil
}

// Obtain the string representation of the root of the filesystem.
// Returns a string representing the root of the filesystem path.
// If the root can not be obtained, an OSError is returned with a nil data pointer.
func getRoot(filePath string) (string, *gopolutils.Exception) {
	var index int = strings.IndexRune(filePath, filepath.Separator)
	if index < 0 {
		return "", gopolutils.NewNamedException("OSError", "Path does not have an associated root.")
	}
	return filePath[:index], nil
}

// Obtain the root of the filesystem as a path.
// Returns a pointer to the path of the root of the filesystem.
// If the absolute path can not be obtained, an OSError is returned with a nil data pointer.
// If the root of the filesystem can not be obtained, an OSError is returned with a nil data pointer.
func (path Path) Root() (*Path, *gopolutils.Exception) {
	var absolute string
	var absoluteError error
	absolute, absoluteError = filepath.Abs(path.raw)
	if absoluteError != nil {
		return nil, gopolutils.NewNamedException("OSError", absoluteError.Error())
	}
	var root string
	var rootExcept *gopolutils.Exception
	root, rootExcept = getRoot(absolute)
	if rootExcept != nil {
		return nil, rootExcept
	}
	return PathFrom(root), nil
}

// Determine if the given path is the root of the filesystem.
// If the given path is evaluated to be the root, an OSError is returned.
func checkRoot(path Path) *gopolutils.Exception {
	var root *Path
	var rootExcept *gopolutils.Exception
	root, rootExcept = NewPath().Root()
	if rootExcept != nil {
		return rootExcept
	} else if strings.Compare(root.raw, path.raw) == 0 {
		return gopolutils.NewNamedException("OSError", "Can not get parent of filesystem root.")
	}
	return nil
}

// Obtain the parent directory of the filesystem path.
// Returns a pointer to a new path containing the parent directory of the path.
// If the parent can not be obtained, an OSError is returned with a nil data pointer.
// If the path is the root of the filesystem, an OSError is returned with a nil data pointer.
func (path Path) Parent() (*Path, *gopolutils.Exception) {
	var rootExcept *gopolutils.Exception = checkRoot(path)
	if rootExcept != nil {
		return nil, rootExcept
	}
	return PathFrom(filepath.Dir(path.raw)), nil
}

// Represent a filesystem path as a string.
// Returns a string representation of the filesystem path.
func (path Path) ToString() string {
	return path.raw
}
