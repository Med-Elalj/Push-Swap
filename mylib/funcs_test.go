package mylib_test

import (
	"fmt"
	"testing"

	"CODE/mylib"
)

func TestIsSorted(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	if !mylib.IsSorted(data, true) {
		t.Error("Expected true, got false")
	}
	data = []int{1, 3, 2, 4, 5}
	if mylib.IsSorted(data, true) {
		t.Error("Expected false, got true")
	}
	data = []int{5, 4, 3, 2, 1}
	if !mylib.IsSorted(data, false) {
		t.Error("Expected true, got false")
	}
	data = []int{5, 3, 4, 2, 1}
	if !mylib.IsSorted(data, false) {
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

func TestExecute(t *testing.T) {
	for _, test := range executetests {
		mylib.Execute(&test.data_a, &test.data_b, test.command)
		if fmt.Sprint(test.data_a) != test.expect_a || fmt.Sprint(test.data_b) != test.expect_b {
			t.Errorf("Expected %s\n%v %v got\n%v,%v", test.command, test.expect_a, test.expect_b, test.data_a, test.data_b)
		}
	}
}

type executetest struct {
	data_a, data_b              []int
	command, expect_a, expect_b string
}

var executetests = []executetest{
	{[]int{3, 2, 1}, []int{5, 9, 7}, "pa", "[5 3 2 1]", "[9 7]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "pb", "[2 1]", "[3 5 9 7]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "sa", "[2 3 1]", "[5 9 7]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "sb", "[3 2 1]", "[9 5 7]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "ss", "[2 3 1]", "[9 5 7]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "rra", "[1 3 2]", "[5 9 7]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "rrb", "[3 2 1]", "[7 5 9]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "rrr", "[1 3 2]", "[7 5 9]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "ra", "[2 1 3]", "[5 9 7]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "rb", "[3 2 1]", "[9 7 5]"},
	{[]int{3, 2, 1}, []int{5, 9, 7}, "rr", "[2 1 3]", "[9 7 5]"},
	{[]int{3, 2, 1, 0}, []int{}, "rra", "[0 3 2 1]", "[]"},
	{[]int{0, 3, 2, 1}, []int{}, "pb", "[3 2 1]", "[0]"},
	{[]int{3, 2, 1}, []int{0}, "sa", "[2 3 1]", "[0]"},
	{[]int{2, 3, 1}, []int{0}, "rra", "[1 2 3]", "[0]"},
	{[]int{1, 2, 3}, []int{0}, "pa", "[0 1 2 3]", "[]"},
}
