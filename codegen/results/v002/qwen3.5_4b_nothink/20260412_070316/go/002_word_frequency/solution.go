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

		// Remove punctuation: keep only letters, digits, and spaces
		// Note: The rule says "keep only letters", but typically words are separated by spaces.
		// We'll split by space, then remove non-letter characters from each word.
		pattern := regexp.MustCompile(`[^a-z]`)
		words := pattern.Split(line, -1)

		for _, word := range words {
			if word != "" {
				wordCount[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create a slice of words for sorting
	type wordCount struct {
		Word   string
		Count  int
	}

	var wordList []wordCount
	for word, count := range wordCount {
		wordList = append(wordList, wordCount{Word: word, Count: count})
	}

	// Sort by count descending, then by word alphabetically ascending
	sort.Slice(wordList, func(i, j int) bool {
		if wordList[i].Count != wordList[j].Count {
			return wordList[i].Count > wordList[j].Count
		}
		return wordList[i].Word < wordList[j].Word
	})

	// Output
	for _, item := range wordList {
		fmt.Printf("%s: %d\n", item.Word, item.Count)
	}
}