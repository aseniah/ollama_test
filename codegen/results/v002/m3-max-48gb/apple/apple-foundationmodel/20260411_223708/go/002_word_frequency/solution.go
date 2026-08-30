package main

import (
	"bufio"
	"fmt"
	"os"
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

	// Initialize a map to store word counts
	wordCount := make(map[string]int)

	// Read from the file
	reader := bufio.NewReader(file)
	for {
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		// Split the line into words and process each word
		words := strings.Fields(line)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			// Strip punctuation and count the word
			if _, _ := strings.Trim(lowerWord, ".,!?;:\""); wordCount[lowerWord] == 0 {
				wordCount[lowerWord] = 1
			} else {
				wordCount[lowerWord]++
			}
		}
	}

	// Sort the words by count descending, then alphabetically ascending
	sortedWords := make([]string, 0, len(wordCount))
	for word, count := range wordCount {
		sortedWords = append(sortedWords, fmt.Sprintf("%s: %d", word, count))
	}

	// Sort by count descending and alphabetically ascending for ties
	sortedWords = sortWordsByCountDescendingAndAscending(sortedWords)

	// Output the sorted words
	for _, line := range sortedWords {
		fmt.Println(line)
	}
}

// Helper function to sort words by count descending and alphabetically ascending
func sortWordsByCountDescendingAndAscending(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		countI, _ := strconv.Atoi(words[i])
		countJ, _ := strconv.Atoi(words[j])
		if countI != countJ {
			return countI > countJ
		} else {
			return words[i] < words[j]
		}
	})
	return words
}