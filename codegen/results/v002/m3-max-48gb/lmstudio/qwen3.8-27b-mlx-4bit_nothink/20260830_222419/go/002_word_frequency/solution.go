package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	counts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		
		// Strip all punctuation (keep only letters)
		var filtered strings.Builder
		for _, ch := range line {
			if (ch >= 'a' && ch <= 'z') {
				filtered.WriteRune(ch)
			} else {
				// If it's not a letter, add a space to separate words
				// But we need to be careful: we want to split on non-letter characters
				filtered.WriteRune(' ')
			}
		}
		
		// Split by spaces
		words := strings.Fields(filtered.String())
		for _, word := range words {
			if word != "" {
				counts[word]++
			}
		}
	}

	// Create slice of wordCount structs
	result := make([]wordCount, 0, len(counts))
	for word, count := range counts {
		result = append(result, wordCount{word: word, count: count})
	}

	// Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(result, func(i, j int) bool {
		if result[i].count != result[j].count {
			return result[i].count > result[j].count
		}
		return result[i].word < result[j].word
	})

	// Output
	for _, wc := range result {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}