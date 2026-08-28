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
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		// Split into words
		fields := strings.Fields(line)
		for _, field := range fields {
			// Strip punctuation, keep only letters
			word := ""
			for _, ch := range field {
				if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
					word += string(ch)
				}
			}
			if word != "" {
				wordCount[word]++
			}
		}
	}

	// Create a slice of word-count pairs
	type wordCount struct {
		word  string
		count int
	}
	var pairs []wordCount
	for w, c := range wordCount {
		pairs = append(pairs, wordCount{w, c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	// Output
	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.word, p.count)
	}
}