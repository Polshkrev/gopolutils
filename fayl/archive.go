package fayl

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Polshkrev/gopolutils"
)

// Extract an either a `.tar`, `.tar.gz`, or `.zip` archive based on a given source [Path].
// If the [Suffix] of the source [Path] can not be obtained, [gopolutils.KeyError] is returned.
// If the file can not be read, an [gopolutils.IOError] is returned.
// If the file can not be written, an [gopolutils.IOError] is returned.
func Extract(source, destination *Path) *gopolutils.Exception {
	var sourceSuffix Suffix
	var except *gopolutils.Exception
	sourceSuffix, except = source.Suffix()
	if except != nil {
		return except
	}
	switch sourceSuffix {
	case Gz:
		var newRawPath string = strings.TrimSuffix(source.ToString(), filepath.Ext(source.ToString()))
		var newPath *Path = PathFrom(newRawPath)
		var content []byte
		var readError *gopolutils.Exception
		content, readError = Read(source)
		if readError != nil {
			return readError
		}
		var raw []byte
		var rawError *gopolutils.Exception
		raw, rawError = Gunzip(content)
		if rawError != nil {
			return rawError
		}
		var writeError *gopolutils.Exception = Write(newPath, raw)
		if writeError != nil {
			return writeError
		}
		return Extract(newPath, destination)
	case Tar:
		return Untar(source, destination)
	default:
		return Unzip(source, destination)
	}
}

// Archive a given variadic group of file [Entry] to a given destination [Path].
// If the [Suffix] of the source [Path] can not be obtained, [gopolutils.KeyError] is returned.
// If the file can not be read, an [gopolutils.IOError] is returned.
// If the file can not be written, an [gopolutils.IOError] is returned.
func Archive(destination *Path, files ...*Entry) *gopolutils.Exception {
	var destinationSuffix Suffix
	var except *gopolutils.Exception
	destinationSuffix, except = destination.Suffix()
	if except != nil {
		return except
	}
	switch destinationSuffix {
	case Gz:
		var newRawPath string = strings.TrimSuffix(destination.ToString(), filepath.Ext(destination.ToString()))
		var newPath *Path = PathFrom(newRawPath)
		var archiveError *gopolutils.Exception = Archive(newPath, files...)
		if archiveError != nil {
			return archiveError
		}
		var content []byte
		var readError *gopolutils.Exception
		content, readError = Read(newPath)
		if readError != nil {
			return readError
		}
		var raw []byte
		var rawError *gopolutils.Exception
		raw, rawError = Gzip(content)
		if rawError != nil {
			return rawError
		}
		var writeError *gopolutils.Exception = Write(destination, raw)
		if writeError != nil {
			return writeError
		}
		return nil
	case Tar:
		return TarFolder(destination, files...)
	default:
		return ZipFolder(destination, files...)
	}
}

// Zip a given variadic group of [Entry] to a given destination [Path].
// If the given destination [Path] does not exist, a [gopolutils.FileExistsError] is returned.
// If the file can not be created, an [gopolutils.OSError] is returned.
// If a handle to any of the given files can not be obtained, an [gopolutils.IOError] is returned.
// If any of the files can not be added to the zip archive, an [gopolutils.IOError] is returned.
func ZipFolder(destination *Path, files ...*Entry) *gopolutils.Exception {
	if destination.Exists() {
		return gopolutils.NewNamedException(gopolutils.FileExistsError, "'%s' already exists.", destination.ToString())
	}
	var target *os.File
	var createError *gopolutils.Exception
	target, createError = createFile(destination.ToString())
	if createError != nil {
		return createError
	}
	var writer *zip.Writer = zip.NewWriter(target)
	var i int
	for i = range files {
		var file *Entry = files[i]
		if file.Is(DirectoryType) {
			continue
		}
		var name string = file.Path().ToString()
		var handle *os.File
		var openError *gopolutils.Exception
		handle, openError = getHandle(name)
		if openError != nil {
			return openError
		}
		defer handle.Close()
		var cleaned string
		var cleanedError error
		cleaned, cleanedError = getRelative(name)
		if cleanedError != nil {
			return gopolutils.NewNamedException(gopolutils.ValueError, cleanedError.Error())
		}
		var stripped string = stripPrefix(cleaned, fmt.Sprintf("..%c", filepath.Separator))
		var destinationHandle io.Writer
		var handleError error
		destinationHandle, handleError = writer.Create(stripped)
		if handleError != nil {
			return gopolutils.NewNamedException(gopolutils.IOError, handleError.Error())
		}
		var copyError *gopolutils.Exception = copyFile(destinationHandle, handle)
		if copyError != nil {
			return copyError
		}
	}
	var closeError error = writer.Close()
	if closeError != nil {
		return gopolutils.NewNamedException(gopolutils.IOError, closeError.Error())
	}
	return nil
}

