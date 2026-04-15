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
		// If file doesn't exist, exit silently as per instructions to output nothing else
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(file)

	// Regex to match words containing only letters (a-z, A-Z)
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)

	for scanner.Scan() {
		line := scanner.Text()
		words := wordRegex.FindAllString(line, -1)
		for _, w := range words {
			lower := strings.ToLower(w)
			wordCount[lower]++
		}
	}

	// Create a slice of unique words
	words := make([]string, 0, len(wordCount))
	for word := range wordCount {
		words = append(words, word)
	}

	// Sort words by count descending, then alphabetically ascending
	sort.Slice(words, func(i, j int) bool {
		countI := wordCount[words[i]]
		countJ := wordCount[words[j]]
		if countI != countJ {
			return countI > countJ
		}
		return words[i] < words[j]
	})

	// Output results
	for _, word := range words {
		fmt.Printf("%s: %d\n", word, wordCount[word])
	}
}