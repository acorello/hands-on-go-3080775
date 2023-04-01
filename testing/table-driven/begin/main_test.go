// testing/table-driven/begin/main_test.go
package main

import "testing"

func TestSum(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		{"one", []int{1}, 1},
		{"two", []int{1, -2}, 1},
		{"three", []int{1, 2, 3}, 6},
	}

	// range over the tests and run them as subtests
	for i, test := range tests {
		result := sum(test.input...)
		if result != test.want {
			// print an int with at least 2 digits
			t.Errorf("Nº%0.2d: %s\twanted %d got %d", i, test.name, test.want, result)
		}
	}
}
