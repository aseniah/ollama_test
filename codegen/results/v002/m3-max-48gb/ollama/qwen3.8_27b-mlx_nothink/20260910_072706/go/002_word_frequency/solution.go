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
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	content := string(data)
	lowerContent := strings.ToLower(content)

	// Split by non-letter characters
	var words []string
	var currentWord []rune

	for _, ch := range lowerContent {
		if unicode.IsLetter(ch) {
			currentWord = append(currentWord, ch)
		} else {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = nil
			}
		}
	}
	// Don't forget the last word
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	// Count frequencies
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	// Create a slice of pairs for sorting
	type wordCount struct {
		word  string
		count int
	}

	var pairs []wordCount
	for w, c := range freq {
		pairs = append(pairs, wordCount{w, c})
	}

	// Sort by count descending, then alphabetically ascending for ties
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	// Output
	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.word, p.count)
	}
}