package tests

import (
	"testing"

	"github.com/Polshkrev/gopolutils/table"
)

func TestFieldString(t *testing.T) {
	var field table.Field = "username"

	if field.String() != "username" {
		t.Fatalf("expected \"username\", got %q", field.String())
	}
}

func TestIdString(t *testing.T) {
	if table.Id.String() != "id" {
		t.Fatalf("expected \"id\", got %q", table.Id.String())
	}
}

func TestGetFieldsEmpty(t *testing.T) {
	var result string = table.GetFields()

	if result != "" {
		t.Fatalf("expected empty string, got %q", result)
	}
}

func TestGetFieldsSingle(t *testing.T) {
	var result string = table.GetFields(table.Id)

	if result != "id" {
		t.Fatalf("expected \"id\", got %q", result)
	}
}

func TestGetFieldsMultiple(t *testing.T) {
	var result string = table.GetFields(
		table.Id,
		table.Field("name"),
		table.Field("email"),
	)

	if result != "id, name, email" {
		t.Fatalf("unexpected result: %q", result)
	}
}

func TestGetFieldsDuplicate(t *testing.T) {
	var result string = table.GetFields(
		table.Id,
		table.Id,
	)

	if result != "id, id" {
		t.Fatalf("unexpected result: %q", result)
	}
}

func TestGetFieldsCustomFields(t *testing.T) {
	var result string = table.GetFields(
		table.Field("first_name"),
		table.Field("last_name"),
		table.Field("created_at"),
	)

	if result != "first_name, last_name, created_at" {
		t.Fatalf("unexpected result: %q", result)
	}
}

func TestGetFieldsEmptyField(t *testing.T) {
	var result string = table.GetFields(
		table.Field(""),
	)

	if result != "" {
		t.Fatalf("expected empty string, got %q", result)
	}
}

func TestGetFieldsMixedEmptyFields(t *testing.T) {
	var result string = table.GetFields(
		table.Id,
		table.Field(""),
		table.Field("name"),
	)

	if result != "id, , name" {
		t.Fatalf("unexpected result: %q", result)
	}
}
