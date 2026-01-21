package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
)

// This functionality is not implemented yet. This test should fail.
func TestGetLevelSuccess(test *testing.T) {
	var level gopolutils.LoggingLevel = gopolutils.Debug
	var logger *gopolutils.Logger = gopolutils.NewLogger("test", level)
	var result gopolutils.LoggingLevel = logger.GetLevel()
	if result != level {
		test.Errorf("Expected: %d got: %d", level, result)
	}
}

// This functionality is not implemented yet. This test should fail.
func TestSetLevelSuccess(test *testing.T) {
	var logger *gopolutils.Logger = gopolutils.NewLogger("test", gopolutils.Debug)
	logger.SetLevel(gopolutils.Info)
	var result gopolutils.LoggingLevel = logger.GetLevel()
	if result != gopolutils.Info {
		test.Errorf("Expected: %d got: %d", gopolutils.Info, result)
	}
}

func TestSetLevelFromSameLevel(test *testing.T) {
	var logger *gopolutils.Logger = gopolutils.NewLogger("test", gopolutils.Debug)
	logger.SetLevel(gopolutils.Debug)
	var result gopolutils.LoggingLevel = logger.GetLevel()
	if result != gopolutils.Debug {
		test.Errorf("Expected: %d got: %d", gopolutils.Debug, result)
	}
}
