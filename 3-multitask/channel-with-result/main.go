package main

import (
	"fmt"
	"strings"
	"unicode"
)

type counter map[string]int

func countDigitsInWords(phrase string) counter {
	stats := make(counter)
	words := strings.Fields(phrase)
	counted := make(chan int)

	go func(words []string) {
		for _, word := range words {
			count := countDigits(word)
			counted <- count
		}
	}(words)

	for _, word := range words {
		count := <-counted
		stats[word] = count
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

func main() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWords(phrase)
	printStats(stats)
}
