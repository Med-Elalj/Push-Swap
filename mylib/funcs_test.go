package mylib_test

import (
	"testing"

	"CODE/mylib"
)

func TestIsSorted(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	if !mylib.IsSorted(data) {
		t.Error("Expected true, got false")
	}
	data = []int{1, 3, 2, 4, 5}
	if mylib.IsSorted(data) {
		t.Error("Expected false, got true")
	}
}

func TestParse(t *testing.T) {
	data := mylib.Parse("1 2 3 4 5")
	if len(data) != 5 {
		t.Error("Expected 5, got", len(data))
	}
	data = mylib.Parse("1 3 2 4 5")
	if len(data) != 5 {
		t.Error("Expected 5, got", len(data))
	}
}
