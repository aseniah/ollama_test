package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func main() {
	// Read the file
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		return
	}

	// Convert to lowercase
	text := strings.ToLower(string(content))

	// Remove all characters that are not letters or whitespace
	// We replace non-letters (except spaces/newlines) with spaces to avoid merging words
	reg := regexp.MustCompile(`[^a-z\s]`)
	text = reg.ReplaceAllString(text, " ")

	// Split by whitespace
	words := strings.Fields(text)

	// Count frequencies
	countsMap := make(map[string]int)
	for _, w := range words {
		countsMap[w]++
	}

	// Convert map to slice for sorting
	var sortedCounts []wordCount
	for w, c := range countsMap {
		sortedCounts = append(sortedCounts, wordCount{w, c})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(sortedCounts, func(i, j int) bool {
		if sortedCounts[i].count != sortedCounts[j].count {
			return sortedCounts[i].count > sortedCounts[j].count
		}
		return sortedCounts[i].word < sortedCounts[j].word
	})

	// Output results
	for _, wc := range sortedCounts {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}