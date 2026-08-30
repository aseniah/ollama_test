package main

import (
	"bufio"
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
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	counts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line
		for _, ch := range line {
			// We'll build words by filtering
		}
		
		// Extract words: keep only letters, treat non-letters as delimiters
		var currentWord strings.Builder
		for _, ch := range line {
			if unicode.IsLetter(ch) {
				currentWord.WriteRune(unicode.ToLower(ch))
			} else {
				if currentWord.Len() > 0 {
					word := currentWord.String()
					counts[word]++
					currentWord.Reset()
				}
			}
		}
		// Don't forget the last word in the line
		if currentWord.Len() > 0 {
			word := currentWord.String()
			counts[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	// Convert to slice for sorting
	var results []wordCount
	for w, c := range counts {
		results = append(results, wordCount{word: w, count: c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(results, func(i, j int) bool {
		if results[i].count != results[j].count {
			return results[i].count > results[j].count
		}
		return results[i].word < results[j].word
	})

	// Output
	for _, rc := range results {
		fmt.Printf("%s: %d\n", rc.word, rc.count)
	}
}