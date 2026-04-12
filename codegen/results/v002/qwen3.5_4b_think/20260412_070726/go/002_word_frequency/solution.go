package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Read entire file
	content, err := bufio.ReadAll(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Convert all words to lowercase
	content = strings.ToLower(content)

	// Remove punctuation (keep only letters, spaces)
	// Replace any non-alphabetic character with space
	content = regexp.MustCompile(`[^a-z]+`).ReplaceAllString(content, " ")

	// Split into words
	words := strings.Fields(content)

	// Count word frequencies
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	// Convert to slice for sorting
	type wordCount struct {
		word string
		count int
	}

	var items []wordCount
	for word, count := range counts {
		items = append(items, wordCount{word: word, count: count})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(items, func(i, j int) bool {
		if items[i].count != items[j].count {
			return items[i].count > items[j].count
		}
		return items[i].word < items[j].word
	})

	// Output
	for _, item := range items {
		fmt.Printf("%s: %d\n", item.word, item.count)
	}
}