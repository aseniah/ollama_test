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

	// Map to store word counts
	words := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		
		// Remove non-alphabetic characters using regex
		// This keeps only letters (a-z) and spaces
		re := regexp.MustCompile("[^a-z ]")
		cleanedLine := re.ReplaceAllString(line, " ")
		
		// Split into words
		wordsInLine := strings.Fields(cleanedLine)
		
		for _, word := range wordsInLine {
			words[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Convert map to slice for sorting
	type wordCount struct {
		word  string
		count int
	}

	var wordCountList []wordCount
	for word, count := range words {
		wordCountList = append(wordCountList, wordCount{word: word, count: count})
	}

	// Sort by count descending, then by word alphabetically ascending
	sort.Slice(wordCountList, func(i, j int) bool {
		if wordCountList[i].count != wordCountList[j].count {
			return wordCountList[i].count > wordCountList[j].count
		}
		return wordCountList[i].word < wordCountList[j].word
	})

	// Output results
	for _, wc := range wordCountList {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}