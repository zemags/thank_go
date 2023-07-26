package profiler

import (
	"sort"
	"strings"
)

func JoinWords(first, second string) []string {
	words1 := split(first)
	words2 := split(second)
	words := join(words1, words2)
	words = lower(words)
	words = sorted(words)
	words = uniq(words)
	return words
}

func split(str string) []string {
	return strings.Fields(str)
}

func join(words1, words2 []string) []string {
	words := make([]string, len(words1)+len(words2))
	idx := 0
	for _, word := range words1 {
		words[idx] = word
		idx++
	}
	for _, word := range words2 {
		words[idx] = word
		idx++
	}
	return words
}

func lower(words []string) []string {
	for idx, word := range words {
		words[idx] = strings.ToLower(word)
	}
	return words
}

func sorted(words []string) []string {
	sort.Strings(words)
	return words
}

func uniq(words []string) []string {
	uniq := []string{}
	current := ""
	for _, word := range words {
		if word == current {
			continue
		}
		if current != "" {
			uniq = append(uniq, current)
			current = ""
		}
		current = word
	}
	if current != "" {
		uniq = append(uniq, current)
	}
	return uniq
}
