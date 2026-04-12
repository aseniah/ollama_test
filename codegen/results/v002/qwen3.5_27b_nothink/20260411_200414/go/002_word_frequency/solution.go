package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Map to store word frequencies
	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert line to lowercase
		line = strings.ToLower(line)

		// Replace non-letter characters with spaces
		var cleaned strings.Builder
		for _, r := range line {
			if (r >= 'a' && r <= 'z') {
				cleaned.WriteRune(r)
			} else {
				cleaned.WriteRune(' ')
			}
		}
		cleanLine := cleaned.String()

		// Split into words
		words := strings.Fields(cleanLine)
		for _, word := range words {
			if len(word) > 0 {
				wordCounts[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Convert map to slice for sorting
	type wordCount struct {
		word  string
		count int
	}
	var sortedWords []wordCount
	for word, count := range wordCounts {
		sortedWords = append(sortedWords, wordCount{word, count})
	}

	// Sort: by count descending, then alphabetically ascending
	sort.Slice(sortedWords, func(i, j int) bool {
		if sortedWords[i].count != sortedWords[j].count {
			return sortedWords[i].count > sortedWords[j].count
		}
		return sortedWords[i].word < sortedWords[j].word
	})

	// Output results
	for _, wc := range sortedWords {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}