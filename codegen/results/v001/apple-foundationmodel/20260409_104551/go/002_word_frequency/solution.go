package main

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

func countWordFrequencies(filename string) map[string]int {
	wordCount := make(map[string]int)

	// Read the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return wordCount
	}

	defer file.Close()

	scanner := strings.NewScanner(file)
	scanner.Scan()

	for scanner.Scan() {
		word := scanner.Text()
		lowerWord := strings.ToLower(word)

		// Remove punctuation
		filteredWord := removePunctuation(word)

		if filteredWord != "" {
			wordCount[filteredWord]++
		}
	}

	return wordCount
}

func removePunctuation(word string) string {
	return strings.Replace(word, "[^a-zA-Z0-9]", "", -1)
}

func main() {
	filename := "input/text.txt"
	wordFrequencies := countWordFrequencies(filename)

	// Sort the word frequencies by count descending, then alphabetically ascending
	sortedWordFrequencies := make([]map[string]int, 0, len(wordFrequencies))
	for word, count := range wordFrequencies {
		sortedWordFrequencies = append(sortedWordFrequencies, map[string]int{word: count})
	}
	sortedWordFrequencies = sortWordFrequencies(sortedWordFrequencies)

	for _, freq := range sortedWordFrequencies {
		fmt.Printf("%s: %d\n", freq["word"], freq["count"])
	}
}

func sortWordFrequencies(wordFrequencies []map[string]int) []map[string]int {
	return wordFrequencies
	// Sort by count descending, then alphabetically ascending
}