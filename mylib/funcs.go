package mylib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// IsSorted returns true if the given array is sorted .
// returns false if it's not sorted or empty.
func IsSorted(data []int, ascending bool) bool {
	if len(data) == 0 {
		return false
	}
	var last int = data[0]
	for _, v := range data {
		if last > v && ascending {
			return false
		}
		last = v
	}
	return true
}

// Parse returns a array of integers from the given space separated string.
func Parse(parse ...string) []int {
	input := strings.Join(parse, " ")
	in := strings.Fields(input)
	data := make([]int, 0, len(in))
	for _, v := range in {
		if v == "" {
			continue
		}
		d, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error")
			os.Exit(0)
		}
		data = append(data, d)
	}
	switch {
	case len(data) < 2:
		fmt.Println("Error")
		os.Exit(0)
	case hasDuplicate(data):
		fmt.Println("Error")
		os.Exit(0)
	}

	return data
}

// hasDuplicate takes an array of integers and check if one of them is a duplicated.
func hasDuplicate(data []int) bool {
	for idx1, int1 := range data {
		for idx2, int2 := range data {
			if int1 == int2 && idx1 != idx2 {
				return true
			}
		}
	}
	return false
}

func Execute(data_a, data_b *[]int, command string) {
	switch command {
	case "sa":
		SendTo(data_a, 0, 1)
	case "sb":
		SendTo(data_b, 0, 1)
	case "ss":
		SendTo(data_a, 0, 1)
		SendTo(data_b, 0, 1)
	case "pa":
		Push(data_a, data_b)
	case "pb":
		Push(data_b, data_a)
	case "ra":
		SendTo(data_a, 0, len(*data_a)-1)
	case "rb":
		SendTo(data_b, 0, len(*data_b)-1)
	case "rr":
		SendTo(data_a, 0, len(*data_a)-1)
		SendTo(data_b, 0, len(*data_b)-1)
	case "rra":
		SendTo(data_a, len(*data_a)-1, 0)
	case "rrb":
		SendTo(data_b, len(*data_b)-1, 0)
	case "rrr":
		SendTo(data_a, len(*data_a)-1, 0)
		SendTo(data_b, len(*data_b)-1, 0)
	default:
		fmt.Println("Error command not found", command)
		os.Exit(2)
	}
}
