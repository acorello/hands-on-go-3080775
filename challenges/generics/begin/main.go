// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// non-generic print functions

type printable interface {
	~string | constraints.Float | constraints.Integer | ~bool
}

func printAny[T printable](s T) { fmt.Println(s) }

// Part 2 sum function refactoring

type numeric interface {
	~int | ~float64
}

// sum sums a slice of any type
func sumAny[T numeric](numbers ...T) (result T) {
	for _, n := range numbers {
		result += n
	}
	return
}

func main() {
	// call non-generic print functions
	printAny("Hello")
	printAny(42)
	printAny(true)

	// call sum function
	fmt.Println("result", sumAny(1, 2, 3))
}
