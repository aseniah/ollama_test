package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// WordCount stores the word and its frequency
type WordCount struct {
	Word  string
	Count int
}

func main() {
	// Read the file input/text.txt
	content, err := os.ReadFile("input/text.txt")
	if err != nil {
		// If the file cannot be read, exit silently as per the requirement 
		// to not output anything else, though in a real scenario we would log this.
		return
	}

	// Convert content to string
	text := string(content)

	// Split the text into tokens by whitespace
	tokens := strings.Fields(text)

	// Map to store the frequency of each word
	counts := make(map[string]int)

	// Regex to identify everything that is NOT a letter
	// This allows us to strip all punctuation and digits
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return
	}

	for _, token := range tokens {
		// Convert to lowercase
		lowerToken := strings.ToLower(token)
		
		// Strip all non-letter characters
		cleanWord := reg.ReplaceAllString(lowerToken, "")

		// Only count if the resulting string is not empty
		if cleanWord != "" {
			counts[clean
				counts[cleanWord]++
			}
		}
	}

	// Convert map to a slice of WordCount structs for sorting
	var wordList []WordCount
	for word, count := range counts {
		wordList = append(wordList, WordCount{Word: word, Count: count})
	}

	// Sort the slice:
	// 1. By count descending
	// 2. By word ascending (alphabetical) for ties
	sort.Slice(wordList, func(i, j int) bool {
		if wordList[i].Count != wordList[j].Count {
			return wordList[i].Count > wordList[j].Count
		}
		return wordList[i].Word < wordList[j].Word
	})

	// Output the results in the format: word: count
	for _, wc := range wordList {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}