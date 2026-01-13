package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils/fayl"
)

func TestSuffixCountHasNotChanged(test *testing.T) {
	const expected uint8 = 33
	if fayl.SuffixCount != expected {
		test.Errorf("SuffixCount has changed. Have: %d. Expected: %d\n", fayl.SuffixCount, expected)
	}
}
