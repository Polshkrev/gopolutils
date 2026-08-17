package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
)

func TestNewSize(t *testing.T) {
	var size *fayl.Size = fayl.NewSize(42, gopolutils.MB)

	if size == nil {
		t.Fatal("expected non-nil Size")
	} else if size.Size() != 42 {
		t.Fatalf("expected size 42, got %d", size.Size())
	} else if size.Unit() != gopolutils.MB {
		t.Fatalf("expected unit %v, got %v", gopolutils.MB, size.Unit())
	}
}

type sizeFromBytesTest struct {
	name         string
	bytes        gopolutils.Size
	expectedSize gopolutils.Size
	expectedUnit gopolutils.ByteSize
}

func TestSizeFromBytes(t *testing.T) {
	var tests []sizeFromBytesTest = []sizeFromBytesTest{
		{
			name:         "Zero",
			bytes:        0,
			expectedSize: 0,
			expectedUnit: gopolutils.Byte,
		},
		{
			name:         "Bytes",
			bytes:        512,
			expectedSize: 512,
			expectedUnit: gopolutils.Byte,
		},
		{
			name:         "OneKB",
			bytes:        gopolutils.Size(gopolutils.KB),
			expectedSize: 1,
			expectedUnit: gopolutils.KB,
		},
		{
			name:         "TwoKB",
			bytes:        2 * gopolutils.Size(gopolutils.KB),
			expectedSize: 2,
			expectedUnit: gopolutils.KB,
		},
		{
			name:         "OneMB",
			bytes:        gopolutils.Size(gopolutils.MB),
			expectedSize: 1,
			expectedUnit: gopolutils.MB,
		},
		{
			name:         "OneGB",
			bytes:        gopolutils.Size(gopolutils.GB),
			expectedSize: 1,
			expectedUnit: gopolutils.GB,
		},
		{
			name:         "OneTB",
			bytes:        gopolutils.Size(gopolutils.TB),
			expectedSize: 1,
			expectedUnit: gopolutils.TB,
		},
		{
			name:         "OnePB",
			bytes:        gopolutils.Size(gopolutils.PB),
			expectedSize: 1,
			expectedUnit: gopolutils.PB,
		},
		{
			name:         "OneEB",
			bytes:        gopolutils.Size(gopolutils.EB),
			expectedSize: 1,
			expectedUnit: gopolutils.EB,
		},
	}
	var i int
	for i = range tests {
		var test sizeFromBytesTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			var size *fayl.Size = fayl.SizeFromBytes(test.bytes)

			if size.Size() != test.expectedSize {
				t.Fatalf("expected size %d, got %d", test.expectedSize, size.Size())
			} else if size.Unit() != test.expectedUnit {
				t.Fatalf("expected unit %v, got %v", test.expectedUnit, size.Unit())
			}
		})
	}
}

func TestSize(t *testing.T) {
	var size *fayl.Size = fayl.NewSize(123, gopolutils.GB)
	var got gopolutils.Size = size.Size()
	if got != 123 {
		t.Fatalf("expected 123, got %d", got)
	}
}

func TestUnit(t *testing.T) {
	var size *fayl.Size = fayl.NewSize(123, gopolutils.GB)
	var got gopolutils.ByteSize = size.Unit()
	if got != gopolutils.GB {
		t.Fatalf("expected %v, got %v", gopolutils.GB, got)
	}
}

type IsEmptyTest struct {
	name     string
	size     *fayl.Size
	expected bool
}

func TestIsEmpty(t *testing.T) {
	var tests []IsEmptyTest = []IsEmptyTest{
		{
			name:     "Empty",
			size:     fayl.NewSize(0, gopolutils.Byte),
			expected: true,
		},
		{
			name:     "NonEmpty",
			size:     fayl.NewSize(1, gopolutils.Byte),
			expected: false,
		},
		{
			name:     "Large",
			size:     fayl.NewSize(100, gopolutils.GB),
			expected: false,
		},
	}
	var i int
	for i = range tests {
		var test IsEmptyTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			var got bool = test.size.IsEmpty()
			if got != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

type sizeStringTest struct {
	name     string
	size     *fayl.Size
	expected string
}

func TestString(t *testing.T) {
	var tests []sizeStringTest = []sizeStringTest{
		{
			name:     "Bytes",
			size:     fayl.NewSize(123, gopolutils.Byte),
			expected: "123B",
		},
		{
			name:     "KB",
			size:     fayl.NewSize(4, gopolutils.KB),
			expected: "4KB",
		},
		{
			name:     "MB",
			size:     fayl.NewSize(12, gopolutils.MB),
			expected: "12MB",
		},
		{
			name:     "GB",
			size:     fayl.NewSize(2, gopolutils.GB),
			expected: "2GB",
		},
		{
			name:     "TB",
			size:     fayl.NewSize(8, gopolutils.TB),
			expected: "8TB",
		},
	}
	var i int
	for i = range tests {
		var test sizeStringTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			var got string = test.size.String()
			if got != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, got)
			}
		})
	}
}

type sizeFromBytesBoundryTest struct {
	name         string
	bytes        gopolutils.Size
	expectedSize gopolutils.Size
	expectedUnit gopolutils.ByteSize
}

func TestSizeFromBytesBoundaryValues(t *testing.T) {
	var tests []sizeFromBytesBoundryTest = []sizeFromBytesBoundryTest{
		{
			name:         "KBMinusOne",
			bytes:        gopolutils.Size(gopolutils.KB) - 1,
			expectedSize: gopolutils.Size(gopolutils.KB) - 1,
			expectedUnit: gopolutils.Byte,
		},
		{
			name:         "MBMinusOne",
			bytes:        gopolutils.Size(gopolutils.MB) - 1,
			expectedSize: 1023,
			expectedUnit: gopolutils.KB,
		},
		{
			name:         "GBMinusOne",
			bytes:        gopolutils.Size(gopolutils.GB) - 1,
			expectedSize: 1023,
			expectedUnit: gopolutils.MB,
		},
	}
	var i int
	for i = range tests {
		var test sizeFromBytesBoundryTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			var size *fayl.Size = fayl.SizeFromBytes(test.bytes)

			if size.Size() != test.expectedSize {
				t.Fatalf("expected size %d, got %d", test.expectedSize, size.Size())
			} else if size.Unit() != test.expectedUnit {
				t.Fatalf("expected unit %v, got %v", test.expectedUnit, size.Unit())
			}
		})
	}
}
