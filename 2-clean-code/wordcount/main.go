package main

import (
	"fmt"
	"os"
	"strings"
)

func wordCount(input string) int {
	// Split the input string into words using spaces as separators
	words := strings.Fields(input)
	// Return the count of words
	return len(words)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wordcount <input_string>")
		return
	}

	input := os.Args[1]
	wordCount := wordCount(input)
	fmt.Printf("%d\n", wordCount)
}
