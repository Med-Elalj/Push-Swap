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
	for len(*stackA) > 3 && !mylib.IsSorted(*stackA, true) {
		smallestIdx := indxOfSmallest(stackA)
		if smallestIdx == 0 {
			mylib.Execute(stackA, stackB, "pb")
			*instructions = append(*instructions, "pb")
		} else if smallestIdx > len(*stackA)/2 {
			for i := smallestIdx; i < len(*stackA); i++ {
				mylib.Execute(stackA, stackB, "rra")
				*instructions = append(*instructions, "rra")
			}
			mylib.Execute(stackA, stackB, "pb")
			*instructions = append(*instructions, "pb")
		} else {
			for i := smallestIdx; i > 0; i-- {
				mylib.Execute(stackA, stackB, "ra")
				*instructions = append(*instructions, "ra")
			}
			mylib.Execute(stackA, stackB, "pb")
			*instructions = append(*instructions, "pb")
		}
	}
	sortThree(stackA, instructions)
	for len(*stackB) > 0 {
		mylib.Execute(stackA, stackB, "pa")
		*instructions = append(*instructions, "pa")
	}
}

func sortThree(stack *[]int, instructions *[]string) {
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
		*instructions = append(*instructions, "sa")
		mylib.Execute(stack, nil, "ra")
		*instructions = append(*instructions, "ra")
	case first && !second && !last:
		// 2 3 1
		mylib.Execute(stack, nil, "rra")
		*instructions = append(*instructions, "rra")
	case !first && !second && !last:
		// 3 2 1
		mylib.Execute(stack, nil, "sa")
		*instructions = append(*instructions, "sa")
		mylib.Execute(stack, nil, "rra")
		*instructions = append(*instructions, "rra")
	case !first && !second && last:
		// 2 3 1
		mylib.Execute(stack, nil, "ra")
		*instructions = append(*instructions, "ra")
	case !first && second && !last:
		// 3 1 2
		mylib.Execute(stack, nil, "ra")
		*instructions = append(*instructions, "ra")
	case !first && second && last:
		// 2 1 3
		mylib.Execute(stack, nil, "sa")
		*instructions = append(*instructions, "sa")
	}
	// fmt.Println(stack)
}

// 123
// 213
// 312
// 321
// 231
// 132
func indxOfSmallest(stack *[]int) int {
	smallest := (*stack)[0]
	smallestIdx := 0
	for i, val := range *stack {
		if val < smallest {
			smallest = val
			smallestIdx = i
		}
	}
	return smallestIdx
}

// arg="2 3 4 1" ;go run . ${arg} | go run ../checker/main.go ${arg}
// arg="2 1 3" ;go run . ${arg} | go run ../checker/main.go ${arg}
// arg="3 1 2" ;go run . ${arg} | go run ../checker/main.go ${arg}
// arg="3 2 1" ;go run . ${arg} | go run ../checker/main.go ${arg}
// arg="2 3 1" ;go run . ${arg} | go run ../checker/main.go ${arg}
// arg="1 3 2" ;go run . ${arg} | go run ../checker/main.go ${arg}
