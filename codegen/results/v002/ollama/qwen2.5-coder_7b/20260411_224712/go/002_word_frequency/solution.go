package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a map to store word counts
	wordCounts := make(map[string]int)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.FieldsFunc(line, func(r rune) bool {
			return !strings.IsLetter(string(r))
		})

		// Convert words to lowercase and count them
		for _, word := range words {
			word = strings.ToLower(word)
			wordCounts[word]++
		}
	}

	// Sort the words by frequency descending and alphabetically ascending for ties
	var sortedWords []string
	for word := range wordCounts {
		sortedWords = append(sortedWords, word)
	}
	sort.SliceStable(sortedWords, func(i, j int) bool {
		if wordCounts[sortedWords[i]] == wordCounts[sortedWords[j]] {
			return sortedWords[i] < sortedWords[j]
		}
		return wordCounts[sortedWords[i]] > wordCounts[sortedWords[j]]
	})

	// Output the sorted words and counts
	for _, word := range sortedWords {
		fmt.Println(word + ": " + strconv.Itoa(wordCounts[word]))
	}

	// Check for errors reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}