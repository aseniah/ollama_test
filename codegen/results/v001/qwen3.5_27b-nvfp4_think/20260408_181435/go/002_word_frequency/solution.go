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
	letterPattern := regexp.MustCompile(`[a-zA-Z]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := letterPattern.FindAllString(line, -1)
		for _, word := range words {
			word = strings.ToLower(word)
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	type wordFreq struct {
		word  string
		count int
	}

	var results []wordFreq
	for word, count := range wordCount {
		results = append(results, wordFreq{word: word, count: count})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].count != results[j].count {
			return results[i].count > results[j].count // Descending by count
		}
		return results[i].word < results[j].word // Ascending by word for ties
	})

	for _, r := range results {
		fmt.Printf("%s: %d\n", r.word, r.count)
	}
}