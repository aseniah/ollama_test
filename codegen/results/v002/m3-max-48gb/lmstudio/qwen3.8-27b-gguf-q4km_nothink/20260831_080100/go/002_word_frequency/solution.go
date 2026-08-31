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

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split by whitespace
		words := strings.Fields(line)
		for _, word := range words {
			// Convert to lowercase and strip non-letters
			var cleaned strings.Builder
			for _, ch := range word {
				if unicode.IsLetter(ch) {
					cleaned.WriteRune(unicode.ToLower(ch))
				}
			}
			cleanedWord := cleaned.String()
			if cleanedWord != "" {
				wordCount[cleanedWord]++
			}
		}
	}

	// Create a slice of word-count pairs
	type WordCount struct {
		Word  string
		Count int
	}

	var pairs []WordCount
	for w, c := range wordCount {
		pairs = append(pairs, WordCount{Word: w, Count: c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count != pairs[j].Count {
			return pairs[i].Count > pairs[j].Count
		}
		return pairs[i].Word < pairs[j].Word
	})

	// Output
	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.Word, p.Count)
	}
}