package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"CODE/mylib"
)

func main() {
	data_a := mylib.Parse(os.Args[1:]...)
	data_b := make([]int, 0, len(data_a))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		// fmt.Println(command, data_a, data_b)
		mylib.Execute(&data_a, &data_b, command)
	}

	// fmt.Println(data_a, data_b)
	if !mylib.IsSorted(data_a, true) || len(data_b) != 0 {
		fmt.Println("KO")
		return
	}
	fmt.Println("OK")
}

func init() {
	if strings.HasSuffix(os.Args[0], ".test") {
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./checker <name>", len(os.Args))
		os.Exit(0)
	}
}
