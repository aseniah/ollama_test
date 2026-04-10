package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filePath := "input/text.txt"
	wordCount := make(map[string]int)

	// Read the file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Process each line
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			if _, exists := wordCount[lowerWord]; !exists {
				wordCount[lowerWord] = 0
			}
			wordCount[lowerWord]++
		}
	}

	// Sort words by count descending, then alphabetically ascending
	sortedWords := make([]string, 0, len(wordCount))
	for word := range wordCount {
		sortedWords = append(sortedWords, word)
	}
	sortedWords = sortWordsByCount(sortedWords, wordCount)

	// Output the results
	for _, word := range sortedWords {
		fmt.Printf("%s: %d\n", word, wordCount[word])
	}
}

// Helper function to sort words by count descending, then alphabetically ascending
func sortWordsByCount(words []string, wordCount map[string]int) []string {
	sorted := make([]struct{ word string, count int }{
		{word: word, count: wordCount[word]},
		{word: word, count: wordCount[word]},
		{word: word, count: wordCount[word]},
		// Add more words as needed
	})

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].count == sorted[j].count {
			return sorted[i].word < sorted[j].word
		}
		return sorted[i].count > sorted[j].count
	})

	return make([]string, 0, len(sorted))
}