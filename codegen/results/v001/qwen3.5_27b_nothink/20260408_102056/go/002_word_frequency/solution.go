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
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			// Convert to lowercase and strip punctuation
			var sb strings.Builder
			for _, r := range word {
				if unicode.IsLetter(r) {
					sb.WriteRune(unicode.ToLower(r))
				}
			}
			cleanWord := sb.String()

			if cleanWord != "" {
				wordCounts[cleanWord]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	// Create a slice of word/count pairs for sorting
	type wordCount struct {
		word  string
		count int
	}

	var results []wordCount
	for word, count := range wordCounts {
		results = append(results, wordCount{word, count})
	}

	// Sort by count descending, then by word alphabetically ascending
	sort.Slice(results, func(i, j int) bool {
		if results[i].count != results[j].count {
			return results[i].count > results[j].count
		}
		return results[i].word < results[j].word
	})

	// Output results
	for _, item := range results {
		fmt.Printf("%s: %d\n", item.word, item.count)
	}
}