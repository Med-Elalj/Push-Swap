package main_test

import (
	"bytes"
	"os/exec"
	"testing"
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
