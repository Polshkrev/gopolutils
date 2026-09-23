package tests

import (
	"sync"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
)

type stringTest struct {
	name     string
	input    string
	expected fayl.Suffix
}

func TestSuffixFromString(t *testing.T) {
	tests := []stringTest{
		{"Go", "go", fayl.Go},
		{"C", "c", fayl.C},
		{"Header h", "h", fayl.Header},
		{"Header hpp", "hpp", fayl.Header},
		{"Yaml", "yaml", fayl.Yaml},
		{"Yaml Alias", "yml", fayl.Yaml},
		{"Zip", "zip", fayl.Zip},
		{"Zip Alias rar", "rar", fayl.Zip},
		{"Zip Alias 7z", "7z", fayl.Zip},
		{"None", "", fayl.None},
	}
	var i int
	for i = range tests {
		var test stringTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			var result fayl.Suffix
			var exception *gopolutils.Exception
			result, exception = fayl.SuffixFromString(test.input)
			if exception != nil {
				t.Fatalf("unexpected exception: %v", exception)
			} else if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestSuffixFromStringUnknown(t *testing.T) {
	var result fayl.Suffix
	var exception *gopolutils.Exception
	result, exception = fayl.SuffixFromString("does_not_exist")

	if exception == nil {
		t.Fatal("expected exception")
	} else if result != fayl.None {
		t.Fatalf("expected None, got %v", result)
	}
}

type fromSuffixTest struct {
	name     string
	input    fayl.Suffix
	expected string
}

func TestStringFromSuffix(t *testing.T) {
	var tests []fromSuffixTest = []fromSuffixTest{
		{"Go", fayl.Go, "go"},
		{"C", fayl.C, "c"},
		{"Header", fayl.Header, "h"},
		{"Yaml", fayl.Yaml, "yaml"},
		{"Zip", fayl.Zip, "zip"},
		{"None", fayl.None, ""},
	}
	var i int
	for i = range tests {
		var test fromSuffixTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			var result string
			var exception *gopolutils.Exception
			result, exception = fayl.StringFromSuffix(test.input)
			if exception != nil {
				t.Fatalf("unexpected exception: %v", exception)
			} else if result != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, result)
			}
		})
	}
}

func TestStringFromSuffixUnknown(t *testing.T) {
	const invalid fayl.Suffix = fayl.Suffix(100)
	var result string
	var exception *gopolutils.Exception
	result, exception = fayl.StringFromSuffix(invalid)

	if exception == nil {
		t.Fatal("expected exception")
	} else if result != "" {
		t.Fatalf("expected empty string, got %q", result)
	}
}

func TestSuffixRoundTrip(t *testing.T) {
	var tests []fayl.Suffix = []fayl.Suffix{
		fayl.A,
		fayl.Asm,
		fayl.Bat,
		fayl.C,
		fayl.Class,
		fayl.Cpp,
		fayl.Cs,
		fayl.Css,
		fayl.Csv,
		fayl.Db,
		fayl.Dll,
		fayl.Exe,
		fayl.Go,
		fayl.Gz,
		fayl.Header,
		fayl.Html,
		fayl.Jar,
		fayl.Java,
		fayl.Js,
		fayl.Json,
		fayl.Lib,
		fayl.Log,
		fayl.Md,
		fayl.None,
		fayl.O,
		fayl.Py,
		fayl.Rs,
		fayl.Sh,
		fayl.So,
		fayl.Sql,
		fayl.Tar,
		fayl.Tex,
		fayl.Toml,
		fayl.Txt,
		fayl.Xml,
		fayl.Yaml,
		fayl.Zip,
	}
	var i int
	for i = range tests {
		var original fayl.Suffix = tests[i]
		var suffixString string
		var exception *gopolutils.Exception
		suffixString, exception = fayl.StringFromSuffix(original)
		if exception != nil {
			t.Fatalf("unexpected exception: %v", exception)
		}
		var result fayl.Suffix
		result, exception = fayl.SuffixFromString(suffixString)
		if exception != nil {
			t.Fatalf("unexpected exception: %v", exception)
		} else if result != original {
			t.Fatalf("round-trip failed: %v -> %q -> %v", original, suffixString, result)
		}
	}
}

func TestSuffixFromStringConcurrent(t *testing.T) {
	var inputs []string = []string{
		"go",
		"c",
		"h",
		"hpp",
		"yaml",
		"yml",
		"zip",
		"rar",
		"7z",
		"",
	}

	var wg sync.WaitGroup
	for range 100 {
		var i int
		for i = range inputs {
			wg.Add(1)

			var input string = inputs[i]
			go func(inputString string) {
				defer wg.Done()
				var exception *gopolutils.Exception
				_, exception = fayl.SuffixFromString(inputString)
				if exception != nil {
					t.Errorf("unexpected exception for %q: %v", inputString, exception)
				}
			}(input)
		}
	}

	wg.Wait()
}

func TestStringFromSuffixConcurrent(t *testing.T) {
	var values []fayl.Suffix = []fayl.Suffix{
		fayl.Go,
		fayl.C,
		fayl.Header,
		fayl.Yaml,
		fayl.Zip,
		fayl.None,
	}

	var wg sync.WaitGroup

	for range 100 {
		var i int
		for i = range values {
			wg.Add(1)

			var value fayl.Suffix = values[i]
			go func(suffix fayl.Suffix) {
				defer wg.Done()
				var exception *gopolutils.Exception
				_, exception = fayl.StringFromSuffix(suffix)
				if exception != nil {
					t.Errorf("unexpected exception for %v: %v", suffix, exception)
				}
			}(value)
		}
	}

	wg.Wait()
}
