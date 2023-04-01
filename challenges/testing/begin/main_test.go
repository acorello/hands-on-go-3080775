// challenges/testing/begin/main_test.go
package main

import "testing"

type testCase struct {
	name  string
	input string
	want  int
}

func TestLetterCounter(t *testing.T) {
	t.Parallel()
	runTests(t, letterCounter(0), []testCase{
		{"empty", "", 0},
		{"just_1_letters", "a", 1},
		{"just_2_letters", "ab", 2},
		{"two_letters,spaces,and_punct", "a b!", 2},
		{"spaces_and_punct_but_no_letters", "?1 .", 0},
	})
}

func TestSymbolsCounter(t *testing.T) {
	t.Parallel()
	runTests(t, symbolCounter(0), []testCase{
		{"empty", "", 0},
		{"just 1 symbol", ".", 1},
		{"just 2 symbols", "%!", 2},
		{"two symbols, spaces, and letters", "a? ^", 2},
		{"spaces and letters but no symbols", "a1 ", 0},
	})
}

func TestNumbersCounter(t *testing.T) {
	t.Parallel()
	runTests(t, numberCounter(0), []testCase{
		{"empty", "", 0},
		{"just 1 number", "1", 1},
		{"just 2 numbers", "12", 2},
		{"two numbers, spaces, and letters", "1 ^2", 2},
		{"spaces and letters but no numbers", "a! ", 0},
	})
}

func runTests(t *testing.T, counter counter, tests []testCase) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			count := counter.count(test.input)
			if count != test.want {
				t.Errorf("expected %d, got %d", test.want, count)
			}
		})
	}
}
