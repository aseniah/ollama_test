package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	inputFile := flag.String("i", "input/text.txt", "Input file path")
	outputFile := flag.String("o", "output/text.txt", "Output file path")
	flag.Parse()

	// Regex to extract words (sequences of letters)
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)

	// Load input file
	words, err := loadFile(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading input file: %v\n", err)
		os.Exit(1)
	}

	// Count word frequencies
	wordCounts := make(map[string]int)
	for _, w := range words {
		// Convert to lowercase and extract only letters
		processed := wordRegex.FindAllString(w, -1)
		for _, word := range processed {
			wordCounts[strings.ToLower(word)]++
		}
	}

	// Prepare output slice
	type WordCount struct {
		Word  string
		Count int
	}
	type SortWord struct {
		Cnt       int
		Word      string
		Original  string
	}
	sorter := make([]SortWord, 0, len(wordCounts))
	for word, count := range wordCounts {
		sortedWords := strings.Fields(word)
		if len(sortedWords) == 0 {
			continue
		}
		// Find the first word in the split
		originalWord := strings.ToLower(sortedWords[0])
		sortedWords = strings.Split(originalWord, " ")
		sorter = append(sortedWord, wordCounts, sortedWord, originalWord)
	}

	// Sort the words by count descending, then by word alphabetically ascending
	sort.Slice(sorter, func(i, j int) bool {
		if sorter[i].Count != sorter[j].Count {
			return sorter[i].Count > sorter[j].Count
		}
		return sorter[i].Word < sorter[j].Word
	})

	// Output the results
	for _, w := range sorter {
		fmt.Printf("%s: %d\n", w.Word, w.Count)
	}
}

func loadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Extract words from the line
		wordRegex := regexp.MustCompile(`\w+`)
		words = append(words, wordRegex.FindAllString(line, -1)...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}