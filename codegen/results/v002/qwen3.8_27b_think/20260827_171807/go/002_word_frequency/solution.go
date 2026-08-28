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
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, field := range fields {
			// Convert to lowercase and strip non-letter characters
			var cleaned strings.Builder
			for _, ch := range strings.ToLower(field) {
				if unicode.IsLetter(ch) {
					cleaned.WriteRune(ch)
				}
			}
			word := cleaned.String()
			if word != "" {
				wordCount[word]++
			}
		}
	}

	// Create slice of word-count pairs
	type entry struct {
		word  string
		count int
	}
	entries := make([]entry, 0, len(wordCount))
	for w, c := range wordCount {
		entries = append(entries, entry{w, c})
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