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

		// Strip punctuation, keeping only letters
		// We use a regex to match sequences of letters
		re := regexp.MustCompile(`[^a-z]+`)
		words := re.FindAllString(line, -1)

		for _, word := range words {
			if word == "" {
				continue
			}
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create a slice of the map items to sort
	type wordCountEntry struct {
		word  string
		count int
	}
	
	var entries []wordCountEntry
	for word, count := range wordCount {
		entries = append(entries, wordCountEntry{word: word, count: count})
	}

	// Sort by count descending, then by word alphabetically ascending
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		return entries[i].word < entries[j].word
	})

	// Output results
	for _, entry := range entries {
		fmt.Printf("%s: %d\n", entry.word, entry.count)
	}
}