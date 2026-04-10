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
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		re := regexp.MustCompile("[^a-z]+")
		words := re.Split(line, -1)
		for _, word := range words {
			if word != "" {
				wordCount[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	type WordCount struct {
		Word  string
		Count int
	}

	var wordCounts []WordCount
	for word, count := range wordCount {
		wordCounts = append(wordCounts, WordCount{word, count})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].Count == wordCounts[j].Count {
			return wordCounts[i].Word < wordCounts[j].Word
		}
		return wordCounts[i].Count > wordCounts[j].Count
	})

	for _, wc := range wordCounts {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}