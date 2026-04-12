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
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCounts := make(map[string]int)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		pattern := regexp.MustCompile(`[^a-z0-9]`)
		words := pattern.Split(strings.ToLower(string(line)), -1)

		for _, w := range words {
			w = strings.TrimSpace(w)
			if len(w) > 0 {
				wordCounts[w]++
			}
		}
	}

	type WordCount struct {
		Word  string
	 Count int
	}

	var wcList []WordCount
	for word, count := range wordCounts {
		wcList = append(wcList, WordCount{Word: word, Count: count})
	}

	sort.Slice(wcList, func(i, j int) bool {
		if wcList[i].Count != wcList[j].Count {
			return wcList[i].Count > wcList[j].Count
		}
		return wcList[i].Word < wcList[j].Word
	})

	for _, w := range wcList {
		fmt.Printf("%s: %d\n", w.Word, w.Count)
	}
}