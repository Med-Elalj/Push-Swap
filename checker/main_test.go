package main_test

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

	checker "CODE/checker"
)

func TestMain(t *testing.T) {
	for _, test := range maintests {
		cmd := exec.Command("go", "run", "main.go", test[1])

		// Create a buffer to simulate the echo input
		var buf bytes.Buffer
		buf.WriteString(test[0])

		// Set the standard input of the command to the buffer
		cmd.Stdin = &buf

		// Run the command
		out, err := cmd.CombinedOutput()
		// Check the output
		if string(out) != test[2] {
			t.Error(string(out), err)
		}
	}
}

var maintests = [][3]string{
	{"rra\npb\nsa\nrra\npa\n", "3 2 1 0", "OK\n"},
	{"rra\npb\nsa\n", "3 2 one 0", "Error\n"},
}

func TestExecute(t *testing.T) {
	for _, test := range executetests {
		checker.Execute(&test.data_a, &test.data_b, test.command)
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
