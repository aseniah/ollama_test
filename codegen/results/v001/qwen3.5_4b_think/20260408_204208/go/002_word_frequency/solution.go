package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Map to store word counts
	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Convert to lowercase
		line = strings.ToLower(line)

		// Remove punctuation and split into words
		// Keep only alphabetic characters and spaces
		words := regexp.MustCompile(`[^a-z\s]+`).ReplaceAllString(line, " ")
		words = regexp.MustCompile(`\s+`).ReplaceAllString(words, " ")

		// Split and count
		wordList := strings.Split(words, " ")
		for _, word := range wordList {
			if word != "" {
				wordCount[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Convert map to slice for sorting
	words := make([]struct {
		word   string
		count  int
	}, 0, len(wordCount))
	for word, count := range wordCount {
		words = append(words, struct {
			word   string
			count  int
		}{word, count})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}
		return words[i].word < words[j].word
	})

	// Output results
	for _, w := range words {
		fmt.Printf("%s: %d\n", w.word, w.count)
	}
}