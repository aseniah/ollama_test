package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type WordCount struct {
	Word  string
	Count int
}

func main() {
	// Read the file
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		panic(err)
	}

	// Convert to lowercase
	text := strings.ToLower(string(content))

	// Strip punctuation: keep only letters, replace others with space
	var cleaned strings.Builder
	cleaned.Grow(len(text))
	for _, r := range text {
		if unicode.IsLetter(r) {
			cleaned.WriteRune(r)
		} else {
			cleaned.WriteRune(' ')
		}
	}

	// Split into words based on whitespace
	words := strings.Fields(cleaned.String())

	// Count frequencies
	frequencies := make(map[string]int)
	for _, word := range words {
		frequencies[word]++
	}

	// Convert map to slice for sorting
	var wordCounts []WordCount
	for word, count := range frequencies {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}

	// Sort by count descending, then word alphabetically ascending
	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].Count != wordCounts[j].Count {
			return wordCounts[i].Count > wordCounts[j].Count
		}
		return wordCounts[i].Word < wordCounts[j].Word
	})

	// Output results
	for _, wc := range wordCounts {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}