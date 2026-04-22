package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	content := strings.ToLower(string(data))
	
	// Split into words by non-letter characters
	words := strings.FieldsFunc(content, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	var uniqueWords []string
	for w := range counts {
		uniqueWords = append(uniqueWords, w)
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		ci := counts[uniqueWords[i]]
		cj := counts[uniqueWords[j]]
		if ci != cj {
			return ci > cj // descending count
		}
		return uniqueWords[i] < uniqueWords[j] // ascending alphabetically
	})

	for _, w := range uniqueWords {
		fmt.Printf("%s: %d\n", w, counts[w])
	}
}