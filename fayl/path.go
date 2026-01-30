package fayl

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
		fmt.Fprintln(os.Stderr, gopolutils.NewNamedException(gopolutils.OSError, err.Error()))
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
// The fileType parametre is the file extension without the preceding dot.
// Returns a new filesystem path containing the absolute path composed of the given parts.
// If the absolute path of the given parts can not be obtained, an OSError is printed to standard error and the programme exits.
// If the path suffix is not defined in suffixToString, a `KeyError` is printed to standard error and the programme exits.
func PathFromParts(folderName, fileName string, fileType Suffix) *Path {
	var buffer strings.Builder = strings.Builder{}
	buffer.WriteString(filepath.Join(folderName, fileName))
	buffer.WriteByte('.')
	var suffixString string
	var suffixExcept *gopolutils.Exception
	suffixString, suffixExcept = StringFromSuffix(fileType)
	if suffixExcept != nil {
		fmt.Fprintln(os.Stderr, suffixExcept.Error())
		os.Exit(1)
	}
	buffer.WriteString(suffixString)
	var absolute string
	var absoluteError error
	absolute, absoluteError = filepath.Abs(buffer.String())
	if absoluteError != nil {
		fmt.Fprintln(os.Stderr, gopolutils.NewNamedException(gopolutils.OSError, absoluteError.Error()))
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
		return nil, gopolutils.NewNamedException(gopolutils.OSError, absoluteError.Error())
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
	var itemPath *Path = PathFrom(item)
	path.Append(*itemPath)
}

// Obtain the suffix of the filesystem path.
// If the suffix can not be obtained, an OSError is returned with a `None` suffix value.
// If the string representation of the suffix is not found within the global map, a `KeyError` is returned with a `None` suffix value.
func (path Path) Suffix() (Suffix, *gopolutils.Exception) {
	var index int = strings.LastIndexByte(path.raw, '.')
	if index < 0 {
		return None, gopolutils.NewNamedException(gopolutils.OSError, "Path does not have an associated suffix.")
	}
	var raw string = path.raw[index+1:]
	return SuffixFromString(raw)
}

// Obtain the string representation of the root of the filesystem.
// Returns a string representing the root of the filesystem path.
// If the root can not be obtained, an OSError is returned with a nil data pointer.
func getRoot(filePath string) (string, *gopolutils.Exception) {
	var index int = strings.IndexRune(filePath, filepath.Separator)
	if index < 0 {
		return "", gopolutils.NewNamedException(gopolutils.OSError, "Path does not have an associated root.")
	}
	return filePath[:index], nil
}

// Obtain the root of the filesystem as a path.
// Returns a pointer to the path of the root of the filesystem.
// If the absolute path can not be obtained, an OSError is returned with a nil data pointer.
// If the root of the filesystem can not be obtained, an OSError is returned with a nil data pointer.
func (path Path) Root() (*Path, *gopolutils.Exception) {
	if OS(runtime.GOOS) != WINDOWS { // ! This will error if value is not in enum list.
		return PathFrom("/"), nil
	}
	var absolute string
	var absoluteError error
	absolute, absoluteError = filepath.Abs(path.raw)
	if absoluteError != nil {
		return nil, gopolutils.NewNamedException(gopolutils.OSError, absoluteError.Error())
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
		return gopolutils.NewNamedException(gopolutils.OSError, "Can not get parent of filesystem root.")
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
