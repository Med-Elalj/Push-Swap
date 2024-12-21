package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
	// "fmt"
)

func TestPushSwap(t *testing.T) {
	stacks := []string{
		"1 2 3 4 ", "2 3 4 1", "4 3 2 1",
		"2 1 3 6 5 8", "0 1 2 3 4 5", "4 9 17 3 97", "4 67 3 87 23",
	}
	for _, stack := range stacks {
		out, err := exec.Command("go", "run", ".", stack).CombinedOutput()
		if err != nil {
			fmt.Println(stack, err)
			t.Fail()
			continue
		}
		fmt.Println(stack, "count", strings.Count(string(out), "\n"))
		cmd := exec.Command("go", "run", "../checker/main.go", stack)
		var buf bytes.Buffer
		buf.Write(out)

		// Set the standard input of the command to the buffer
		cmd.Stdin = &buf

		// Run the command
		out, err = cmd.CombinedOutput()
		// Check the output
		if string(out) != "OK\n" || err != nil {
			fmt.Println(len(stack), err)
			t.Fail()
		}
	}
}
