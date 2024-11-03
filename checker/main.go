package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"CODE/mylib"
)

func main() {
	data_a := mylib.Parse(os.Args[1])
	data_b := make([]int, 0, len(data_a))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		// fmt.Println(command, data_a, data_b)
		Execute(&data_a, &data_b, command)
		// fmt.Println(mylib.IsSorted(data_a), command)
	}

	// fmt.Println(data_a, data_b)
	if !mylib.IsSorted(data_a) && len(data_b) == 0 {
		fmt.Println("KO", data_a, data_b)
		return
	}
	fmt.Println("OK")
}

func init() {
	if strings.HasSuffix(os.Args[0], ".test") {
		return
	}
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./checker <name>", len(os.Args))
		os.Exit(0)
	}
}

func Execute(data_a, data_b *[]int, command string) {
	switch command {
	case "sa":
		mylib.SendTo(data_a, 0, 1)
	case "sb":
		mylib.SendTo(data_b, 0, 1)
	case "ss":
		mylib.SendTo(data_a, 0, 1)
		mylib.SendTo(data_b, 0, 1)
	case "pa":
		mylib.Push(data_a, data_b)
	case "pb":
		mylib.Push(data_b, data_a)
	case "ra":
		mylib.SendTo(data_a, 0, len(*data_a)-1)
	case "rb":
		mylib.SendTo(data_b, 0, len(*data_b)-1)
	case "rr":
		mylib.SendTo(data_a, 0, len(*data_a)-1)
		mylib.SendTo(data_b, 0, len(*data_b)-1)
	case "rra":
		mylib.SendTo(data_a, len(*data_a)-1, 0)
	case "rrb":
		mylib.SendTo(data_b, len(*data_b)-1, 0)
	case "rrr":
		mylib.SendTo(data_a, len(*data_a)-1, 0)
		mylib.SendTo(data_b, len(*data_b)-1, 0)
	default:
		fmt.Println("Error command not found", command)
		os.Exit(2)
	}
}
