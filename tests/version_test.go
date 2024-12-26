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
		test.Errorf("Version major is less than the compare value. Got: %t. Expected: %t", result, true)
	}
}

func TestVersionCompareMajorFailure(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 0, 0)
	var compare *gopolutils.Version = gopolutils.VersionConvert(2, 0, 0)
	var result bool = version.CompareMajor(compare.Major())
	if result {
		test.Errorf("Version major is less than the compare value. Got: %t. Expected: %t", result, false)
	}
}

func TestVersionCompareMinor(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 2, 0)
	var compare *gopolutils.Version = gopolutils.VersionConvert(0, 1, 0)
	var result bool = version.CompareMinor(compare.Minor())
	if !result {
		test.Errorf("Version minor is less than the compare value. Got: %t. Expected: %t", result, true)
	}
}

func TestVersionCompareMinorFailure(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 1, 0)
	var compare *gopolutils.Version = gopolutils.VersionConvert(0, 2, 0)
	var result bool = version.CompareMinor(compare.Minor())
	if result {
		test.Errorf("Version major is less than the compare value. Got: %t. Expected: %t", result, false)
	}
}

func TestVersionComparePatch(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(2, 0, 0)
	var compare *gopolutils.Version = gopolutils.VersionConvert(1, 0, 0)
	var result bool = version.ComparePatch(compare.Patch())
	if !result {
		test.Errorf("Version patch is less than the compare value. Got: %t. Expected: %t", result, true)
	}
}

func TestVersionComparePatchFailure(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 0, 1)
	var compare *gopolutils.Version = gopolutils.VersionConvert(0, 0, 2)
	var result bool = version.ComparePatch(compare.Patch())
	if result {
		test.Errorf("Version major is less than the compare value. Got: %t. Expected: %t", result, false)
	}
}

func TestVersionCompare(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(2, 0, 0)
	var compare *gopolutils.Version = gopolutils.VersionConvert(1, 0, 0)
	var result bool = version.Compare(*compare)
	if !result {
		test.Errorf("Version is less than the compare value. Got: %t. Expected: %t", result, true)
	}
}

func TestVersionCompareFailure(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 0, 0)
	var compare *gopolutils.Version = gopolutils.VersionConvert(2, 0, 0)
	var result bool = version.Compare(*compare)
	if result {
		test.Errorf("Version major is less than the compare value. Got: %t. Expected: %t", result, false)
	}
}

func TestVersionIsZeroTrue(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 0, 0)
	if !version.IsZero() {
		test.Errorf("Version: '%+v' is not evaluated to be zero.", *version)
	}
}

func TestVersionIsZeroFalse(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 1, 1)
	if version.IsZero() {
		test.Errorf("Version: '%+v' is evaluated to be zero.", *version)
	}
}

func TestVersionIsZeroMajorFalse(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 0, 0)
	if version.IsZero() {
		test.Errorf("Version: '%+v' is evaluated to be zero.", *version)
	}
}

func TestVersionIsZeroMinorFalse(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 1, 0)
	if version.IsZero() {
		test.Errorf("Version: '%+v' is evaluated to be zero.", *version)
	}
}

func TestVersionIsZeroPatchFalse(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 0, 1)
	if version.IsZero() {
		test.Errorf("Version: '%+v' is evaluated to be zero.", *version)
	}
}

func TestVersionIsPublic(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 0, 0)
	var result bool = version.IsPublic()
	if !result {
		test.Errorf("Version is less than the compare value. Got: %t. Expected: %t", result, true)
	}
}

func TestVersionIsPublicFailure(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 1, 0)
	var result bool = version.IsPublic()
	if result {
		test.Errorf("Version is less than the compare value. Got: %t. Expected: %t", result, false)
	}
}

func TestVersionPublish(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(0, 0, 0)
	var result *gopolutils.Exception = version.Publish()
	if result != nil {
		test.Errorf("Version publish returned a non nil exception. Got: %+v. Expected: %+v", result, nil)
	}
}

func TestVersionPublishAlredyPublic(test *testing.T) {
	var version *gopolutils.Version = gopolutils.VersionConvert(1, 0, 0)
	var result *gopolutils.Exception = version.Publish()
	if result == nil {
		test.Errorf("Version publish returned a nil exception. Got: %+v. Expected: %+v", result, nil)
	}
}

func TestVersionRelease(test *testing.T) {
	var version *gopolutils.Version = gopolutils.NewVersion()
	version.Release()
	if version.Major() != 1 {
		test.Errorf("Version release did not increment correctly. Got: %d, Expected: %d.", version.Major(), 1)
	}
}

func TestVersionUpdate(test *testing.T) {
	var version *gopolutils.Version = gopolutils.NewVersion()
	version.Update()
	if version.Minor() != 1 {
		test.Errorf("Version update did not increment correctly. Got: %d, Expected: %d.", version.Minor(), 1)
	}
}

func TestVersionFix(test *testing.T) {
	var version *gopolutils.Version = gopolutils.NewVersion()
	version.Fix()
	if version.Patch() != 1 {
		test.Errorf("Version release did not increment correctly. Got: %d, Expected: %d.", version.Patch(), 1)
	}
}
