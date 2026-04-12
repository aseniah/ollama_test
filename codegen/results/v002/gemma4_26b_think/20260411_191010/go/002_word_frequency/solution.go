package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// wordEntry holds the word and its corresponding frequency.
type wordEntry struct {
	word  string
	count int
}

func main() {
	// Read the file from the specified path
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		// If the file cannot be read, we terminate.
		// The requirement specifies "Do not output anything else".
		return
	}

	// Convert all text to lowercase as per rule 1
	lowerContent := strings.ToLower(string(content))

	// Rule 2: Strip all punctuation (keep only letters).
	// We use a regular expression to replace any non-alphabetic character (a-z)
	// with a space, ensuring we only process letters.
	reg := regexp.MustCompile(`[^a-z]+`)
	cleanedContent := reg.ReplaceAllString(lowerContent, " ")

	// Split the cleaned string into individual words using whitespace as a delimiter
	words := strings.Fields(cleanedContent)

	// Count the frequency of each word
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	// Transfer the map entries into a slice of structs to allow for sorting
	var sortedEntries []wordEntry
	for w, c := range counts {
		sortedEntries = append(sortedEntries, wordEntry{word: w, count: c})
	}

	// Rule 4: Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(sortedEntries, func(i, j int) bool {
		if sortedEntries[i].count != sortedEntries[j].count {
			// Primary sort: count descending
			return sortedEntries[i].count > sortedEntries[j].count
		}
		// Secondary sort: word ascending
		return sortedEntries[i].word < sortedEntries[j].word
	})

	// Rule 3: Output one line per unique word in the format: word: count
	for _, entry := range sortedEntries {
		fmt.Printf("%s: %d\n", entry.word, entry.count)
	}
}