// Unzip a folder from a given source [Path] to a given destination [Path].
// If the source [Path] does not exist, a [gopolutils.FileNotFoundError] is returned.
// If the destination [Path] already exists, a [gopolutils.FileExistsError] is returned.
// If file can not be opened, a [gopolutils.ValueError] is returned.
// If the destination path is not valid, a [gopolutils.ValueError] is returned.
// If a handle to the any of the read files can not be obtained, an [gopolutils.OSError] is returned.
// If any of the files can not be opened, an [gopolutils.IOError] is returned.
// If any of the files can not be copied, an [gopolutils.IOError] is returned.
func Unzip(source, destination *Path) *gopolutils.Exception {
	if !source.Exists() {
		return gopolutils.NewNamedException(gopolutils.FileNotFoundError, "'%s' does not exist.", source.ToString())
	} else if destination.Exists() {
		gopolutils.NewNamedException(gopolutils.FileExistsError, "'%s' already exists.", destination.ToString())
	}
	var makeDirectoryError *gopolutils.Exception = makeDirectory(destination.ToString())
	if makeDirectoryError != nil {
		return makeDirectoryError
	}
	var reader *zip.ReadCloser
	var readerError error
	reader, readerError = zip.OpenReader(source.ToString())
	if readerError != nil {
		return gopolutils.NewNamedException(gopolutils.OSError, readerError.Error())
	}
	defer reader.Close()
	var i int
	for i = range reader.File {
		var file *zip.File = reader.File[i]
		var fullPath string = fmt.Sprintf("%s%c%s", destination.ToString(), filepath.Separator, file.Name)
		if !validate(fullPath, destination.ToString()) {
			return gopolutils.NewNamedException(gopolutils.ValueError, "Invalid file path: %s", fullPath)
		} else if file.FileInfo().IsDir() {
			var makeDirectoryError *gopolutils.Exception = makeDirectory(fullPath)
			if makeDirectoryError != nil {
				return makeDirectoryError
			}
			continue
		}
		var makeDirectoryError *gopolutils.Exception = makeDirectory(filepath.Dir(fullPath))
		if makeDirectoryError != nil {
			return makeDirectoryError
		}
		var fileHandle *os.File
		var handleError *gopolutils.Exception
		fileHandle, handleError = createFile(fullPath)
		if handleError != nil {
			return handleError
		}
		defer fileHandle.Close()
		var readCloser io.ReadCloser
		var openError error
		readCloser, openError = file.Open()
		if openError != nil {
			return gopolutils.NewNamedException(gopolutils.IOError, openError.Error())
		}
		defer readCloser.Close()
		var copyError *gopolutils.Exception = copyFile(fileHandle, readCloser)
		if copyError != nil {
			return copyError
		}
	}
	return nil
}

// Unzip a tar file from a given source [Path] to a given destination [Path].
// If the source [Path] does not exist, a [gopolutils.FileNotFoundError] is returned.
// If the destination [Path] already exists, a [gopolutils.FileExistsError] is returned.
// If the destination can not be created, an [gopolutils.OSError] is returned.
// If source file can not be opened, an [gopolutils.OSError] is returned.
// If the destination path is not valid, a [gopolutils.ValueError] is returned.
// If any of the files can not be opened, an [gopolutils.IOError] is returned.
// If any of the files can not be copied, an [gopolutils.IOError] is returned.
func Untar(source, destination *Path) *gopolutils.Exception {
	if !source.Exists() {
		return gopolutils.NewNamedException(gopolutils.FileNotFoundError, "'%s' does not exist.", source.ToString())
	} else if destination.Exists() {
		gopolutils.NewNamedException(gopolutils.FileExistsError, "'%s' already exists.", destination.ToString())
	}
	var makeDirectoryError *gopolutils.Exception = makeDirectory(destination.ToString())
	if makeDirectoryError != nil {
		return makeDirectoryError
	}
	var handle *os.File
	var handleError *gopolutils.Exception
	handle, handleError = getHandle(source.ToString())
	if handleError != nil {
		return gopolutils.NewNamedException(gopolutils.OSError, handleError.Error())
	}
	var tarReader *tar.Reader = tar.NewReader(handle)
	for {
		var header *tar.Header
		var nextError error
		header, nextError = tarReader.Next()
		if header == nil {
			return nil
		} else if nextError != nil {
			return gopolutils.NewNamedException(gopolutils.IOError, nextError.Error())
		} else if nextError == io.EOF {
			return nil
		}
		var fullPath string
		if !filepath.IsAbs(header.Name) {
			fullPath = fmt.Sprintf("%s%c%s", destination.ToString(), filepath.Separator, header.Name)
		} else {
			var cleaned string
			var cleanedError error
			cleaned, cleanedError = cleanPath(header.Name)
			if cleanedError != nil {
				return gopolutils.NewNamedException(gopolutils.IOError, cleanedError.Error())
			}
			fullPath = fmt.Sprintf("%s%c%s", destination.ToString(), filepath.Separator, cleaned)
		}
		if !validate(fullPath, destination.ToString()) {
			return gopolutils.NewNamedException(gopolutils.ValueError, "Invalid file path: %s", fullPath)
		}
		switch header.Typeflag {
		case tar.TypeDir:
			var makeDirectoryError *gopolutils.Exception = makeDirectory(fullPath)
			if makeDirectoryError != nil {
				return makeDirectoryError
			}
		case tar.TypeReg:
			var except *gopolutils.Exception = untarFile(fullPath, tarReader)
			if except != nil {
				return except
			}
		}
	}
}

