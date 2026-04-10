package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// main function to execute the program
func main() {
	// Read the file
	file, err := os.Open("input/text.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()

	// Create a map to store word frequencies
	wordCount := make(map[string]int)

	// Read the file content
	reader := bufio.NewReader(file)
	for {
		// Read a line from the file
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		// If there's an error, log it and exit
		if err != nil {
			log.Printf("Error reading line: %v", err)
			return
		}

		// Trim whitespace and convert to lowercase
		line = strings.TrimSpace(line)
		line = strings.ToLower(line)

		// Split the line into words
		words := strings.Fields(line)

		// Process each word
		for _, word := range words {
			wordCount[word]++
		}
	}

	// Prepare the output
	sortedWords := make([]string, 0, len(wordCount))
	for word, count := range wordCount {
		sortedWords = append(sortedWords, fmt.Sprintf("%s:%d", word, count))
	}

	// Sort the words by count descending, then alphabetically ascending
	sort.Slice(sortedWords, func(i, j int) bool {
		if wordCount[sortedWords[i]] == wordCount[sortedWords[j]] {
			return sortedWords[i] < sortedWords[j]
		}
		return wordCount[sortedWords[i]] > wordCount[sortedWords[j]]
	})

	// Print each word and its count
	for _, word := range sortedWords {
		fmt.Println(word)
	}
}