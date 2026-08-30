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
		return
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	// Regex to match sequences of letters
	re := regexp.MustCompile(`[a-zA-Z]+`)

	for scanner.Scan() {
		line := scanner.Text()
		words := re.FindAllString(line, -1)
		for _, w := range words {
			lower := strings.ToLower(w)
			wordCounts[lower]++
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	// Create a slice of (word, count) pairs
	type wordCount struct {
		word  string
		count int
	}
	var wcList []wordCount
	for word, count := range wordCounts {
		wcList = append(wcList, wordCount{word, count})
	}

	// Sort by count descending, then word ascending
	sort.Slice(wcList, func(i, j int) bool {
		if wcList[i].count == wcList[j].count {
			return wcList[i].word < wcList[j].word
		}
		return wcList[i].count > wcList[j].count
	})

	for _, wc := range wcList {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}