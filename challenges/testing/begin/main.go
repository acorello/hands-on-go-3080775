// challenges/testing/begin/main.go
package main

import "unicode"

type counter interface {
	name() string
	count(input string) int
}

type letterCounter uint

func (n letterCounter) name() string {
	return "letters"
}
func (l letterCounter) count(input string) int {
	result := 0
	for _, char := range input {
		if unicode.IsLetter(char) {
			result++
		}
	}
	return result
}

type numberCounter uint

func (n numberCounter) name() string {
	return "numbers"
}

func (n numberCounter) count(input string) int {
	result := 0
	for _, char := range input {
		if unicode.IsNumber(char) {
			result++
		}
	}
	return result
}

type symbolCounter uint

func (s symbolCounter) name() string {
	return "symbols"
}

func (s symbolCounter) count(input string) int {
	result := 0
	for _, char := range input {
		if  unicode.IsSymbol(char) || unicode.IsPunct(char)  {
			result++
		}
	}
	return result
}
