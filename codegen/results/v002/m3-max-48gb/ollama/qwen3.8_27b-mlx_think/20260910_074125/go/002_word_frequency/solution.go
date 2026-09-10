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
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	text := string(data)
	words := strings.Fields(text)

	counts := make(map[string]int)
	for _, w := range words {
		// Convert to lowercase and strip non-letter characters
		cleaned := make([]rune, 0, len(w))
		for _, ch := range strings.ToLower(w) {
			if unicode.IsLetter(ch) {
				cleaned = append(cleaned, ch)
			}
		}
		if len(cleaned) > 0 {
			counts[string(cleaned)]++
		}
	}

	// Build slice of (word, count) pairs
	type wordCount struct {
		word  string
		count int
	}
	pairs := make([]wordCount, 0, len(counts))
	for w, c := range counts {
		pairs = append(pairs, wordCount{w, c})
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