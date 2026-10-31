package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils"
	"github.com/Polshkrev/gopolutils/collections"
)

func TestNewNode(test *testing.T) {
	var node *collections.Node[string] = collections.NewNode("value")
	if node == nil {
		test.Fatal("NewNode() returned nil")
	} else if node.Data() != "value" {
		test.Errorf("NewNode().Data() = %q, want %q", node.Data(), "value")
	} else if node.HasNext() {
		test.Error("NewNode() unexpectedly has a next node")
	} else if node.HasPrevious() {
		test.Error("NewNode() unexpectedly has a previous node")
	}
}

func TestNodeSetData(test *testing.T) {
	var node *collections.Node[int] = collections.NewNode(10)
	node.SetData(20)
	if node.Data() != 20 {
		test.Errorf("Data() = %d, want %d", node.Data(), 20)
	}
}

func TestNodeSetNext(test *testing.T) {
	var node *collections.Node[string] = collections.NewNode("first")
	var next *collections.Node[string] = collections.NewNode("second")

	node.SetNext(next)

	if !node.HasNext() {
		test.Fatal("HasNext() = false, want true")
	}

	var actual collections.Node[string]
	var exception *gopolutils.Exception

	actual, exception = node.Next()
	if exception != nil {
		test.Fatalf("Next() returned an unexpected exception: %v", exception)
	} else if actual.Data() != "second" {
		test.Errorf("Next().Data() = %q, want %q", actual.Data(), "second")
	}
}

func TestNodeSetPrevious(test *testing.T) {
	var node *collections.Node[string] = collections.NewNode("second")
	var previous *collections.Node[string] = collections.NewNode("first")

	node.SetPrevious(previous)

	if !node.HasPrevious() {
		test.Fatal("HasPrevious() = false, want true")
	}

	var actual collections.Node[string]
	var exception *gopolutils.Exception

	actual, exception = node.Previous()
	if exception != nil {
		test.Fatalf("Previous() returned an unexpected exception: %v", exception)
	} else if actual.Data() != "first" {
		test.Errorf("Previous().Data() = %q, want %q", actual.Data(), "first")
	}
}

func TestNodeNextWithoutLink(test *testing.T) {
	var node *collections.Node[string] = collections.NewNode("value")

	var actual collections.Node[string]
	var exception *gopolutils.Exception

	actual, exception = node.Next()

	if exception == nil {
		test.Fatal("Next() returned no exception, want an error")
	} else if !exception.Is(gopolutils.ValueError) {
		test.Errorf("Next() exception type = %s, want %v", exception.Name(), gopolutils.ValueError)
	} else if actual.Data() != "" {
		test.Errorf("Next() returned data %q, want zero value", actual.Data())
	}
}

func TestNodePreviousWithoutLink(test *testing.T) {
	var node *collections.Node[string] = collections.NewNode("value")

	var actual collections.Node[string]
	var exception *gopolutils.Exception

	actual, exception = node.Previous()

	if exception == nil {
		test.Fatal("Previous() returned no exception, want an error")
	} else if !exception.Is(gopolutils.ValueError) {
		test.Errorf("Previous() exception type = %s, want %v", exception.Name(), gopolutils.ValueError)
	} else if actual.Data() != "" {
		test.Errorf("Previous() returned data %q, want zero value", actual.Data())
	}
}

func TestNodeHasLinks(test *testing.T) {
	var first *collections.Node[int] = collections.NewNode(1)
	var second *collections.Node[int] = collections.NewNode(2)

	if first.HasNext() {
		test.Error("HasNext() = true before linking, want false")
	} else if first.HasPrevious() {
		test.Error("HasPrevious() = true before linking, want false")
	}

	first.SetNext(second)
	second.SetPrevious(first)

	if !first.HasNext() {
		test.Error("HasNext() = false after linking, want true")
	} else if !second.HasPrevious() {
		test.Error("HasPrevious() = false after linking, want true")
	}
}

func TestNodeString(test *testing.T) {
	var node *collections.Node[string] = collections.NewNode("value")
	var actual string = node.String()
	var expected string = "value"

	if actual != expected {
		test.Errorf("String() = %q, want %q", actual, expected)
	}
}

func TestNodeStringWithInteger(test *testing.T) {
	var node *collections.Node[int] = collections.NewNode(42)
	var actual string = node.String()
	var expected string = "42"

	if actual != expected {
		test.Errorf("String() = %q, want %q", actual, expected)
	}
}
