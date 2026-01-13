package tests

import "testing"

func TestSuffixCountHasNotChanged(test *testing.T) {
	const expected fayl.Suffix = fayl.ZIP
	if fayl.SuffixCount != expected {
		test.Errorf("SuffixCount has changed. Have: %d. Expected: %d\n", fayl.SuffixCount, expected)
	}
}
