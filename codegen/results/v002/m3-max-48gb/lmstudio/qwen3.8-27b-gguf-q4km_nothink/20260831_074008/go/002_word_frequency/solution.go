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

	scanner := bufio.NewScanner(file)
	wordCount := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		// Extract only alphabetic characters into words
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
				wordCount[w]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create a slice of word-count pairs for sorting
	type wordCount struct {
		word  string
		count int
	}

	var pairs []wordCount
	for w, c := range wordCount {
		pairs = append(pairs, wordCount{w, c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	// Output
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, p := range pairs {
		fmt.Fprintf(out, "%s: %d\n", p.word, p.count)
	}
}