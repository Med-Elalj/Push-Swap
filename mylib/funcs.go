package mylib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// IsSorted returns true if the given array is sorted .
// returns false if it's not sorted or empty.
func IsSorted(data []int) bool {
	if len(data) == 0 {
		return false
	}
	var last int = data[0]
	for _, v := range data {
		if last > v {
			return false
		}
		last = v
	}
	return true
}

// Parse returns a array of integers from the given space separated string.
func Parse(input string) []int {
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
