package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
)

func TestNewVersion(t *testing.T) {
	var version *gopolutils.Version = gopolutils.NewVersion()

	if version == nil {
		t.Fatal("Expected non-nil version")
	} else if version.Name() != "" {
		t.Errorf("Expected empty name, got %q", version.Name())
	} else if version.Description() != "" {
		t.Errorf("Expected empty description, got %q", version.Description())
	} else if version.Major() != 0 {
		t.Errorf("Expected major 0, got %d", version.Major())
	} else if version.Minor() != 0 {
		t.Errorf("Expected minor 0, got %d", version.Minor())
	} else if version.Patch() != 0 {
		t.Errorf("Expected patch 0, got %d", version.Patch())
	}
}

func TestVersionConvert(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 2, 3)
	if version.Major() != 1 {
		t.Errorf("Expected major 1, got %d", version.Major())
	} else if version.Minor() != 2 {
		t.Errorf("Expected minor 2, got %d", version.Minor())
	} else if version.Patch() != 3 {
		t.Errorf("Expected patch 3, got %d", version.Patch())
	}
}

func TestNewNamedVersion(t *testing.T) {
	var version *gopolutils.Version = gopolutils.NewNamedVersion("Test")

	if version.Name() != "Test" {
		t.Errorf("Expected name Test, got %q", version.Name())
	} else if version.Description() != "" {
		t.Errorf("Expected empty description")
	}
}

func TestNewStringVersion(t *testing.T) {
	var version *gopolutils.Version = gopolutils.NewStringVersion("Test", "Description")

	if version.Name() != "Test" {
		t.Errorf("Expected name Test")
	} else if version.Description() != "Description" {
		t.Errorf("Expected description Description")
	}
}

func TestNewFullVersion(t *testing.T) {
	var version *gopolutils.Version = gopolutils.NewFullVersion("Test", "Description", 1, 2, 3)

	if version.Name() != "Test" {
		t.Errorf("Unexpected name")
	} else if version.Description() != "Description" {
		t.Errorf("Unexpected description")
	} else if version.Major() != 1 || version.Minor() != 2 || version.Patch() != 3 {
		t.Errorf("Unexpected version numbers")
	}
}

func TestSetters(t *testing.T) {
	var version *gopolutils.Version = gopolutils.NewVersion()

	version.SetName("Example")
	version.SetDescription("Testing")
	version.SetMajor(5)
	version.SetMinor(6)
	version.SetPatch(7)

	if version.Name() != "Example" {
		t.Errorf("Unexpected name")
	} else if version.Description() != "Testing" {
		t.Errorf("Unexpected description")
	} else if version.Major() != 5 {
		t.Errorf("Unexpected major")
	} else if version.Minor() != 6 {
		t.Errorf("Unexpected minor")
	} else if version.Patch() != 7 {
		t.Errorf("Unexpected patch")
	}
}

func TestCompareMajor(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(3, 0, 0)

	if !version.CompareMajor(3) {
		t.Error("Expected CompareMajor(3) to be true")
	} else if !version.CompareMajor(2) {
		t.Error("Expected CompareMajor(2) to be true")
	} else if version.CompareMajor(4) {
		t.Error("Expected CompareMajor(4) to be false")
	}
}

func TestCompareMinor(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 4, 0)

	if !version.CompareMinor(4) {
		t.Error("Expected CompareMinor(4)")
	} else if !version.CompareMinor(2) {
		t.Error("Expected CompareMinor(2)")
	} else if version.CompareMinor(5) {
		t.Error("Expected CompareMinor(5) to be false")
	}
}

func TestComparePatch(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 0, 8)

	if !version.ComparePatch(8) {
		t.Error("Expected ComparePatch(8)")
	} else if !version.ComparePatch(2) {
		t.Error("Expected ComparePatch(2)")
	} else if version.ComparePatch(9) {
		t.Error("Expected ComparePatch(9) to be false")
	}
}

