package fayl

import (
	"os"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/goserialize"
)

// Concurrently read a file passed on its given path.
func readConcurrent(path *Path, dataChannel chan<- []byte, errorChannel chan<- error) {
	var data []byte
	var readError error
	data, readError = os.ReadFile(path.ToString())
	dataChannel <- data
	errorChannel <- readError
	defer close(dataChannel)
	defer close(errorChannel)
}

// Read the raw contents of a file.
// Returns a byte slice representing the raw file content.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
func Read(filePath *Path) ([]byte, *gopolutils.Exception) {
	var dataChannel chan []byte = make(chan []byte, 1)
	var errorChannel chan error = make(chan error, 1)
	go readConcurrent(filePath, dataChannel, errorChannel)
	var file []byte = <-dataChannel
	var readError error = <-errorChannel
	if readError != nil {
		return nil, gopolutils.NewNamedException(gopolutils.IOError, readError.Error())
	}
	return file, nil
}

// Helper method to marshall a single object from a file.
// Returns a pointer to the marshalled object type.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// If the given reader returns an error, an IOError is returned with a nil data pointer.
func readRawObject[Type any](filePath *Path, reader goserialize.Reader) (*Type, *gopolutils.Exception) {
	var raw []byte
	var readError *gopolutils.Exception
	raw, readError = Read(filePath)
	if readError != nil {
		return nil, readError
	}
	var result *Type = new(Type)
	var marshallError error = reader(raw, result)
	if marshallError != nil {
		return nil, gopolutils.NewNamedException(gopolutils.IOError, marshallError.Error())
	}
	return result, nil
}

// Helper method to marshall a slice of objects from a file.
// Returns a pointer to the marshalled slice of objects.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// If the given reader returns an error, an IOError is returned with a nil data pointer.
func readRawList[Type any](filePath *Path, reader goserialize.Reader) ([]Type, *gopolutils.Exception) {
	var raw []byte
	var readError *gopolutils.Exception
	raw, readError = Read(filePath)
	if readError != nil {
		return nil, readError
	}
	var result []Type = make([]Type, len(raw))
	var readerError error = reader(raw, &result)
	if readerError != nil {
		return nil, gopolutils.NewNamedException(gopolutils.IOError, readerError.Error())
	}
	return result, nil
}

// Helper function to dispatch the reader function based on the type of the given file path.
// Returns a slice of marshalled data from a file.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// Alternatively, if the data can not be marshalled, an IOError is returned with a nil data pointer.
// In addition, if the file type can not be evaluated, an Exception is returned with a nil data pointer.
func readerListDispatch[Type any](filePath *Path) ([]Type, *gopolutils.Exception) {
	var fileType Suffix
	var except *gopolutils.Exception
	fileType, except = filePath.Suffix()
	if except != nil {
		return nil, except
	}
	switch fileType {
	case Yaml:
		return readRawList[Type](filePath, goserialize.YAMLReader)
	case Toml:
		return readRawList[Type](filePath, goserialize.TOMLReader)
	case Csv:
		return readRawList[Type](filePath, goserialize.CSVReader)
	default:
		return readRawList[Type](filePath, goserialize.JSONReader)
	}
}

// Read a file as a view into a collection of objects.
// Returns a view into a collection of a file containing a list of data.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// Alternatively, if the data can not be marshalled, an IOError is returned with a nil data pointer.
// In addition, if the file type can not be evaluated, an Exception is returned with a nil data pointer.
func ReadList[Type any](filePath *Path) (collections.View[Type], *gopolutils.Exception) {
	var raw []Type
	var rawError *gopolutils.Exception
	raw, rawError = readerListDispatch[Type](filePath)
	if rawError != nil {
		return nil, rawError
	}
	var result collections.View[Type] = goserialize.SliceToView[Type](raw)
	return result, nil
}

// Read a file as an object.
// Returns a pointer to an object of data from a file.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// Alternatively, if the data can not be marshalled, an IOError is returned with a nil data pointer.
// In addition, if the file type can not be evaluated, an Exception is returned with a nil data pointer.
func ReadObject[Type any](filePath *Path) (*Type, *gopolutils.Exception) {
	var fileType Suffix
	var except *gopolutils.Exception
	fileType, except = filePath.Suffix()
	if except != nil {
		return nil, except
	}
	switch fileType {
	case Yaml:
		return readRawObject[Type](filePath, goserialize.YAMLReader)
	case Toml:
		return readRawObject[Type](filePath, goserialize.TOMLReader)
	case Csv:
		return readRawObject[Type](filePath, goserialize.CSVReader)
	default:
		return readRawObject[Type](filePath, goserialize.JSONReader)
	}
}
