package fayl

import (
	"os"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/goserialize"
)

// Concurrently write a slice of bytes to a file of a given path.
func writeConcurrent(path *Path, content []byte, errorChannel chan<- error) {
	errorChannel <- os.WriteFile(path.ToString(), content, 0644)
	defer close(errorChannel)
}

// Write a slice of bytes to a file.
// If the file can not be written, an IOError is returned.
func Write(filePath *Path, content []byte) *gopolutils.Exception {
	var errorChannel chan error = make(chan error, 1)
	go writeConcurrent(filePath, content, errorChannel)
	var writeError error = <-errorChannel
	if writeError != nil {
		return gopolutils.NewNamedException("IOError", writeError.Error())
	}
	return nil
}

// Helper method to marshall a single object to a file.
// If the file can not be written, an IOError is returned.
// If the given writer returns an error, an IOError is returned.
func writeRawObject[Type any](filePath *Path, content *Type, writer goserialize.Writer) *gopolutils.Exception {
	var data []byte
	var marshalError error
	data, marshalError = writer(content)
	if marshalError != nil {
		return gopolutils.NewNamedException("IOError", marshalError.Error())
	}
	return Write(filePath, data)
}

// Helper method to marshall a slice of objects to a file.
// If the file can not be written, an IOError is returned.
// If the given writer returns an error, an IOError is returned.
func writeRawList[Type any](filePath *Path, content []Type, writer goserialize.Writer) *gopolutils.Exception {
	var data []byte
	var marshalError error
	data, marshalError = writer(content)
	if marshalError != nil {
		return gopolutils.NewNamedException("IOError", marshalError.Error())
	}
	return Write(filePath, data)
}

// Write a view of a type into a file.
// If the file can not be written, an IOError is returned.
// Alternatively, if the data can not be marshalled, an IOError is returned.
// In addition, if the file type can not be evaluated, an Exception is returned.
func WriteList[Type any](filePath *Path, content collections.View[Type]) *gopolutils.Exception {
	var fileType string
	var except *gopolutils.Exception
	fileType, except = filePath.Suffix()
	if except != nil {
		return except
	}
	switch fileType {
	case goserialize.YAMLType:
		return writeRawList[Type](filePath, content.Collect(), goserialize.YAMLWriter)
	case "yml":
		return writeRawList[Type](filePath, content.Collect(), goserialize.YAMLWriter)
	case goserialize.TOMLType:
		return writeRawList[Type](filePath, content.Collect(), goserialize.TOMLWriter)
	default:
		return writeRawList[Type](filePath, content.Collect(), goserialize.JSONWriter)
	}
}

// Write a file as an object.
// If the file can not be written, an IOError is returned.
// Alternatively, if the data can not be marshalled, an IOError is returned.
// In addition, if the file type can not be evaluated, an Exception is returned.
func WriteObject[Type any](filePath *Path, content *Type) *gopolutils.Exception {
	var fileType string
	var except *gopolutils.Exception
	fileType, except = filePath.Suffix()
	if except != nil {
		return except
	}
	switch fileType {
	case goserialize.YAMLType:
		return writeRawObject[Type](filePath, content, goserialize.YAMLWriter)
	case "yml":
		return writeRawObject[Type](filePath, content, goserialize.YAMLWriter)
	case goserialize.TOMLType:
		return writeRawObject[Type](filePath, content, goserialize.TOMLWriter)
	default:
		return writeRawObject[Type](filePath, content, goserialize.JSONWriter)
	}
}
