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
		return // Do not output anything else if file cannot be opened
	}
	defer file.Close()

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		
		for _, word := range words {
			// Convert to lowercase
			cleaned := strings.ToLower(word)
			
			// Strip punctuation: keep only letters
			filtered := strings.Builder{}
			for _, ch := range cleaned {
				if unicode.IsLetter(ch) {
					filtered.WriteRune(ch)
				}
			}
			finalWord := filtered.String()

			if finalWord != "" {
				wordCounts[finalWord]++
			}
		}
	}

	// Convert map to slice for sorting
	type wordCount struct {
		Word  string
		Count int
	}

	sorted := make([]wordCount, 0, len(wordCounts))
	for w, c := range wordCounts {
		sorted = append(sorted, wordCount{Word: w, Count: c})
	}

	// Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Count != sorted[j].Count {
			return sorted[i].Count > sorted[j].Count
		}
		return sorted[i].Word < sorted[j].Word
	})

	// Output results
	for _, item := range sorted {
		fmt.Printf("%s: %d\n", item.Word, item.Count)
	}
}