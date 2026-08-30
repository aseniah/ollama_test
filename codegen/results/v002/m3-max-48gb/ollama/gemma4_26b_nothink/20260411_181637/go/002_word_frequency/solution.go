package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// WordFreq stores a word and its occurrence count
type WordFreq struct {
	Word  string
	Count int
}

func main() {
	// Define the file path
	filePath := "input/text.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		// If file doesn't exist in the current environment, 
		// we exit silently as per the requirement "Do not output anything else"
		// (assuming standard error handling behavior for missing files in a script context)
		return
	}
	defer file.Close()

	counts := make(map[string]int)

	// Regex to keep only letters
	// We define a regex that identifies anything that is NOT a letter
	re := regexp.MustCompile(`[^a-z]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert line to lowercase
		line := strings.ToLower(scanner.Text())

		// Replace all non-letter characters with spaces
		// This effectively strips punctuation and handles newlines/tabs
		cleanedLine := re.ReplaceAllString(line, " ")

		// Split by whitespace
		words := strings.Fields(cleanedLine)

		for _, word := range words {
			if word != "" {
				counts[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	// Convert map to a slice of structs for sorting
	var freqList []WordFreq
	for word, count := range counts {
		freqList = append(freqList, WordFreq{Word: word, Count: count})
	}

	// Sort by count descending; for ties, sort alphabetically ascending
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].Count != freqList[j].Count {
			return freqList[i].Count > freqList[j].Count
		}
		return freqList[i].Word < freqList[j].Word
	})

	// Output the results
	for _, entry := range freqList {
		fmt.Printf("%s: %d\n", entry.Word, entry.Count)
	}
}