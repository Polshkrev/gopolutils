package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
	"github.com/Polshkrev/gopolutils/fayl"
)

func TestWrite(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")
	var expected []byte = []byte("hello world")
	var exception *gopolutils.Exception = fayl.Write(fayl.PathFrom(path), expected)
	if exception != nil {
		t.Fatal(exception)
	}
	var actual []byte
	var readError error
	actual, readError = os.ReadFile(path)
	if readError != nil {
		t.Fatal(readError)
	} else if !bytes.Equal(actual, expected) {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestWriteInvalidPath(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "missing", "file.txt")
	var exception *gopolutils.Exception = fayl.Write(fayl.PathFrom(path), []byte("test"))
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestWriteObjectJSON(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.json")

	var object *testObject = &testObject{
		Name: "John",
		Age:  42,
	}
	var exception *gopolutils.Exception = fayl.WriteObject(fayl.PathFrom(path), object)
	if exception != nil {
		t.Fatal(exception)
	}
	var actual []byte
	var readError error
	actual, readError = os.ReadFile(path)
	if readError != nil {
		t.Fatal(readError)
	} else if len(actual) == 0 {
		t.Fatal("expected data")
	}
}

func TestWriteObjectYAML(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.yaml")

	var object *testObject = &testObject{
		Name: "John",
		Age:  42,
	}
	var exception *gopolutils.Exception = fayl.WriteObject(fayl.PathFrom(path), object)
	if exception != nil {
		t.Fatal(exception)
	}
	var actual []byte
	var readError error
	actual, readError = os.ReadFile(path)
	if readError != nil {
		t.Fatal(readError)
	} else if len(actual) == 0 {
		t.Fatal("expected data")
	}
}

func TestWriteObjectTOML(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.toml")

	var object *testObject = &testObject{
		Name: "John",
		Age:  42,
	}
	var exception *gopolutils.Exception = fayl.WriteObject(fayl.PathFrom(path), object)
	if exception != nil {
		t.Fatal(exception)
	}
	var actual []byte
	var readError error
	actual, readError = os.ReadFile(path)
	if readError != nil {
		t.Fatal(readError)
	} else if len(actual) == 0 {
		t.Fatal("expected data")
	}
}

func TestWriteObjectCSV(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.csv")

	var object *testObject = &testObject{
		Name: "John",
		Age:  42,
	}
	var exception *gopolutils.Exception = fayl.WriteObject(fayl.PathFrom(path), object)
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestWriteObjectUnknownSuffix(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.invalid")

	var object *testObject = &testObject{
		Name: "John",
		Age:  42,
	}
	var exception *gopolutils.Exception = fayl.WriteObject(fayl.PathFrom(path), object)
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestWriteListJSON(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.json")
	var view collections.Collection[testObject] = collections.NewArray[testObject]()
	view.Append(testObject{"Alice", 20})
	view.Append(testObject{"Bob", 30})
	var exception *gopolutils.Exception = fayl.WriteList(fayl.PathFrom(path), view)
	if exception != nil {
		t.Fatal(exception)
	}
	var data []byte
	var writeError error
	data, writeError = os.ReadFile(path)
	if writeError != nil {
		t.Fatal(writeError)
	} else if len(data) == 0 {
		t.Fatal("expected data")
	}
}

func TestWriteListYAML(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.yaml")

	var view collections.Collection[testObject] = collections.NewArray[testObject]()
	view.Append(testObject{"Alice", 20})
	view.Append(testObject{"Bob", 30})
	var exception *gopolutils.Exception = fayl.WriteList(fayl.PathFrom(path), view)
	if exception != nil {
		t.Fatal(exception)
	}
}

func TestWriteListTOML(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.toml")

	var view collections.Collection[testObject] = collections.NewArray[testObject]()
	view.Append(testObject{"Alice", 20})
	view.Append(testObject{"Bob", 30})
	var exception *gopolutils.Exception = fayl.WriteList(fayl.PathFrom(path), view)
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestWriteListCSV(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.csv")

	var view collections.Collection[testObject] = collections.NewArray[testObject]()
	view.Append(testObject{"Alice", 20})
	view.Append(testObject{"Bob", 30})
	var exception *gopolutils.Exception = fayl.WriteList(fayl.PathFrom(path), view)
	if exception != nil {
		t.Fatal(exception)
	}
}

func TestWriteListUnknownSuffix(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.invalid")

	var view collections.Collection[testObject] = collections.NewArray[testObject]()
	view.Append(testObject{"Alice", 20})
	var exception *gopolutils.Exception = fayl.WriteList(fayl.PathFrom(path), view)
	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestWriteEmptyBytes(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "empty.txt")
	var exception *gopolutils.Exception = fayl.Write(fayl.PathFrom(path), []byte{})
	if exception != nil {
		t.Fatal(exception)
	}
	var info os.FileInfo
	var statError error
	info, statError = os.Stat(path)
	if statError != nil {
		t.Fatal(statError)
	} else if info.Size() != 0 {
		t.Fatalf("expected empty file, got %d bytes", info.Size())
	}
}

func TestWriteNilBytes(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "nil.txt")
	var exception *gopolutils.Exception = fayl.Write(fayl.PathFrom(path), nil)
	if exception != nil {
		t.Fatal(exception)
	}
	var info os.FileInfo
	var statError error
	info, statError = os.Stat(path)
	if statError != nil {
		t.Fatal(statError)
	} else if info.Size() != 0 {
		t.Fatalf("expected empty file, got %d bytes", info.Size())
	}
}

func TestWriteObjectNilPointer(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.json")

	var object *testObject = nil

	// Some serializers marshal nil as "null", others return an error.
	// Accept either behavior, but ensure we don't panic.
	_ = fayl.WriteObject(fayl.PathFrom(path), object)
}

func TestWriteListEmpty(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "empty.json")

	var view collections.Collection[testObject] = collections.NewArray[testObject]()
	var exception *gopolutils.Exception = fayl.WriteList(fayl.PathFrom(path), view)
	if exception != nil {
		t.Fatal(exception)
	}
	var data []byte
	var readError error
	data, readError = os.ReadFile(path)
	if readError != nil {
		t.Fatal(readError)
	} else if len(data) == 0 {
		t.Fatal("expected serialized empty collection")
	}
}

func TestWriteOverwrite(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "file.txt")
	var exception *gopolutils.Exception = fayl.Write(fayl.PathFrom(path), []byte("first"))
	if exception != nil {
		t.Fatal(exception)
	} else if exception := fayl.Write(fayl.PathFrom(path), []byte("second")); exception != nil {
		t.Fatal(exception)
	}
	var data []byte
	var readError error
	data, readError = os.ReadFile(path)
	if readError != nil {
		t.Fatal(readError)
	} else if string(data) != "second" {
		t.Fatalf("expected %q, got %q", "second", string(data))
	}
}

func TestWriteReadRoundTripJSON(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.json")

	var expected *testObject = &testObject{
		Name: "John",
		Age:  42,
	}
	var exception *gopolutils.Exception = fayl.WriteObject(fayl.PathFrom(path), expected)
	if exception != nil {
		t.Fatal(exception)
	}
	var actual *testObject
	actual, exception = fayl.ReadObject[testObject](fayl.PathFrom(path))
	if exception != nil {
		t.Fatal(exception)
	} else if actual.Name != expected.Name {
		t.Fatalf("expected %q, got %q", expected.Name, actual.Name)
	} else if actual.Age != expected.Age {
		t.Fatalf("expected %d, got %d", expected.Age, actual.Age)
	}
}

func TestWriteReadRoundTripListJSON(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.json")

	var expected collections.Collection[testObject] = collections.NewArray[testObject]()
	expected.Append(testObject{"Alice", 20})
	expected.Append(testObject{"Bob", 30})
	var exception *gopolutils.Exception = fayl.WriteList(fayl.PathFrom(path), expected)
	if exception != nil {
		t.Fatal(exception)
	}
	var actual collections.View[testObject]
	actual, exception = fayl.ReadList[testObject](fayl.PathFrom(path))
	if exception != nil {
		t.Fatal(exception)
	} else if actual.Size() != expected.Size() {
		t.Fatalf("expected %d objects, got %d", expected.Size(), actual.Size())
	}
	var i gopolutils.Size
	for i = 0; i < actual.Size(); i++ {
		var left testObject = expected.Collect()[i]
		var right testObject = actual.Collect()[i]

		if left.Name != right.Name {
			t.Fatalf("index %d: expected %q, got %q", i, left.Name, right.Name)
		} else if left.Age != right.Age {
			t.Fatalf("index %d: expected %d, got %d", i, left.Age, right.Age)
		}
	}
}
