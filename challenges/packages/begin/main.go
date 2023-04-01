// challenges/packages/begin/proverbs.go
package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

// import the proverbs package

// getRandomProverb returns a random proverb from the proverbs package
func getRandomProverb() string {
	return proverbs.Random().Saying
}

func main() {
	// print the result of calling your getRandomProverb function
	fmt.Println(getRandomProverb())
}
