package main

import (
	"fmt"
	"strings"
	"unicode"
)

type nextFunc func() string

type counter map[string]int

type pair struct {
	word  string
	count int
}

func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	counted := make(chan pair)
	go countWords(pending, counted)

	return fillStats(counted)
}

func submitWords(next nextFunc, pending chan<- string) {
	for {
		word := next()
		if word == "" {
			close(pending)
			return
		}
		pending <- word
	}
}

func countWords(pending <-chan string, counted chan<- pair) {
	for {
		word, ok := <-pending
		if !ok {
			close(counted)
			return
		}
		counted <- pair{
			word:  word,
			count: countDigits(word),
		}
	}
}

func fillStats(counted <-chan pair) counter {
	stats := make(counter)
	for pair := range counted {
		stats[pair.word] = pair.count
	}
	return stats
}

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
