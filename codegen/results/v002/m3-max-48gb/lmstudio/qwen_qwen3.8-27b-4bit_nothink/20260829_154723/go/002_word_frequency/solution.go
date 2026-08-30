package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCounts := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		// Split into words by any non-letter characters
		// We want to strip all punctuation and keep only letters
		// So we can split by non-letter runes and filter empty strings
		var words []string
		var currentWord strings.Builder
		for _, r := range line {
			if unicode.IsLetter(r) {
				currentWord.WriteRune(r)
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
		for _, word := range words {
			if word != "" {
				wordCounts[word]++
			}
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading file")
		return
	}

	// Create a slice of word-count pairs
	type wordCount struct {
		word  string
		count int
	}

	var pairs []wordCount
	for word, count := range wordCounts {
		pairs = append(pairs, wordCount{word, count})
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

	for _, pair := range pairs {
		fmt.Fprintf(out, "%s: %d\n", pair.word, pair.count)
	}
}