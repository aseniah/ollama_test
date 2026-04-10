package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// Read the file content
	data, err := os.ReadFile("input/text.txt")
	if err != nil {
		panic(err)
	}

	// Convert to lowercase
	text := strings.ToLower(string(data))

	// Strip all punctuation (keep only letters a-z)
	// Replace any sequence of non-lowercase-letters with a space
	re := regexp.MustCompile(`[^a-z]+`)
	text = re.ReplaceAllString(text, " ")

	// Split into words by whitespace
	words := strings.Fields(text)

	// Count word frequencies
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	// Prepare slice for sorting
	type WordCount struct {
		word  string
		count int
	}
	
	var sortedWords []WordCount
	for w, c := range counts {
		sortedWords = append(sortedWords, WordCount{word: w, count: c})
	}

	// Sort: Count descending, then Word ascending
	sort.Slice(sortedWords, func(i, j int) bool {
		if sortedWords[i].count != sortedWords[j].count {
			return sortedWords[i].count > sortedWords[j].count
		}
		return sortedWords[i].word < sortedWords[j].word
	})

	// Output results
	for _, wc := range sortedWords {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}