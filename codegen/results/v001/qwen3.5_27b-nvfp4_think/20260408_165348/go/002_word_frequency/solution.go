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
	splitter := regexp.MustCompile(`\s+`)
	stripper := regexp.MustCompile(`[^a-z]+`)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := splitter.Split(line, -1)

		for _, word := range words {
			cleanWord := stripper.ReplaceAllString(word, "")
			if len(cleanWord) > 0 {
				wordCount[cleanWord]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	type wordCountPair struct {
		word  string
		count int
	}

	pairs := make([]wordCountPair, 0, len(wordCount))
	for word, count := range wordCount {
		pairs = append(pairs, wordCountPair{word: word, count: count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	for _, pair := range pairs {
		fmt.Printf("%s: %d\n", pair.word, pair.count)
	}
}