// Zip a given variadic group of [Entry] into a destination tar [Path].
// If the destination [Path] already exists, a [gopolutils.FileExistsError] is returned.
// If the file can not be created, an [gopolutils.OSError] is returned.
// If the tar header can not be created or written, an [gopolutils.IOError] is returned.
// If a handle to any of the given files can not be obtained, an [gopolutils.OSError] is returned.
// If any of the files can not be copied, an [gopolutils.IOError] is returned.
func TarFolder(destination *Path, files ...*Entry) *gopolutils.Exception {
	if destination.Exists() {
		return gopolutils.NewNamedException(gopolutils.FileExistsError, "'%s' already exists.", destination.ToString())
	}
	var targetHandle *os.File
	var createError *gopolutils.Exception
	targetHandle, createError = createFile(destination.ToString())
	if createError != nil {
		return createError
	}
	var tarWriter *tar.Writer = tar.NewWriter(targetHandle)
	defer tarWriter.Close()
	var i int
	for i = range files {
		var file *Entry = files[i]
		if file.Is(DirectoryType) {
			continue
		}
		var name string = file.Path().ToString()
		var info fs.FileInfo
		var statError error
		info, statError = os.Stat(name)
		if statError != nil {
			return gopolutils.NewNamedException(gopolutils.OSError, statError.Error())
		}
		var header *tar.Header
		var headerError error
		header, headerError = tar.FileInfoHeader(info, name)
		if headerError != nil {
			return gopolutils.NewNamedException(gopolutils.IOError, headerError.Error())
		}
		var cleaned string
		var cleanedError error
		cleaned, cleanedError = getRelative(name)
		if cleanedError != nil {
			return gopolutils.NewNamedException(gopolutils.ValueError, cleanedError.Error())
		}
		var stripped string = stripPrefix(cleaned, fmt.Sprintf("..%c", filepath.Separator))
		header.Name = stripped
		var writeHeaderError error = tarWriter.WriteHeader(header)
		if writeHeaderError != nil {
			return gopolutils.NewNamedException(gopolutils.IOError, writeHeaderError.Error())
		}
		var openFile *os.File
		var openError *gopolutils.Exception
		openFile, openError = getHandle(name)
		if openError != nil {
			return gopolutils.NewNamedException(gopolutils.OSError, openError.Error())
		}
		defer openFile.Close()
		var copyError *gopolutils.Exception = copyFile(tarWriter, openFile)
		if copyError != nil {
			return copyError
		}
	}
	return nil
}

// Validate a given source and destination string paths to prevent a "Zip-Slip" vulnerability.
// Returns true if the source and destination string paths are valid, else false.
func validate(source, destination string) bool {
	return strings.HasPrefix(source, filepath.Clean(destination)+string(os.PathSeparator))
}

// Concurrently create a directory on the filesystem.
// If the directory can not be created, an [gopolutils.OSError] is sent to the exception channel.
func makeDirectoryConcurrent(target string, exceptionChannel chan<- *gopolutils.Exception) {
	defer close(exceptionChannel)
	var makeDirectoryError error = os.MkdirAll(target, 0755)
	if makeDirectoryError != nil {
		exceptionChannel <- gopolutils.NewNamedException(gopolutils.OSError, makeDirectoryError.Error())
		return
	}
	exceptionChannel <- nil
}

// Create a directory on the filesystem.
// If the directory can not be created, an [gopolutils.OSError] returned.
func makeDirectory(target string) *gopolutils.Exception {
	var exceptionChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception, 1)
	go makeDirectoryConcurrent(target, exceptionChannel)
	var except *gopolutils.Exception = <-exceptionChannel
	return except
}

