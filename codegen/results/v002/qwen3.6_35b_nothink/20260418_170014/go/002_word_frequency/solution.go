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
	// Read the file
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Count word frequencies
	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split into tokens
		// We need to extract words: sequences of letters only
		var currentWord strings.Builder
		for _, r := range line {
			if unicode.IsLetter(r) {
				currentWord.WriteRune(unicode.ToLower(r))
			} else {
				if currentWord.Len() > 0 {
					word := currentWord.String()
					wordCount[word]++
					currentWord.Reset()
				}
			}
		}
		// Don't forget the last word if the line ends with a letter
		if currentWord.Len() > 0 {
			word := currentWord.String()
			wordCount[word]++
			currentWord.Reset()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Collect words and their counts
	type wordFreq struct {
		word  string
		count int
	}

	var words []wordFreq
	for word, count := range wordCount {
		words = append(words, wordFreq{word: word, count: count})
	}

	// Sort by count descending, then alphabetically ascending for ties
	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}
		return words[i].word < words[j].word
	})

	// Output
	for _, wf := range words {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}