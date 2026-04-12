package main

import (
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	// Read the input file
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		os.Exit(1)
	}

	text := string(content)

	// Replace all non-letter characters with space
	text = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return ' '
	}, text)

	// Split into words
	words := strings.Fields(text)

	// Count frequency
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// Prepare results slice
	type WordCount struct {
		Word string
		Count int
	}

	var results []WordCount
	for word, count := range wordCount {
		results = append(results, WordCount{Word: word, Count: count})
	}

	// Sort by count descending, then by word ascending alphabetically
	sort.Slice(results, func(i, j int) bool {
		if results[i].Count != results[j].Count {
			return results[i].Count > results[j].Count
		}
		return results[i].Word < results[j].Word
	})

	// Print results
	for _, r := range results {
		fmt.Printf("%s: %d\n", r.Word, r.Count)
	}
}