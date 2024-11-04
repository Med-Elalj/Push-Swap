package main

import (
	"fmt"
	"os"

	"CODE/mylib"
)

func main() {
	// Read input from command line arguments
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage:./push_swap <numbers>", os.Args[1], os.Args)
		return // No input, do nothing
	}

	stackA := mylib.Parse(os.Args[1:]...)

	stackB := make([]int, 0, len(stackA))
	instructions := []string{}

	// Main sorting logic
	sortStacks(&stackA, &stackB, &instructions)

	// Output instructions
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}

func sortStacks(stackA *[]int, stackB *[]int, instructions *[]string) {
	for !mylib.IsSorted(*stackA, true) || len(*stackB) > 0 {
		if !mylib.IsSorted(*stackA, true) {
			sort(stackA, stackB, instructions, true)
		}
		if !mylib.IsSorted(*stackB, false) {
			sort(stackA, stackB, instructions, false)
		}
		if len(*stackB) > 0 && mylib.IsSorted(*stackB, false) && mylib.IsSorted(*stackA, true) {
			*instructions = append(*instructions, "pa")
			mylib.Execute(stackA, stackB, "pa")
		}
	}
}

func sort(stackA, stackB *[]int, instructions *[]string, order bool) {
	var sa, ra bool
	var t, nt string

	if order {
		if len(*stackA) < 2 {
			return
		}
		sa = (*stackA)[0] > (*stackA)[1]
		ra = (*stackA)[0] > (*stackA)[len(*stackA)-1]
		t = "a"
		nt = "b"
	} else {
		if len(*stackB) < 2 {
			return
		}
		sa = (*stackB)[0] < (*stackB)[1]
		ra = (*stackB)[0] < (*stackB)[len(*stackB)-1]
		t = "b"
		nt = "a"
	}

	switch {
	case sa && !ra:
		*instructions = append(*instructions, "s"+t)
		mylib.Execute(stackA, stackB, "s"+t)
	case !sa && !ra:
		*instructions = append(*instructions, "p"+nt)
		mylib.Execute(stackA, stackB, "p"+nt)
	case sa && ra:
		*instructions = append(*instructions, "r"+t)
		mylib.Execute(stackA, stackB, "r"+t)
	case !sa && ra:
		mylib.Execute(stackA, stackB, "rr"+t)
		*instructions = append(*instructions, "rr"+t)
	}
}
