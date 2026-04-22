package common

import (
	"testing"
)

func TestNewError(t *testing.T) {
	err := NewError("something went wrong")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if err.Error() == "" {
		t.Fatal("expected non-empty error message")
	}
}

func TestNewErrorf(t *testing.T) {
	err := NewErrorf("error code %d: %s", 42, "bad input")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	want := "error code 42: bad input"
	if err.Error() != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}

func TestNewErrorMultipleArgs(t *testing.T) {
	err := NewError("part1", "part2", 123)
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	// Sprintln separates with spaces and adds trailing newline
	if len(err.Error()) == 0 {
		t.Error("expected non-empty error message")
	}
}

func TestRecover_NoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("unexpected panic")
		}
	}()
	result := Recover("test")
	if result != nil {
		t.Errorf("expected nil result when no panic, got %v", result)
	}
}
