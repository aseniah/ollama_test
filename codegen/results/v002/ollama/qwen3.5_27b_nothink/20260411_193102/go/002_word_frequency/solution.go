package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// Read the entire file
	content, err := ioutil.ReadFile("input/text.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert to lowercase
	text := strings.ToLower(string(content))

	// Remove all non-letter characters (keep only a-z)
	re := regexp.MustCompile(`[^a-z]+`)
	cleaned := re.ReplaceAllString(text, " ")

	// Split into words (by whitespace)
	words := strings.Fields(cleaned)

	// Count frequencies
	counts := make(map[string]int)
	for _, word := range words {
		if word == "" {
			continue
		}
		counts[word]++
	}

	// Convert map to slice for sorting
	type WordCount struct {
		Word  string
		Count int
	}

	var sorted []WordCount
	for w, c := range counts {
		sorted = append(sorted, WordCount{w, c})
	}

	// Sort: descending by count, then ascending by word
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Count != sorted[j].Count {
			return sorted[i].Count > sorted[j].Count
		}
		return sorted[i].Word < sorted[j].Word
	})

	// Output
	for _, item := range sorted {
		fmt.Printf("%s: %d\n", item.Word, item.Count)
	}
}