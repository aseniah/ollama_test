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

	scanner := bufio.NewScanner(file)
	wordCounts := make(map[string]int)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		// Strip all punctuation, keep only letters
		// We'll split by non-letter characters
		var words []string
		var currentWord strings.Builder

		for _, ch := range line {
			if unicode.IsLetter(ch) {
				currentWord.WriteRune(ch)
			} else {
				if currentWord.Len() > 0 {
					words = append(words, currentWord.String())
					currentWord.Reset()
				}
			}
		}
		if currentWord.Len() > 0 {
			words = append(words, currentWord.String())
		}

		for _, w := range words {
			if w != "" {
				wordCounts[w]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create a slice of word-count pairs for sorting
	type wordCount struct {
		word  string
		count int
	}

	var pairs []wordCount
	for w, c := range wordCounts {
		pairs = append(pairs, wordCount{word: w, count: c})
	}

	// Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	// Output
	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.word, p.count)
	}
}