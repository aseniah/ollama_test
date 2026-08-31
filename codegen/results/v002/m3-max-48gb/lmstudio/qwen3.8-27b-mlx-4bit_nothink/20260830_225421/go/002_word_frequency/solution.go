package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
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
		// Strip all punctuation - keep only letters and spaces
		var cleaned strings.Builder
		for _, ch := range line {
			if unicode.IsLetter(ch) || unicode.IsSpace(ch) {
				cleaned.WriteRune(ch)
			}
		}
		// Split into words by spaces
		words := strings.Fields(cleaned.String())
		for _, w := range words {
			counts[w]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create slice of wordCount
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

	for _, wc := range result {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}