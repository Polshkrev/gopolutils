package gopolutils

// Standardization of a enum value.
type Enum uint8

// Standardization of a enum string value.
type StringEnum string

// Standardization of a size of a type or collection.
type Size uint64

// Representation of a generic number type.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~complex64 | ~complex128
}

// Standardization of an empty ptyhon-like None type.
type None struct{}

// Standardization of a default callback.
type Callback func()
