package tests

import (
	"bytes"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
)

func TestGzipSuccess(test *testing.T) {
	var content []byte = []byte("Hello World")
	var except *gopolutils.Exception
	_, except = fayl.Gzip(content)
	if except != nil {
		test.Errorf("Gzip has failed: %s\n", except.Error())
	}
}

func TestGunzipSuccess(test *testing.T) {
	var content []byte = []byte("Hello World")
	var compressed []byte
	var except *gopolutils.Exception
	compressed, except = fayl.Gzip(content)
	if except != nil {
		test.Errorf("Gzip has failed: %s\n", except.Error())
	}
	_, except = fayl.Gunzip(compressed)
	if except != nil {
		test.Errorf("Gunzip has failed: %s\n", except.Error())
	}
}

func TestGzipIsSame(test *testing.T) {
	var content []byte = []byte("Hello World")
	var compressed []byte
	var except *gopolutils.Exception
	compressed, except = fayl.Gzip(content)
	if except != nil {
		test.Errorf("Gzip has failed: %s\n", except.Error())
	}
	var uncompressed []byte
	uncompressed, except = fayl.Gunzip(compressed)
	if except != nil {
		test.Errorf("Gunzip has failed: %s\n", except.Error())
	} else if !bytes.Equal(content, uncompressed) {
		test.Errorf("The raw and uncompressed data is not the same: %v - %v\n", content, uncompressed)
	}
}
