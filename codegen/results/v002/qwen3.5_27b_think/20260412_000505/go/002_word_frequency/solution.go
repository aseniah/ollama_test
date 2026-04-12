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
	// Open the file
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the file content
	scanner := bufio.NewScanner(file)
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString(" ")
	}

	// Convert to lowercase
	text := strings.ToLower(content.String())

	// Count word frequencies
	wordCount := make(map[string]int)
	var currentWord strings.Builder

	for _, ch := range text {
		if unicode.IsLetter(ch) {
			currentWord.WriteRune(ch)
		} else {
			if currentWord.Len() > 0 {
				word := currentWord.String()
				wordCount[word]++
				currentWord.Reset()
			}
		}
	}

	// Don't forget the last word
	if currentWord.Len() > 0 {
		word := currentWord.String()
		wordCount[word]++
	}

	// Convert to slice for sorting
	type wordFreq struct {
		Word  string
		Count int
	}

	words := make([]wordFreq, 0, len(wordCount))
	for word, count := range wordCount {
		words = append(words, wordFreq{Word: word, Count: count})
	}

	// Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(words, func(i, j int) bool {
		if words[i].Count != words[j].Count {
			return words[i].Count > words[j].Count
		}
		return words[i].Word < words[j].Word
	})

	// Output results
	for _, w := range words {
		fmt.Printf("%s: %d\n", w.Word, w.Count)
	}
}