package fayl

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

// Read the raw contents of a file.
// Returns a byte slice representing the raw file content.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
func ReadFile(filePath string) ([]byte, *gopolutils.Exception) {
	var absolute string
	var absoluteError error
	absolute, absoluteError = filepath.Abs(filePath)
	if absoluteError != nil {
		return nil, gopolutils.NewNamedException("IOError", absoluteError.Error())
	}
	var file []byte
	var readError error
	file, readError = os.ReadFile(absolute)
	if readError != nil {
		return nil, gopolutils.NewNamedException("IOError", readError.Error())
	}
	return file, nil
}

// Convert a slice to a collection.
func sliceToCollection[Type any](items []Type, collection collections.Collection[Type]) {
	var item Type
	for _, item = range items {
		collection.Append(item)
	}
}

// Read a json file as a view.
// Returns a view into a json file containing a list of data.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// Alternatively, if the data can not be marshalled, an IOError is returned with a nil data pointer.
func ReadJSONList[Type any](filePath string) (collections.View[Type], *gopolutils.Exception) {
	var raw []byte
	var rawError *gopolutils.Exception
	raw, rawError = ReadFile(filePath)
	if rawError != nil {
		return nil, rawError
	}
	var result collections.Collection[Type] = collections.NewArray[Type]()
	var rawList []Type = make([]Type, 0)
	var readError error = json.Unmarshal(raw, &rawList)
	if readError != nil {
		return nil, gopolutils.NewNamedException("IOError", readError.Error())
	}
	sliceToCollection[Type](rawList, result)
	return result, nil
}

// Read a json file as a view.
// Returns a pointer to an object of data from a json file.
// If the absolute path of the file can not be obtained, or the file can not be read, an IOError is returned with a nil data pointer.
// Alternatively, if the data can not be marshalled, an IOError is returned with a nil data pointer.
func ReadJSONObject[Type any](filePath string) (*Type, *gopolutils.Exception) {
	var raw []byte
	var rawError *gopolutils.Exception
	raw, rawError = ReadFile(filePath)
	if rawError != nil {
		return nil, rawError
	}
	var rawObject *Type = new(Type)
	var readError error = json.Unmarshal(raw, rawObject)
	if readError != nil {
		return nil, rawError
	}
	return rawObject, nil
}
