package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		// Extract words: keep only letters, group consecutive letters
		var words []string
		var current []rune
		for _, r := range line {
			if unicode.IsLetter(r) {
				current = append(current, r)
			} else {
				if len(current) > 0 {
					words = append(words, string(current))
					current = nil
				}
			}
		}
		if len(current) > 0 {
			words = append(words, string(current))
		}
		for _, w := range words {
			wordCount[w]++
		}
	}

	// Create a slice of (word, count) pairs
	type wordCount struct {
		word  string
		count int
	}
	items := make([]wordCount, 0, len(wordCount))
	for w, c := range wordCount {
		items = append(items, wordCount{w, c})
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