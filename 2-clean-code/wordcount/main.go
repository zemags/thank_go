package main

import (
	"fmt"
	"os"
	"strings"
)

func wordCount(input string) int {
	words := strings.Fields(input)
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
