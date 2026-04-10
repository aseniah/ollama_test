package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// wordEntry holds the word and its frequency for sorting purposes
type wordEntry struct {
	word  string
	count int
}

func main() {
	// Open the file
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	
	// Regex to keep only letters (strips punctuation and numbers)
	// This matches non-letter characters to replace them with space
	reg, err := regexp.Compile(`[^a-zA-Z\s]+`)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		// 1. Convert to lowercase
		line = strings.ToLower(line)
		
		// 2. Strip punctuation (replace non-letters with space)
		line = reg.ReplaceAllString(line, " ")
		
		// 3. Split into words and count
		words := strings.Fields(line)
		for _, w := range words {
			counts[w]++
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	// Convert map to a slice of structs for sorting
	var entries []wordEntry
	for word, count := range counts {
		entries = append(entries, wordEntry{word, count})
	}

	// 4. Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		return entries[i].word < entries[j].word
	})

	// 5. Output formatted results
	for _, entry := range entries {
		fmt.Printf("%s: %d\n", entry.word, entry.count)
	}
}