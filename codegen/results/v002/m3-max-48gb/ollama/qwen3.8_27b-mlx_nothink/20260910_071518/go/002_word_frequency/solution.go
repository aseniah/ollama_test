package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	// Read the file
	data, err := os.ReadFile("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Split into words by whitespace
	lines := string(data)
	words := strings.Fields(lines)

	// Count word frequencies
	counts := make(map[string]int)

	for _, word := range words {
		// Convert to lowercase
		word = strings.ToLower(word)
		// Strip punctuation - keep only letters
		var builder strings.Builder
		for _, r := range word {
			if unicode.IsLetter(r) {
				builder.WriteRune(r)
			}
		}
		cleanWord := builder.String()
		if cleanWord != "" {
			counts[cleanWord]++
		}
	}

	// Create a slice of (word, count) pairs
	type wordCount struct {
		word  string
		count int
	}

	var pairs []wordCount
	for w, c := range counts {
		pairs = append(pairs, wordCount{word: w, count: c})
	}

	// Sort by count descending, then alphabetically ascending
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