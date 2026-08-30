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
	f, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// Extract words: keep only letters, convert to lowercase
		// Split by non-letter characters
		var words []string
		var current strings.Builder
		for _, r := range line {
			if unicode.IsLetter(r) {
				current.WriteRune(unicode.ToLower(r))
			} else {
				if current.Len() > 0 {
					words = append(words, current.String())
					current.Reset()
				}
			}
		}
		if current.Len() > 0 {
			words = append(words, current.String())
		}

		for _, w := range words {
			wordCount[w]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create a slice of word-count pairs
	type wordEntry struct {
		word  string
		count int
	}
	entries := make([]wordEntry, 0, len(wordCount))
	for w, c := range wordCount {
		entries = append(entries, wordEntry{word: w, count: c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		return entries[i].word < entries[j].word
	})

	// Output
	for _, e := range entries {
		fmt.Printf("%s: %d\n", e.word, e.count)
	}
}