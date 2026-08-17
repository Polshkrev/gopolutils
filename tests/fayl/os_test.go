package tests

import (
	"runtime"
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
)

func TestCurrentOperatingSystem(t *testing.T) {
	var expected fayl.OperatingSystem = fayl.OperatingSystem(runtime.GOOS)

	var actual fayl.OperatingSystem = fayl.CurrentOperatingSystem()

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

type operatingSystemConstantTest struct {
	name     string
	system   fayl.OperatingSystem
	expected string
}

func TestOperatingSystemConstants(t *testing.T) {
	var tests []operatingSystemConstantTest = []operatingSystemConstantTest{
		{"Windows", fayl.Windows, "windows"},
		{"Mac", fayl.Mac, "darwin"},
		{"Linux", fayl.Linux, "linux"},
		{"Android", fayl.Android, "android"},
		{"FreeBSD", fayl.Freebsd, "freebsd"},
		{"IOS", fayl.Ios, "ios"},
		{"NetBSD", fayl.Netbsd, "netbsd"},
		{"OpenBSD", fayl.Openbsd, "openbsd"},
	}
	var i int
	for i = range tests {
		var test operatingSystemConstantTest = tests[i]
		t.Run(test.name, func(t *testing.T) {
			if string(test.system) != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, string(test.system))
			}
		})
	}
}

func TestCurrentOperatingSystemKnownValue(t *testing.T) {
	var known map[fayl.OperatingSystem]gopolutils.None = map[fayl.OperatingSystem]gopolutils.None{
		fayl.Windows: {},
		fayl.Mac:     {},
		fayl.Linux:   {},
		fayl.Android: {},
		fayl.Freebsd: {},
		fayl.Ios:     {},
		fayl.Netbsd:  {},
		fayl.Openbsd: {},
	}

	var current fayl.OperatingSystem = fayl.CurrentOperatingSystem()
	var ok bool
	_, ok = known[current]
	if !ok {
		t.Logf("current operating system %q is not one of the predefined enum values", current)
	}
}

type operatingSystemStringTest struct {
	system fayl.OperatingSystem
	value  string
}

func TestOperatingSystemString(t *testing.T) {
	var tests []operatingSystemStringTest = []operatingSystemStringTest{
		{fayl.Windows, "windows"},
		{fayl.Mac, "darwin"},
		{fayl.Linux, "linux"},
		{fayl.Android, "android"},
		{fayl.Freebsd, "freebsd"},
		{fayl.Ios, "ios"},
		{fayl.Netbsd, "netbsd"},
		{fayl.Openbsd, "openbsd"},
	}
	var i int
	for i = range tests {
		var test operatingSystemStringTest = tests[i]
		var got string = string(test.system)
		if got != test.value {
			t.Fatalf("expected %q, got %q", test.value, got)
		}
	}
}
