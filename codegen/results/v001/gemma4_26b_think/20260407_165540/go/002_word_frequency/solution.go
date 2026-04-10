package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// wordFreq stores the word and its occurrence count for sorting purposes.
type wordFreq struct {
	word  string
	count int
}

func main() {
	// Define the path to the input file.
	filePath := "input/text.txt"

	// Read the entire content of the file.
	content, err := os.ReadFile(filePath)
	if err != nil {
		// If the file cannot be read (e.g., doesn't exist), we exit silently 
		// as per the requirement to not output anything else.
		return
	}

	// Map to store the frequency of each word.
	counts := make(map[string]int)

	// Split the content into initial tokens based on whitespace.
	words := strings.Fields(string(content))

	for _, word := range words {
		// Build a new string containing only letters from the current word.
		var sb strings.Builder
		for _, r := range word {
			if unicode.IsLetter(r) {
				// Convert to lowercase immediately while building.
				sb.WriteRune(unicode.ToLower(r))
			}
		}

		cleanedWord := sb.String()

		// Only process the word if it contains letters after stripping punctuation.
		if cleanedWord != "" {
			counts[cleanedWord]++
		}
	}

	// Convert the map into a slice of structs so we can sort it.
	var freqList []wordFreq
	for word, count := range counts {
		freqList = append(freqList, wordFreq{word: word, count: count})
	}

	// Sort the slice:
	// 1. By count descending.
	// 2. By word alphabetically ascending (for ties).
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count
		}
		return freqList[i].word < freqList[j].word
	})

	// Output the results in the format "word: count".
	for _, entry := range freqList {
		fmt.Printf("%s: %d\n", entry.word, entry.count)
	}
}