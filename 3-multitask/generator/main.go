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
	counted := make(chan pair)

	go func() {
		for {
			word := next()
			if word != "" {
				counted <- pair{word, countDigits(word)}
			} else {
				counted <- pair{}
				break
			}
		}
	}()

	stats := make(counter)
	for {
		p := <-counted
		if p.word == "" {
			break
		}
		stats[p.word] = p.count
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
