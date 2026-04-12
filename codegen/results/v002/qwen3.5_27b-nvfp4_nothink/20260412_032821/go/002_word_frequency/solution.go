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
		return // In a real app, you might log the error, but per instructions: do not output anything else
	}
	defer file.Close()

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line) // Splits by whitespace

		for _, word := range words {
			// Convert to lowercase
			cleanWord := strings.ToLower(word)
			// Keep only letters
			var builder strings.Builder
			for _, r := range cleanWord {
				if 'a' <= r && r <= 'z' {
					builder.WriteRune(r)
				}
			}

			finalWord := builder.String()
			if len(finalWord) > 0 {
				wordCounts[finalWord]++
			}
		}
	}

	// Convert map to slice for sorting
	type wordCount struct {
		word  string
		count int
	}

	var sorted []wordCount
	for w, c := range wordCounts {
		sorted = append(sorted, wordCount{w, c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].count != sorted[j].count {
			return sorted[i].count > sorted[j].count
		}
		return sorted[i].word < sorted[j].word
	})

	// Output results
	for _, item := range sorted {
		fmt.Printf("%s: %d\n", item.word, item.count)
	}
}