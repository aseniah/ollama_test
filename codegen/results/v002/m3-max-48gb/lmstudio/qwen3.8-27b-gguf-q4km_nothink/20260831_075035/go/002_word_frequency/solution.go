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

	scanner := bufio.NewScanner(file)
	counts := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		// Split by whitespace
		words := strings.Fields(line)
		for _, w := range words {
			// Convert to lowercase
			w = strings.ToLower(w)
			// Strip all punctuation - keep only letters
			var cleaned strings.Builder
			for _, r := range w {
				if unicode.IsLetter(r) {
					cleaned.WriteRune(r)
				}
			}
			word := cleaned.String()
			if word != "" {
				counts[word]++
			}
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

	// Sort by count descending, then alphabetically ascending
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