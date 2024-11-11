package main

import (
	"fmt"
	"os"

	"CODE/mylib"
)

func main() {
	// Read input from command line arguments
	if len(os.Args) < 2 {
		if mylib.PrintUsage {
			fmt.Fprintln(os.Stderr, "Usage:./push_swap <numbers>")
		}
		return
	}

	stackA := mylib.Parse(os.Args[1:]...)

	stackB := make([]int, 0, len(stackA))
	var instructions string

	// Main sorting logic
	sortStacks(&stackA, &stackB, &instructions)

	// Output instructions
	fmt.Print(instructions)
}

func sortStacks(stackA *[]int, stackB *[]int, instructions *string) {
	for len(*stackA) > 3 && !mylib.IsSorted(*stackA, true) {
		smallestIdx := mylib.MinElementIndex(stackA)
		if smallestIdx == 0 {
			mylib.Execute(stackA, stackB, "pb")
			*instructions += "pb\n"
		} else if smallestIdx > len(*stackA)/2 {
			for i := smallestIdx; i < len(*stackA); i++ {
				mylib.Execute(stackA, stackB, "rra")
				*instructions += "rra\n"
			}
			mylib.Execute(stackA, stackB, "pb")
			*instructions += "pb\n"
		} else {
			for i := smallestIdx; i > 0; i-- {
				mylib.Execute(stackA, stackB, "ra")
				*instructions += "ra\n"
			}
			mylib.Execute(stackA, stackB, "pb")
			*instructions += "pb\n"
		}
	}
	sortThree(stackA, instructions)
	for len(*stackB) > 0 {
		mylib.Execute(stackA, stackB, "pa")
		*instructions += "pa\n"
	}
}

func sortThree(stack *[]int, instructions *string) {
	first := (*stack)[0] < (*stack)[1]
	second := (*stack)[1] < (*stack)[2]
	last := (*stack)[2] > (*stack)[0]
	switch {
	case first && second:
		// 1 2 3
		return
	case first && !second && last:
		// 2 1 3
		mylib.Execute(stack, nil, "sa")
		*instructions += "sa\n"
		mylib.Execute(stack, nil, "ra")
		*instructions += "ra\n"
	case first && !second && !last:
		// 2 3 1
		mylib.Execute(stack, nil, "rra")
		*instructions += "rra\n"
	case !first && !second && !last:
		// 3 2 1
		mylib.Execute(stack, nil, "sa")
		*instructions += "sa\n"
		mylib.Execute(stack, nil, "rra")
		*instructions += "rra\n"
	case !first && !second && last:
		// 2 3 1
		mylib.Execute(stack, nil, "ra")
		*instructions += "ra\n"
	case !first && second && !last:
		// 3 1 2
		mylib.Execute(stack, nil, "ra")
		*instructions += "ra\n"
	case !first && second && last:
		// 2 1 3
		mylib.Execute(stack, nil, "sa")
		*instructions += "sa\n"
	}
}
