package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
	// "fmt"
)

func TestPushSwap(t *testing.T) {
	stacks := []string{
		"1 2 3 4 ", "2 3 4 1", "4 3 2 1", "4 67 3 87 23", "4 9 17 3 97",
		"2 1 3 6 5 8", "0 1 2 3 4 5",
	}
	for _, stack := range stacks {
		out, err := exec.Command("go", "run", ".", stack).CombinedOutput()
		if err != nil {
			fmt.Println(stack, err)
			t.Fail()
			continue
		}
		// fmt.Printf("%12s count %d len %d\n", stack, strings.Count(string(out), "\n"), len(mylib.Parse(stack)))
		cmd := exec.Command("go", "run", "../checker/main.go", stack)
		var buf bytes.Buffer
		buf.Write(out)

		// Set the standard input of the command to the buffer
		cmd.Stdin = &buf

		// Run the command
		out, err = cmd.CombinedOutput()
		// Check the output
		if string(out) != "OK\n" || err != nil {
			fmt.Println(len(stack), string(out), err)
			t.Fail()
		}
	}
}