// Concurrently create a file on the filesystem.
// If the file can not be created, an [gopolutils.OSError] is sent to the exceptionChannel.
func createConcurrent(target string, resultChannel chan<- *os.File, exceptionChannel chan<- *gopolutils.Exception) {
	defer close(resultChannel)
	defer close(exceptionChannel)
	var result *os.File
	var createError error
	result, createError = os.Create(target)
	if createError != nil {
		resultChannel <- nil
		exceptionChannel <- gopolutils.NewNamedException(gopolutils.OSError, createError.Error())
		return
	}
	resultChannel <- result
	exceptionChannel <- nil
}

// Create a file on the filesystem.
// Returns a handle to the created file.
// If the file can not be created, an [gopolutils.OSError] is returned.
func createFile(target string) (*os.File, *gopolutils.Exception) {
	var resultChannel chan *os.File = make(chan *os.File, 1)
	var exceptionChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception, 1)
	go createConcurrent(target, resultChannel, exceptionChannel)
	var result *os.File = <-resultChannel
	var except *gopolutils.Exception = <-exceptionChannel
	return result, except
}

// Concurrently copy a file.
// If the file can not be copied, an [gopolutils.IOError] is sent to the exception channel.
func copyConcurrent(destination io.Writer, source io.Reader, exceptionChannel chan<- *gopolutils.Exception) {
	defer close(exceptionChannel)
	var copyError error
	_, copyError = io.Copy(destination, source)
	if copyError != nil {
		exceptionChannel <- gopolutils.NewNamedException(gopolutils.IOError, copyError.Error())
		return
	}
	exceptionChannel <- nil
}

// Copy a file on the filesystem.
// If the file can not be copied, an [gopolutils.IOError] returned.
func copyFile(destination io.Writer, source io.Reader) *gopolutils.Exception {
	var exceptionChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception, 1)
	go copyConcurrent(destination, source, exceptionChannel)
	var except *gopolutils.Exception = <-exceptionChannel
	return except
}

// Concurrently open a file.
func openConcurrent(path string, handleChannel chan<- *os.File, exceptionChannel chan<- *gopolutils.Exception) {
	defer close(handleChannel)
	defer close(exceptionChannel)
	var openFile *os.File
	var openError error
	openFile, openError = os.Open(path)
	if openError != nil {
		handleChannel <- nil
		exceptionChannel <- gopolutils.NewNamedException(gopolutils.OSError, openError.Error())
		return
	}
	handleChannel <- openFile
	exceptionChannel <- nil
}

// Obtain a handle to a file from a given path.
// Returns a handle to the open file of the given path.
// If the handle can not be obtained, an [gopolutils.OSError] is returned.
func getHandle(path string) (*os.File, *gopolutils.Exception) {
	var handleChannel chan *os.File = make(chan *os.File, 1)
	var exceptionChannel chan *gopolutils.Exception = make(chan *gopolutils.Exception, 1)
	go openConcurrent(path, handleChannel, exceptionChannel)
	var handle *os.File = <-handleChannel
	var except *gopolutils.Exception = <-exceptionChannel
	return handle, except
}

// Strip the given prefix from the given if contains the given prefix.
// Returns the given path without the preceding prefix.
func stripPrefix(path string, prefix string) string {
	if strings.HasPrefix(path, prefix) {
		return strings.TrimPrefix(path, prefix)
	}
	return path
}

// Clean a given path.
// Returns a cleaned path.
func cleanPath(path string) (string, error) {
	var relativeName string
	var relativeError error
	relativeName, relativeError = getRelative(path)
	if relativeError != nil {
		return "", relativeError
	}
	var evaluated string
	var evaluatedError error
	evaluated, evaluatedError = filepath.EvalSymlinks(relativeName)
	if evaluatedError != nil {
		return "", evaluatedError
	}
	return stripPrefix(evaluated, fmt.Sprintf("..%c", filepath.Separator)), nil
}

// Obtain the relative path from the given path parametre.
// Returns the given path relative to the current working directory.
func getRelative(path string) (string, error) {
	if !filepath.IsAbs(path) {
		return path, nil
	}
	var current string
	var currentError error
	current, currentError = os.Getwd()
	if currentError != nil {
		return "", currentError
	}
	return filepath.Rel(current, path)
}

// Untar a single target file from its given [os.FileMode] and [tar.Reader].
// If the directory can not be created, an [gopolutils.OSError] returned.
// If the file can not be copied, an [gopolutils.IOError] returned.
func untarFile(target string, reader *tar.Reader) *gopolutils.Exception {
	var makeDirectoryError *gopolutils.Exception = makeDirectory(filepath.Dir(target))
	if makeDirectoryError != nil {
		return makeDirectoryError
	}
	var file *os.File
	var openFileError *gopolutils.Exception
	file, openFileError = createFile(target)
	if openFileError != nil {
		return openFileError
	}
	defer file.Close()
	return copyFile(file, reader)
}
