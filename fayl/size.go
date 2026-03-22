package fayl

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

// Representation of a file size.
type Size struct {
	size gopolutils.Size
	unit gopolutils.ByteSize
}

// Construct a new [Size] based on a [gopolutils.Size] and a given unit.
// Returns a new [Size] based on a [gopolutils.Size] and a given [gopolutils.ByteSize].
func NewSize(size gopolutils.Size, unit gopolutils.ByteSize) *Size {
	var result *Size = new(Size)
	result.size = size
	result.unit = unit
	return result
}

// Construct a new [Size] from the size in bytes.
// Returns a new [Size] based on a given size in bytes.
func SizeFromBytes(size gopolutils.Size) *Size {
	switch {
	case size >= gopolutils.Size(gopolutils.KB):
		return NewSize(size/gopolutils.Size(gopolutils.KB), gopolutils.KB)
	case size >= gopolutils.Size(gopolutils.MB):
		return NewSize(size/gopolutils.Size(gopolutils.MB), gopolutils.MB)
	case size >= gopolutils.Size(gopolutils.GB):
		return NewSize(size/gopolutils.Size(gopolutils.GB), gopolutils.GB)
	case size >= gopolutils.Size(gopolutils.TB):
		return NewSize(size/gopolutils.Size(gopolutils.TB), gopolutils.TB)
	case size >= gopolutils.Size(gopolutils.PB):
		return NewSize(size/gopolutils.Size(gopolutils.PB), gopolutils.PB)
	case size >= gopolutils.Size(gopolutils.EB):
		return NewSize(size/gopolutils.Size(gopolutils.EB), gopolutils.EB)
	default:
		return NewSize(size, gopolutils.Byte)
	}
}

// Obtain the unit of the size.
// Returns a [gopolutils.ByteSize] unit of the size.
func (size Size) Unit() gopolutils.ByteSize {
	return size.unit
}

// Obtain the size property of the size.
// Returns the size property of the size.
func (size Size) Size() gopolutils.Size {
	return size.size
}

// Determine if the size is empty.
// Returns true if the size of the size is equal to zero.
func (size Size) IsEmpty() bool {
	return size.size == 0
}

// Represent the size as a string.
// Returns a string representation of the size.
func (size Size) String() string {
	return fmt.Sprintf("%d%s", size.size, size.unit)
}
