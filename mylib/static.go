package mylib

import (
	"fmt"
	"os"
)

// Push transfers the last element of the `from` slice to the end of the `to` slice.
// It modifies both slices in place. If `from` is empty, it prints an error and exits.
func Push(to, from *[]int) {
	if len(*from) == 0 {
		fmt.Println("Error empty from slice")
		os.Exit(1)
	}
	*to = append([]int{(*from)[0]}, (*to)...)
	*from = (*from)[1:]
}

// SendTo transfers the element at index i to index j keeping the order of the rest.
// If indices are out of bounds or data is nil, it prints an error and exits.
func SendTo(data *[]int, i, j int) {
	if len(*data) <= i || len(*data) <= j || data == nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	v := (*data)[i]
	*data = append((*data)[:i], (*data)[i+1:]...)
	x := append([]int{v}, (*data)[j:]...)
	*data = append((*data)[:j], x...)
}
