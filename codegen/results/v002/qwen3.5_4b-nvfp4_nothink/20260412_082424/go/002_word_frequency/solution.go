package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var words []string
	scanner := bufio.Scanner{Scopes: 1000}
	for scanner.Scan() {
		line := scanner.Text()
		wordRegexp := regexp.MustCompile(`[a-z]+`)
		matches := wordRegexp.FindAllString(line, -1)
		for _, w := range matches {
			words = append(words, w)
		}
		if err := scanner.Err(); err != nil {
			break
		}
	}

	if err := scanner.Err(); err != nil && file == nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	wordCounts := make(map[string]int)
	for _, w := range words {
		wordCounts[w]++
	}

	type WordCount struct {
		Word   string
		Count  int
		Sorted string
	}

	var results []WordCount
	for word, count := range wordCounts {
		results = append(results, WordCount{Word: word, Count: count})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Count != results[j].Count {
			return results[i].Count > results[j].Count
		}
		return results[i].Word < results[j].Word
	})

	for _, wc := range results {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}