package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type wordCount struct {
	word  string
	count int
}

func main() {
	// Read the file input/text.txt
	data, err := os.ReadFile("input/text.txt")
	if err != nil {
		return
	}

	content := string(data)

	// Clean the text: keep only letters and spaces, convert to lowercase
	var builder strings.Builder
	for _, r := range content {
		if unicode.IsLetter(r) {
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(' ')
		}
	}

	// Split into words based on whitespace
	words := strings.Fields(builder.String())

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

	// Sort by count descending, then alphabetically ascending
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