package main

import (
    "fmt"
    "strings"
)

func main() {
    file, err := os.Open("input/text.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }

    wordCount := make(map[string]int)

    for _, line := range file.ReadLines() {
        lineText := strings.TrimSpace(strings.ToLower(line.Text()))
        words := strings.Split(lineText, " ")

        for _, word := range words {
            word = strings.Replace(word, "[^a-zA-Z]", "", -1) // Remove punctuation
            if _, exists := wordCount[word]; !exists {
                wordCount[word] = 0
            }
            wordCount[word]++
        }
    }

    sortedWords := make([]string, 0, len(wordCount))
    for word, count := range wordCount {
        sortedWords = append(sortedWords, []string{word, count})
    }

    sortedWords = sortWordsByCountAndName(sortedWords)

    for _, entry := range sortedWords {
        fmt.Printf("%s: %d\n", entry[0], entry[1])
    }
}

func sortWordsByCountAndName(words []string) []string {
    for i := 0; i < len(words)-1; i++ {
        for j := i+1; j < len(words); j++ {
            if (words[i][1] < words[j][1]) ||
               (words[i][1] == words[j][1] && words[i][0] < words[j][0]) {
                words[i], words[j] = words[j], words[i]
            }
        }
    }
    return words
}