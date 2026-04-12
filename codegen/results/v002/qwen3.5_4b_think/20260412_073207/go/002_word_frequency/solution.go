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
	// Read the input file
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Map to store word frequencies
	wordCounts := make(map[string]int)

	// Regex to match words (sequences of letters only)
	wordRegex := regexp.MustCompile(`[a-z]+`)

	// Read and process the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text()) // Convert to lowercase
		words := wordRegex.FindAllString(line, -1)
		for _, word := range words {
			if len(word) > 0 {
				wordCounts[word]++
			}
		}
	}

	// Convert map to slice for sorting
	var words []struct {
		word string
		count int
	}
	for word, count := range wordCounts {
		words = append(words, struct {
			word string
			count int
		}{word, count})
	}

	// Sort: by count descending, then alphabetically ascending
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