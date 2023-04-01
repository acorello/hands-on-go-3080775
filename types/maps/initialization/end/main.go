package main

import "fmt"

type author struct {
	name string
	surn string
}

func main() {
	authors := map[string]struct{ name, surn string }{
		"tm": {name: "Toni", surn: "Morrison"},
		"ma": {name: "Marcus", surn: "Aurelius"},
	}

	// alternative syntax to declare and initialize the author map
	// authors := map[string]author{
	// 	"tm": author{name: "Toni Morrison"},
	// 	"ma": author{name: "Marcus Aurelius"},
	// }

	// print the map with fmt.Printf
	fmt.Printf("%#v\n", authors)

	// read a value from the map with a known key
	// fmt.Println(authors["tm"])
}
