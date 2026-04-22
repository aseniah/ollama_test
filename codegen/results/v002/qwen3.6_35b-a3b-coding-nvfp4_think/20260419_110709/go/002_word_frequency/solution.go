package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	text := string(data)
	// Split text into words, treating any non-letter as a separator
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	counts := make(map[string]int)
	for _, w := range words {
		lower := strings.ToLower(w)
		if lower == "" {
			continue
		}
		counts[lower]++
	}

	// Collect unique words into a slice for sorting
	unique := make([]string, 0, len(counts))
	for w := range counts {
		unique = append(unique, w)
	}

	// Sort by frequency descending, then alphabetically ascending
	sort.Slice(unique, func(i, j int) bool {
		ci := counts[unique[i]]
		cj := counts[unique[j]]
		if ci != cj {
			return ci > cj
		}
		return unique[i] < unique[j]
	})

	// Output results
	for _, w := range unique {
		fmt.Printf("%s: %d\n", w, counts[w])
	}
}