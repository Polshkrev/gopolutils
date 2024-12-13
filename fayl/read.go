package fayl

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Polshkrev/gopolutils/collections"
)

// Read the raw contents of a file.
// Returns a byte slice representing the raw file content.
func ReadFile(fileName string) []byte {
	var file []byte
	var readError error
	file, readError = os.ReadFile(fileName)
	if readError != nil {
		fmt.Fprintf(os.Stderr, "%s\n", readError.Error())
		os.Exit(1)
	}
	return file
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
func ReadJSONList[Type any](fileName string) collections.View[Type] {
	var raw []byte = ReadFile(fileName)
	var result collections.Collection[Type] = collections.NewArray[Type]()
	var rawList []Type = make([]Type, 0)
	var readError error = json.Unmarshal(raw, &rawList)
	if readError != nil {
		fmt.Fprintf(os.Stderr, "%s\n", readError.Error())
		os.Exit(1)
	}
	sliceToCollection[Type](rawList, result)
	return result
}

// Read a json file as a view.
// Returns a view into a json file containing an object of data.
func ReadJSONObject[Type any](fileName string) *Type {
	var raw []byte = ReadFile(fileName)
	var rawObject *Type = new(Type)
	var readError error = json.Unmarshal(raw, rawObject)
	if readError != nil {
		fmt.Fprintf(os.Stderr, "%s\n", readError.Error())
		os.Exit(1)
	}
	return rawObject
}
