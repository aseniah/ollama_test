package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func main() {
	data, err := os.ReadFile("input/text.txt")
	if err != nil {
		panic(err)
	}

	// Convert to lowercase
	content := strings.ToLower(string(data))
	
	// Strip all non-letter characters and replace with spaces
	nonLetters := regexp.MustCompile(`[^a-z]+`)
	content = nonLetters.ReplaceAllString(content, " ")

	// Split into words
	words := strings.Fields(content)

	// Count frequencies
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	// Convert map to slice for sorting
	var result []WordCount
	for word, count := range counts {
		result = append(result, WordCount{word, count})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(result, func(i, j int) bool {
		if result[i].Count != result[j].Count {
			return result[i].Count > result[j].Count
		}
		return result[i].Word < result[j].Word
	})

	// Output results
	for _, item := range result {
		fmt.Printf("%s: %d\n", item.Word, item.Count)
	}
}