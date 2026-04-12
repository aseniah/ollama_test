package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// WordCount holds the word and its frequency.
type WordCount struct {
	word  string
	count int
}

func main() {
	// Read the file content
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		return
	}

	// Map to store word frequencies
	counts := make(map[string]int)
	var sb strings.Builder

	// Iterate over the content rune by rune to handle Unicode correctly
	for _, r := range string(content) {
		if unicode.IsLetter(r) {
			// Convert to lowercase and append to buffer
			sb.WriteRune(unicode.ToLower(r))
		} else {
			// Non-letter character acts as a delimiter
			if sb.Len() > 0 {
				word := sb.String()
				counts[word]++
				sb.Reset()
			}
		}
	}
	// Process the last word if the file doesn't end with a delimiter
	if sb.Len() > 0 {
		word := sb.String()
		counts[word]++
	}

	// Convert map to a slice for sorting
	entries := make([]WordCount, 0, len(counts))
	for w, c := range counts {
		entries = append(entries, WordCount{word: w, count: c})
	}

	// Sort the entries
	sort.Slice(entries, func(i, j int) bool {
		// Sort by count descending
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		// For ties, sort alphabetically ascending
		return entries[i].word < entries[j].word
	})

	// Output the results
	for _, e := range entries {
		fmt.Printf("%s: %d\n", e.word, e.count)
	}
}