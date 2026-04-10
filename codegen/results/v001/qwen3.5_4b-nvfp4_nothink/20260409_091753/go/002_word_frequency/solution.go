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
	// Open the input file
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Regex to match alphanumeric characters and remove punctuation
	lowercaseWordRe := regexp.MustCompile(`[a-z]+`)
	wordCounter := make(map[string]int)

	// Read and process the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		// Replace all non-alphabetic characters with spaces, then split into words
		text = lowercaseWordRe.ReplaceAllString(text, " ")
		words := strings.Fields(text)

		for _, word := range words {
			wordCounter[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	// Convert to a slice for sorting
	type wordCount struct {
		word  string
		count int
	}

	var wordsToSort []wordCount
	for word, count := range wordCounter {
		wordsToSort = append(wordsToSort, wordCount{word: word, count: count})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(wordsToSort, func(i, j int) bool {
		if wordsToSort[i].count != wordsToSort[j].count {
			return wordsToSort[i].count > wordsToSort[j].count
		}
		return wordsToSort[i].word < wordsToSort[j].word
	})

	// Output results
	for _, wc := range wordsToSort {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}