package fayl

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/goserialize"
	"gopkg.in/yaml.v2"
)

const (
	// Default json file extension.
	//
	// Depricated: A new serialization library makes this obsolete.
	JSONType string = "json"
	// Default yaml file extenstion.
	//
	// Depricated: A new serialization library makes this obsolete.
	YAMLType string = "yaml"
	// Default toml file extension.
	//
	// Depricated: A new serialization library makes this obsolete.
	TOMLType string = "toml"
)

// Generic unmarshal type. The reader type takes in the raw file content and a pointer to the object.
//
// Depricated: A new serialization library makes this obsolete.
type Reader = func([]byte, any) error

var (
	// Default json reader.
	//
	// Depricated: A new serialization library makes this obsolete.
	JSONReader Reader = json.Unmarshal
	// Default yaml reader.
	//
	// Depricated: A new serialization library makes this obsolete.
	YAMLReader Reader = yaml.Unmarshal
	// Default toml reader.
	//
	// Depricated: A new serialization library makes this obsolete.
	TOMLReader Reader = toml.Unmarshal
)

// Helper function to get the file type of a given file path.
// Returns the file type of the given filepath based on its extension.
// If the file does not have a valid extension, or the extension can not be indexed, an Exception is returned with an empty string.
//
// Depricated: As of gopolutils 1.7.0, this method has been made obsolete. Use the fayl.Path.Suffix method instead.
func getFileType(filePath string) (string, *gopolutils.Exception) {
	var index int = strings.LastIndexByte(filePath, '.')
	if index < 0 {
		return "", gopolutils.NewException("Can not determine index of file extension.")
	}
	return filePath[index+1:], nil
}

// Read the raw contents of a file.
// Returns a byte slice representing the raw file content.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
func Read(filePath Path) ([]byte, *gopolutils.Exception) {
	var absolute *Path
	var absoluteError *gopolutils.Exception
	absolute, absoluteError = filePath.Absolute()
	if absoluteError != nil {
		return nil, gopolutils.NewNamedException("IOError", absoluteError.Message())
	}
	var file []byte
	var readError error
	file, readError = os.ReadFile(absolute.ToString())
	if readError != nil {
		return nil, gopolutils.NewNamedException("IOError", readError.Error())
	}
	return file, nil
}

// Helper method to marshall a single object from a file.
// Returns a pointer to the marshalled object type.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// If the given reader returns an error, an IOError is returned with a nil data pointer.
func readRawObject[Type any](filePath Path, reader goserialize.Reader) (*Type, *gopolutils.Exception) {
	var raw []byte
	var readError *gopolutils.Exception
	raw, readError = Read(filePath)
	if readError != nil {
		return nil, readError
	}
	var result *Type = new(Type)
	var marshallError error = reader(raw, result)
	if marshallError != nil {
		return nil, gopolutils.NewNamedException("IOError", marshallError.Error())
	}
	return result, nil
}

// Helper method to marshall a slice of objects from a file.
// Returns a pointer to the marshalled slice of objects.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// If the given reader returns an error, an IOError is returned with a nil data pointer.
func readRawList[Type any](filePath Path, reader goserialize.Reader) ([]Type, *gopolutils.Exception) {
	var raw []byte
	var readError *gopolutils.Exception
	raw, readError = Read(filePath)
	if readError != nil {
		return nil, readError
	}
	var result []Type = make([]Type, len(raw))
	var readerError error = reader(raw, &result)
	if readerError != nil {
		return nil, gopolutils.NewNamedException("IOError", readerError.Error())
	}
	return result, nil
}

// Helper function to dispatch the reader function based on the type of the given file path.
// Returns a slice of marshalled data from a file.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// Alternatively, if the data can not be marshalled, an IOError is returned with a nil data pointer.
// In addition, if the file type can not be evaluated, an Exception is returned with a nil data pointer.
func readerListDispatch[Type any](filePath Path) ([]Type, *gopolutils.Exception) {
	var fileType string
	var except *gopolutils.Exception
	fileType, except = filePath.Suffix()
	if except != nil {
		return nil, except
	}
	switch fileType {
	case goserialize.YAMLType:
		return readRawList[Type](filePath, goserialize.YAMLReader)
	case "yml":
		return readRawList[Type](filePath, goserialize.YAMLReader)
	case goserialize.TOMLType:
		return readRawList[Type](filePath, goserialize.TOMLReader)
	default:
		return readRawList[Type](filePath, goserialize.JSONReader)
	}
}

// Convert a slice to a collection.
//
// Depricated: A new serialization library makes this obsolete.
func sliceToCollection[Type any](items []Type, collection collections.Collection[Type]) {
	var item Type
	for _, item = range items {
		collection.Append(item)
	}
}

// Read a file as a view into a collection of objects.
// Returns a view into a collection of a file containing a list of data.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// Alternatively, if the data can not be marshalled, an IOError is returned with a nil data pointer.
// In addition, if the file type can not be evaluated, an Exception is returned with a nil data pointer.
func ReadList[Type any](filePath Path) (collections.View[Type], *gopolutils.Exception) {
	var raw []Type
	var rawError *gopolutils.Exception
	raw, rawError = readerListDispatch[Type](filePath)
	if rawError != nil {
		return nil, rawError
	}
	var result collections.Collection[Type] = collections.NewArray[Type]()
	sliceToCollection[Type](raw, result)
	return result, nil
}

// Read a file as an object.
// Returns a pointer to an object of data from a file.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// Alternatively, if the data can not be marshalled, an IOError is returned with a nil data pointer.
// In addition, if the file type can not be evaluated, an Exception is returned with a nil data pointer.
func ReadObject[Type any](filePath Path) (*Type, *gopolutils.Exception) {
	var fileType string
	var except *gopolutils.Exception
	fileType, except = filePath.Suffix()
	if except != nil {
		return nil, except
	}
	switch fileType {
	case goserialize.YAMLType:
		return readRawObject[Type](filePath, goserialize.YAMLReader)
	case "yml":
		return readRawObject[Type](filePath, goserialize.YAMLReader)
	case goserialize.TOMLType:
		return readRawObject[Type](filePath, goserialize.TOMLReader)
	default:
		return readRawObject[Type](filePath, goserialize.JSONReader)
	}
}
