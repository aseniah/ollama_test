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

	// Regex to extract words (sequences of letters)
	re := regexp.MustCompile(`[a-zA-Z]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		// Extract words using regex
		words := re.FindAllString(line, -1)
		for _, word := range words {
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create a slice of structs to hold word and count for sorting
	type WordCount struct {
		Word  string
		Count int
	}

	var wordList []WordCount
	for word, count := range wordCount {
		wordList = append(wordList, WordCount{Word: word, Count: count})
	}

	// Sort by count descending, then by word alphabetically ascending
	sort.Slice(wordList, func(i, j int) bool {
		if wordList[i].Count != wordList[j].Count {
			return wordList[i].Count > wordList[j].Count
		}
		return wordList[i].Word < wordList[j].Word
	})

	// Output results
	for _, wc := range wordList {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}