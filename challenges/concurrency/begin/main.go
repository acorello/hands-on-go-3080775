// challenges/concurrency/begin/main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"unicode"
)

type counter interface {
	name() string
	count(input string)
	value() uint
}

type letterCounter uint

func (l letterCounter) value() uint {
	return uint(l)
}

func (l letterCounter) name() string {
	return "letters"
}

func (l *letterCounter) count(input string) {
	for _, char := range input {
		if unicode.IsLetter(char) {
			(*l)++
		}
	}
}

type numberCounter uint

func (n numberCounter) value() uint {
	return uint(n)
}

func (n numberCounter) name() string {
	return "numbers"
}

func (n *numberCounter) count(input string) {
	for _, char := range input {
		if unicode.IsNumber(char) {
			(*n)++
		}
	}
}

type symbolCounter uint

func (s symbolCounter) value() uint {
	return uint(s)
}

func (s symbolCounter) name() string {
	return "symbols"
}

func (s *symbolCounter) count(input string) {
	for _, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			(*s)++
		}
	}
}

type callCounter uint

func (s callCounter) value() uint {
	return uint(s)
}

func (s callCounter) name() string {
	return "words"
}

func (s *callCounter) count(input string) {
	(*s)++
}

func main() {
	fileBody, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	defer fileBody.Close()
	analists := analysts()
	doAnalisys(fileBody, analists...)
	for _, analist := range analists {
		fmt.Printf("%s: %d\n", analist.name(), analist.value())
	}
}

func analysts() []counter {
	var (
		l = letterCounter(0)
		n = numberCounter(0)
		s = symbolCounter(0)
		c = callCounter(0)
	)
	var analists = []counter{&l, &n, &s, &c}
	return analists
}

func doAnalisys(dataReader io.Reader, analists ...counter) {
	scanner := bufio.NewScanner(dataReader)
	scanner.Split(bufio.ScanWords)
	// one-channel and go-routine per analist
	channels := make([]chan string, 0)
	wg := sync.WaitGroup{}
	for _, analyst := range analists {
		wg.Add(1)
		ch := make(chan string)
		channels = append(channels, ch)
		go func(analyst counter) {
			for word := range ch {
				analyst.count(word)
			}
			wg.Done()
		}(analyst)
	}

	for scanner.Scan() {
		word := scanner.Text()
		for _, channel := range channels {
			channel <- word
		}
	}
	for _, channel := range channels {
		close(channel)
	}
	wg.Wait()
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
}
