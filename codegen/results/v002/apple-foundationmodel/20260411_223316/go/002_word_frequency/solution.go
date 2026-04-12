package main

import (
	"fmt"
	"strings"
)

func main() {
	// Read the file
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var wordCount map[string]int
	var inputChars = []byte{}

	for {
		ch := <-file.ReadChan()
		inputChars = append(inputChars, ch)
		if ch == 0 {
			break
		}
	}

	// Convert to lowercase and remove punctuation
	words := strings.Fields(string(inputChars))
	wordCount = make(map[string]int)

	for _, word := range words {
		// Remove punctuation
		word = strings.Replace(word, "[^a-zA-Z]", "", -1)
		// Convert to lowercase
		word = strings.ToLower(word)

		// Increment count
		wordCount[word]++
	}

	// Sort words by count descending and alphabetically ascending for ties
	sortedWords := make([]string, 0, len(wordCount))
	for word := range wordCount {
		sortedWords = append(sortedWords, word)
	}

	sortedWords = append(sortedWords, strings.Fields(strings.Join(sortedWords, " ")))
	sortedWords = strings.Fields(strings.Join(sortedWords, " "))

	// Output the result
	for _, word := range sortedWords {
		fmt.Printf("%s:%d\n", word, wordCount[word])
	}
}