func TestCompare(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(2, 5, 7)

	if !version.Compare(*gopolutils.VersionConvert(2, 5, 7)) {
		t.Error("Expected equal versions")
	} else if !version.Compare(*gopolutils.VersionConvert(1, 4, 6)) {
		t.Error("Expected comparison to succeed")
	} else if version.Compare(*gopolutils.VersionConvert(3, 0, 0)) {
		t.Error("Expected comparison to fail")
	} else if version.Compare(*gopolutils.VersionConvert(2, 6, 0)) {
		t.Error("Expected comparison to fail")
	} else if version.Compare(*gopolutils.VersionConvert(2, 5, 8)) {
		t.Error("Expected comparison to fail")
	}
}

func TestIsZero(t *testing.T) {
	if !gopolutils.NewVersion().IsZero() {
		t.Error("Expected zero version")
	} else if gopolutils.VersionConvert(0, 0, 1).IsZero() {
		t.Error("Expected non-zero version")
	}
}

func TestIsPublic(t *testing.T) {
	if gopolutils.VersionConvert(0, 5, 0).IsPublic() {
		t.Error("Expected private version")
	} else if !gopolutils.VersionConvert(1, 0, 0).IsPublic() {
		t.Error("Expected public version")
	}
}

func TestPublish(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 5, 7)

	var except *gopolutils.Exception = version.Publish()
	if except != nil {
		t.Fatalf("unexpected error: %v", except)
	} else if version.Major() != 1 {
		t.Errorf("Expected major 1")
	} else if version.Minor() != 0 {
		t.Errorf("Expected minor 0")
	} else if version.Patch() != 0 {
		t.Errorf("Expected patch 0")
	}
}

func TestPublishAlreadyPublic(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 2, 3)

	var except *gopolutils.Exception = version.Publish()

	if except == nil {
		t.Fatal("Expected error")
	} else if version.Major() != 1 || version.Minor() != 2 || version.Patch() != 3 {
		t.Error("version should not have changed")
	}
}

func TestRelease(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 5, 8)

	version.Release()
	if version.Major() != 2 {
		t.Errorf("Expected major 2")
	} else if version.Minor() != 0 {
		t.Errorf("Expected minor 0")
	} else if version.Patch() != 0 {
		t.Errorf("Expected patch 0")
	}
}

func TestUpdate(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(2, 5, 8)

	version.Update()

	if version.Major() != 2 {
		t.Errorf("major should not change")
	} else if version.Minor() != 6 {
		t.Errorf("Expected minor 6")
	} else if version.Patch() != 0 {
		t.Errorf("Expected patch 0")
	}
}

func TestFix(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(2, 3, 4)

	version.Fix()

	if version.Major() != 2 {
		t.Errorf("major should not change")
	} else if version.Minor() != 3 {
		t.Errorf("minor should not change")
	} else if version.Patch() != 5 {
		t.Errorf("Expected patch 5")
	}
}

type stringTest struct {
	name     string
	version  *gopolutils.Version
	expected string
}

func TestString(t *testing.T) {
	var tests []stringTest = []stringTest{
		{
			name:     "numbers only",
			version:  gopolutils.VersionConvert(1, 2, 3),
			expected: "1.2.3",
		},
		{
			name:     "description only",
			version:  gopolutils.NewFullVersion("", "beta", 1, 2, 3),
			expected: "1.2.3 - beta",
		},
		{
			name:     "name only",
			version:  gopolutils.NewFullVersion("App", "", 1, 2, 3),
			expected: "App: 1.2.3",
		},
		{
			name:     "name and description",
			version:  gopolutils.NewFullVersion("App", "beta", 1, 2, 3),
			expected: "App: 1.2.3 - beta",
		},
	}

	var i int
	for i = range tests {
		var test stringTest = tests[i]
		var result string = test.version.String()
		if result != test.expected {
			t.Errorf("%s: Expected %q, got %q", test.name, test.expected, result)
		}
	}
}

func TestNumberString(t *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(7, 8, 9)

	if version.NumberString() != "7.8.9" {
		t.Errorf("Expected 7.8.9, got %q", version.NumberString())
	}
}
