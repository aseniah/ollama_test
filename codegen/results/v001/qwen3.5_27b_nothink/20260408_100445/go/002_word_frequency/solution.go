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

	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Normalize: convert to lowercase
		line = strings.ToLower(line)

		// Remove punctuation: keep only letters
		var cleanedLine strings.Builder
		for _, r := range line {
			if unicode.IsLetter(r) {
				cleanedLine.WriteRune(r)
			} else {
				// Replace non-letter with space to separate words
				cleanedLine.WriteRune(' ')
			}
		}

		words := strings.Fields(cleanedLine.String())
		for _, word := range words {
			if word != "" {
				wordCounts[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Convert map to slice for sorting
	type wordCount struct {
		Word  string
		Count int
	}
	var results []wordCount
	for w, c := range wordCounts {
		results = append(results, wordCount{Word: w, Count: c})
	}

	// Sort: by count descending, then by word ascending
	sort.Slice(results, func(i, j int) bool {
		if results[i].Count != results[j].Count {
			return results[i].Count > results[j].Count
		}
		return results[i].Word < results[j].Word
	})

	// Output
	for _, item := range results {
		fmt.Printf("%s: %d\n", item.Word, item.Count)
	}
}