package mylib_test

import (
	"fmt"
	"testing"

	"CODE/mylib"
)

func TestPush(t *testing.T) {
	for _, test := range pushTests {
		mylib.Push(&test.data_a, &test.data_b)
		if fmt.Sprint(test.data_a) != test.expected[0] || fmt.Sprint(test.data_b) != test.expected[1] {
			t.Errorf("GOT \n%v  %v EXPECTED\n%s  %s", test.data_a, test.data_b, test.expected[0], test.expected[1])
		}
	}
}

func TestSendTo(t *testing.T) {
	for _, test := range sendToTests {
		mylib.SendTo(&test.data, test.i, test.j)
		if fmt.Sprint(test.data) != test.expected {
			t.Errorf("Expected\n'%s' \ngot\n'%v'", test.expected, fmt.Sprint(test.data))
		}
	}
}

type testPush struct {
	data_a, data_b []int
	expected       [2]string
}

var pushTests = []testPush{
	{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, [2]string{"[6 1 2 3 4 5]", "[7 8 9 10]"}},
	{[]int{5, 4, 3, 2, 1}, []int{10, 9, 8, 7, 6}, [2]string{"[10 5 4 3 2 1]", "[9 8 7 6]"}},
}

type testSendTo struct {
	data     []int
	i, j     int
	expected string
}

var sendToTests = []testSendTo{
	{[]int{1, 2, 3, 4, 5}, 0, 4, "[2 3 4 5 1]"},
	{[]int{5, 4, 3, 2, 1}, 0, 4, "[4 3 2 1 5]"},
	{[]int{2, 3, 1, 4, 5}, 0, 1, "[3 2 1 4 5]"},
	{[]int{2, 3, 1, 4, 5}, 4, 0, "[5 2 3 1 4]"},
}
