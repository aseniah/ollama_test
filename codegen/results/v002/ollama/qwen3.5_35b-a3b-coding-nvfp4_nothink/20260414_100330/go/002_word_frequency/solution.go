package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// Regex to match sequences of letters (a-z, A-Z)
	// We'll use this to extract words
	wordRegex := regexp.MustCompile(`[A-Za-z]+`)

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Find all words in the line
		words := wordRegex.FindAllString(line, -1)
		for _, w := range words {
			lower := strings.ToLower(w)
			wordCounts[lower]++
		}
	}

	// Convert map to slice of structs for sorting
	type wordCount struct {
		word  string
		count int
	}

	var counts []wordCount
	for word, count := range wordCounts {
		counts = append(counts, wordCount{word, count})
	}

	// Sort: descending by count, then ascending alphabetically for ties
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count == counts[j].count {
			return counts[i].word < counts[j].word
		}
		return counts[i].count > counts[j].count
	})

	for _, wc := range counts {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}