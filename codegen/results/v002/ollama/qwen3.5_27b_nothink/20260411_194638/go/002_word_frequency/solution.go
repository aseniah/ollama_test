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
		// Split line into words manually to handle punctuation attached to words
		// and ensure we process character by character for cleaning
		words := strings.Fields(line)
		for _, w := range words {
			cleaned := cleanWord(w)
			if cleaned != "" {
				wordCounts[cleaned]++
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

	var result []wordCount
	for word, count := range wordCounts {
		result = append(result, wordCount{Word: word, Count: count})
	}

	// Sort: descending by count, then ascending alphabetically by word
	sort.Slice(result, func(i, j int) bool {
		if result[i].Count != result[j].Count {
			return result[i].Count > result[j].Count
		}
		return result[i].Word < result[j].Word
	})

	// Output
	for _, item := range result {
		fmt.Printf("%s: %d\n", item.Word, item.Count)
	}
}

func cleanWord(s string) string {
	// Convert to lowercase and keep only letters
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			sb.WriteRune(unicode.ToLower(r))
		}
	}
	return sb.String()
}