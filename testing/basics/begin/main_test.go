// testing/basics/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

// write a test for sum
func TestSum(t *testing.T) {
	expected := 3
	actual := sum(1, 2)
	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

// write a TestMain for setup and teardown
func TestMain(t *testing.M) {
	fmt.Println("SETUP")
	t.Run()
	fmt.Println("TEARDOWN")
}
