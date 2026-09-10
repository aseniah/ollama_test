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

	content := string(data)

	// Split into words by whitespace
	rawWords := strings.Fields(content)

	// Map to store word frequencies
	freq := make(map[string]int)

	for _, raw := range rawWords {
		// Convert to lowercase
		// Strip all punctuation (keep only letters)
		var cleaned strings.Builder
		for _, ch := range raw {
			if unicode.IsLetter(ch) {
				cleaned.WriteRune(unicode.ToLower(ch))
			}
		}
		word := cleaned.String()
		if word != "" {
			freq[word]++
		}
	}

	// Create a slice of word-count pairs
	type WordCount struct {
		Word  string
		Count int
	}

	var pairs []WordCount
	for word, count := range freq {
		pairs = append(pairs, WordCount{Word: word, Count: count})
	}

	// Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count != pairs[j].Count {
			return pairs[i].Count > pairs[j].Count
		}
		return pairs[i].Word < pairs[j].Word
	})

	// Output
	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.Word, p.Count)
	}
}