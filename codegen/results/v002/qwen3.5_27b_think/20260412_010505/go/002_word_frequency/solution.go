package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	// Read the content of the file
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		return
	}

	// Map to count word frequencies
	frequencies := make(map[string]int)
	var wordBuilder strings.Builder

	// Iterate over the text as runes to handle unicode and lowercase correctly
	for _, r := range string(content) {
		if unicode.IsLetter(r) {
			// Append lowercase letter to current word
			wordBuilder.WriteRune(unicode.ToLower(r))
		} else {
			// Non-letter acts as a separator
			if wordBuilder.Len() > 0 {
				word := wordBuilder.String()
				frequencies[word]++
				wordBuilder.Reset()
			}
		}
	}
	// Don't forget the last word if file doesn't end with punctuation
	if wordBuilder.Len() > 0 {
		word := wordBuilder.String()
		frequencies[word]++
	}

	// Convert map to slice for sorting
	type WordEntry struct {
		Word  string
		Count int
	}
	entries := make([]WordEntry, 0, len(frequencies))
	for word, count := range frequencies {
		entries = append(entries, WordEntry{Word: word, Count: count})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count != entries[j].Count {
			return entries[i].Count > entries[j].Count
		}
		return entries[i].Word < entries[j].Word
	})

	// Output results
	for _, entry := range entries {
		fmt.Printf("%s: %d\n", entry.Word, entry.Count)
	}
}