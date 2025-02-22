package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
)

// This functionality is not implemented yet. This test should fail.
func TestGetLevelSuccess(test *testing.T) {
	var level gopolutils.LoggingLevel = gopolutils.DEBUG
	var logger *gopolutils.Logger = gopolutils.NewLogger("test", level)
	var result gopolutils.LoggingLevel = logger.GetLevel()
	if result != level {
		test.Errorf("Expected: %d got: %d", level, result)
	}
}

// This functionality is not implemented yet. This test should fail.
func TestSetLevelSuccess(test *testing.T) {
	var logger *gopolutils.Logger = gopolutils.NewLogger("test", gopolutils.DEBUG)
	logger.SetLevel(gopolutils.INFO)
	var result gopolutils.LoggingLevel = logger.GetLevel()
	if result != gopolutils.INFO {
		test.Errorf("Expected: %d got: %d", gopolutils.INFO, result)
	}
}

func TestSetLevelFromSameLevel(test *testing.T) {
	var logger *gopolutils.Logger = gopolutils.NewLogger("test", gopolutils.DEBUG)
	logger.SetLevel(gopolutils.DEBUG)
	var result gopolutils.LoggingLevel = logger.GetLevel()
	if result != gopolutils.DEBUG {
		test.Errorf("Expected: %d got: %d", gopolutils.DEBUG, result)
	}
}
