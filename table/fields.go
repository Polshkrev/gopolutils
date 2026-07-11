package table

import (
	"strings"

	"github.com/Polshkrev/gopolutils"
)

// Representation of a database field.
type Field gopolutils.StringEnum

const (
	Id Field = "id" // Default id field.
)

// Represent a field as a string.
// Returns a string representation of a field.
func (field Field) String() string {
	return string(field)
}

// Represent a variadic amount of [Field]s as a slice of string.
// Represents a variadic amount of [Field]s as a slice of string.
func fieldsToString(fields ...Field) []string {
	var result []string = make([]string, 0)
	var i int
	for i = range fields {
		var field Field = fields[i]
		result = append(result, field.String())
	}
	return result
}

// Represent a variadic amount of [Field]s as a slice of string.
// Represents a variadic amount of [Field]s as a slice of string.
func GetFields(fields ...Field) string {
	return strings.Join(fieldsToString(fields...), ", ")
}
