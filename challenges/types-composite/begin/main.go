// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author
}

// define a library type to track a list of books
type library []book

// define addBook to add a book to the library
func (me *library) addBook(b book) {
	*me = append(*me, b)
}

// define a lookupByAuthor function to find books by an author's name
func (me library) lookupByAuthor(name string) (res []book) {
	for _, b := range me {
		if b.name == name {
			res = append(res, b)
		}
	}
	return
}

func main() {
	// create a new library
	bookshop := library([]book{
		{
			title: "This is life",
			author: author{
				name: "John Fante",
			},
		}, {
			title: "Moby Dick",
			author: author{
				name: "John Fante",
			},
		},
	})

	// add 1 book to the library by a different author
	bookshop = append(bookshop, book{
		title: "Divina Commedia",
		author: author{
			name: "Dante Alighieri",
		},
	})

	// dump the library with spew
	spew.Dump(bookshop)

	// lookup books by known author in the library
	fantesBooks := bookshop.lookupByAuthor("John Fante")

	// print out the first book's title and its author's name
	fmt.Println("First Fante's book", fantesBooks[0])

}
