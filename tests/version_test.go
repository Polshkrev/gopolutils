package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
)

func TestNewVersionZerosOut(test *testing.T) {
	var version *gopolutils.Version = gopolutils.NewVersion()
	if version.Major() != 0 {
		test.Errorf("New version major is not set correctly: %d. Expected: %d\n", version.Major(), 0)
	} else if version.Minor() != 0 {
		test.Errorf("New version minor is not set correctly: %d. Expected: %d\n", version.Minor(), 0)
	} else if version.Patch() != 0 {
		test.Errorf("New version patch is not set correctly: %d. Expected: %d\n", version.Patch(), 0)
	}
}

func TestVersionConvertSetsCorrectly(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 1, 1)
	if version.Major() != 1 {
		test.Errorf("New version major is not set correctly: %d. Expected: %d\n", version.Major(), 0)
	} else if version.Minor() != 1 {
		test.Errorf("New version minor is not set correctly: %d. Expected: %d\n", version.Minor(), 0)
	} else if version.Patch() != 1 {
		test.Errorf("New version patch is not set correctly: %d. Expected: %d\n", version.Patch(), 1)
	}
}

func TestVersionCompareMajor(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(2, 0, 0)
	var compare *gopolutils.Version = gopolutils.VersionConvert(1, 0, 0)
	var result bool = version.CompareMajor(compare.Major())
	if !result {
		test.Errorf("Version major is less than the compare value. Got: %t. Expected: %t", result, false)
	}
}
