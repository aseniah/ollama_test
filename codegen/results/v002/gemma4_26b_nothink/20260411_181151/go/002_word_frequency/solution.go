package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// wordFreq holds the word and its associated count for sorting
type wordFreq struct {
	word  string
	count int
}

func main() {
	// Path to the input file
	filePath := "input/text.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		// If the file doesn't exist in the environment, we exit silently 
		// as per the requirement to not output anything else.
		return
	}
	defer file.Close()

	// Map to store frequencies
	counts := make(map[string]int)

	// Regex to keep only letters (removes punctuation and digits)
	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert line to lowercase
		line := strings.ToLower(scanner.Text())

		// Replace everything that isn't a letter with a space
		// This ensures words separated by punctuation are split correctly
		cleanedLine := reg.ReplaceAllString(line, " ")

		// Split by whitespace
		words := strings.Fields(cleanedLine)

		for _, w := range words {
			if w != "" {
				counts[w]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	// Transfer map to a slice for sorting
	var freqList []wordFreq
	for word, count := range counts {
		freqList = append(freqList, wordFreq{word, count})
	}

	// Sort rules:
	// 1. Count descending
	// 2. Alphabetical ascending for ties
	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].count != freqList[j].count {
			return freqList[i].count > freqList[j].count
		}
		return freqList[i].word < freqList[j].word
	})

	// Print output
	for _, item := range freqList {
		fmt.Printf("%s: %d\n", item.word, item.count)
	}
}