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

type testObject struct {
	Name string `json:"name" yaml:"name" toml:"name" csv:"name"`
	Age  int    `json:"age" yaml:"age" toml:"age" csv:"age"`
}

func TestReadSuccess(t *testing.T) {
	var expected []byte = []byte("hello world")

	var path string = filepath.Join(t.TempDir(), "file.txt")
	var writeError error = os.WriteFile(path, expected, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var actual []byte
	var exception *gopolutils.Exception
	actual, exception = fayl.Read(fayl.PathFrom(path))
	if exception != nil {
		t.Fatal(exception)
	} else if !bytes.Equal(actual, expected) {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestReadMissingFileExcept(t *testing.T) {
	var exception *gopolutils.Exception
	_, exception = fayl.Read(fayl.PathFrom(filepath.Join(t.TempDir(), "missing.txt")))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestReadObjectJSON(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.json")

	var data []byte = []byte(`{"name":"John","age":42}`)
	var writeError error = os.WriteFile(path, data, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var object *testObject
	var exception *gopolutils.Exception
	object, exception = fayl.ReadObject[testObject](fayl.PathFrom(path))
	if exception != nil {
		t.Fatal(exception)
	} else if object.Name != "John" {
		t.Fatal("incorrect name")
	} else if object.Age != 42 {
		t.Fatal("incorrect age")
	}
}

func TestReadObjectInvalidJSON(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.json")
	var writeError error = os.WriteFile(path, []byte("{"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var exception *gopolutils.Exception
	_, exception = fayl.ReadObject[testObject](fayl.PathFrom(path))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestReadListJSON(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.json")

	var data []byte = []byte(`
[
	{"name":"Alice","age":20},
	{"name":"Bob","age":30}
]
`)
	var writeError error = os.WriteFile(path, data, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var list collections.View[testObject]
	var exception *gopolutils.Exception
	list, exception = fayl.ReadList[testObject](fayl.PathFrom(path))
	if exception != nil {
		t.Fatal(exception)
	} else if list.Size() != 2 {
		t.Fatalf("expected 2 objects, got %d", list.Size())
	}

	var first testObject = list.Collect()[0]

	if first.Name != "Alice" {
		t.Fatal("unexpected first object")
	}
}

func TestReadListInvalidJSON(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.json")
	var writeError error = os.WriteFile(path, []byte("["), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var exception *gopolutils.Exception
	_, exception = fayl.ReadList[testObject](fayl.PathFrom(path))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestReadObjectMissingFile(t *testing.T) {
	var exception *gopolutils.Exception
	_, exception = fayl.ReadObject[testObject](fayl.PathFrom(filepath.Join(t.TempDir(), "missing.json")))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestReadListMissingFile(t *testing.T) {
	var exception *gopolutils.Exception
	_, exception = fayl.ReadList[testObject](fayl.PathFrom(filepath.Join(t.TempDir(), "missing.json")))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestReadObjectUnknownSuffix(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.invalid")
	var writeError error = os.WriteFile(path, []byte("{}"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var exception *gopolutils.Exception
	_, exception = fayl.ReadObject[testObject](fayl.PathFrom(path))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestReadListUnknownSuffix(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.invalid")
	var writeError error = os.WriteFile(path, []byte("[]"), 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var exception *gopolutils.Exception
	_, exception = fayl.ReadList[testObject](fayl.PathFrom(path))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestReadObjectYAML(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.yaml")

	data := []byte(`
name: John
age: 42
`)
	var writeError error = os.WriteFile(path, data, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var object *testObject
	var exception *gopolutils.Exception
	object, exception = fayl.ReadObject[testObject](fayl.PathFrom(path))
	if exception != nil {
		t.Fatal(exception)
	} else if object.Name != "John" {
		t.Fatalf("expected John, got %q", object.Name)
	} else if object.Age != 42 {
		t.Fatalf("expected 42, got %d", object.Age)
	}
}

func TestReadObjectTOML(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.toml")

	data := []byte(`
name = "John"
age = 42
`)
	var writeError error = os.WriteFile(path, data, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var object *testObject
	var exception *gopolutils.Exception
	object, exception = fayl.ReadObject[testObject](fayl.PathFrom(path))
	if exception != nil {
		t.Fatal(exception)
	} else if object.Name != "John" {
		t.Fatalf("expected John, got %q", object.Name)
	} else if object.Age != 42 {
		t.Fatalf("expected 42, got %d", object.Age)
	}
}

func TestReadObjectCSV(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "object.csv")

	var data []byte = []byte(`name,age
John,42
`)
	var writeError error = os.WriteFile(path, data, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var exception *gopolutils.Exception
	_, exception = fayl.ReadObject[testObject](fayl.PathFrom(path))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestReadListYAML(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.yaml")

	var data []byte = []byte(`
- name: Alice
  age: 20
- name: Bob
  age: 30
`)
	var writeError error = os.WriteFile(path, data, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var list collections.View[testObject]
	var exception *gopolutils.Exception
	list, exception = fayl.ReadList[testObject](fayl.PathFrom(path))
	if exception != nil {
		t.Fatal(exception)
	} else if list.Size() != 2 {
		t.Fatalf("expected 2 objects, got %d", list.Size())
	} else if list.Collect()[0].Name != "Alice" {
		t.Fatal("unexpected first object")
	} else if list.Collect()[1].Name != "Bob" {
		t.Fatal("unexpected second object")
	}
}

func TestReadListTOML(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.toml")

	var data []byte = []byte(`
[[items]]
name = "Alice"
age = 20

[[items]]
name = "Bob"
age = 30
`)
	var writeError error = os.WriteFile(path, data, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var exception *gopolutils.Exception
	_, exception = fayl.ReadList[testObject](fayl.PathFrom(path))

	if exception == nil {
		t.Fatal("expected exception")
	}
}

func TestReadListCSV(t *testing.T) {
	var path string = filepath.Join(t.TempDir(), "list.csv")

	var data []byte = []byte(`name,age
Alice,20
Bob,30
`)
	var writeError error = os.WriteFile(path, data, 0644)
	if writeError != nil {
		t.Fatal(writeError)
	}
	var list collections.View[testObject]
	var exception *gopolutils.Exception
	list, exception = fayl.ReadList[testObject](fayl.PathFrom(path))
	if exception != nil {
		t.Fatal(exception)
	} else if list.Size() != 2 {
		t.Fatalf("expected 2 objects, got %d", list.Size())
	} else if list.Collect()[0].Name != "Alice" {
		t.Fatal("unexpected first object")
	} else if list.Collect()[1].Name != "Bob" {
		t.Fatal("unexpected second object")
	}
}
