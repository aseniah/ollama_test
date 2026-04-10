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

	var counts map[string]int = make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
		matches := wordRegex.FindAllString(line, -1)

		for _, word := range matches {
			normalizedWord := strings.ToLower(word)
			counts[normalizedWord]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	type wordCount struct {
		Word string
		Count int
	}

	var wordCounts []wordCount
	for word, count := range counts {
		wordCounts = append(wordCounts, wordCount{Word: word, Count: count})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].Count != wordCounts[j].Count {
			return wordCounts[i].Count > wordCounts[j].Count
		}
		return wordCounts[i].Word < wordCounts[j].Word
	})

	for _, wc := range wordCounts {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}