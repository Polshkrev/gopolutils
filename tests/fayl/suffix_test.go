package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/fayl"
)

func TestSuffixCountHasNotChanged(test *testing.T) {
	const expected uint8 = 34
	if fayl.SuffixCount != expected {
		test.Errorf("SuffixCount has changed. Have: %d. Expected: %d\n", fayl.SuffixCount, expected)
	}
}

func TestSuffixFromStringHasMappedSameValue(test *testing.T) {
	var expect fayl.Suffix = fayl.Go
	var value string = "go"
	var item fayl.Suffix
	var except *gopolutils.Exception
	item, except = fayl.SuffixFromString(value)
	if except != nil {
		test.Errorf("%s\n", except.Error())
	} else if item != expect {
		test.Errorf("Value '%s' is not mapped correctly. '%s' has a different mapped value.", value, value)
	}
}

func TestSuffixFromStringDoesNotHaveMappedDifferentValue(test *testing.T) {
	var expect fayl.Suffix = fayl.C
	var value string = "go"
	var item fayl.Suffix
	var except *gopolutils.Exception
	item, except = fayl.SuffixFromString(value)
	if except != nil {
		test.Errorf("%s\n", except.Error())
	} else if item == expect {
		test.Errorf("Value '%s' is not mapped correctly. '%s' has the same mapped value.", value, value)
	}
}

func TestStringFromSuffixHasMappedSameValue(test *testing.T) {
	var expect string = "go"
	var value fayl.Suffix = fayl.Go
	var item string
	var except *gopolutils.Exception
	item, except = fayl.StringFromSuffix(value)
	if except != nil {
		test.Errorf("%s\n", except.Error())
	} else if item != expect {
		test.Errorf("Value '%s' is not mapped correctly. '%s' has a different mapped value.", value, value)
	}
}

func TestStringFromDoesNotHaveMappedDifferentValue(test *testing.T) {
	var expect string = "go"
	var value fayl.Suffix = fayl.C
	var item fayl.Suffix
	var except *gopolutils.Exception
	item, except = fayl.SuffixFromString(value)
	if except != nil {
		test.Errorf("%s\n", except.Error())
	} else if item == expect {
		test.Errorf("Value '%s' is not mapped correctly. '%s' has the same mapped value.", value, value)
	}
}
