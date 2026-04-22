package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Fields(scanner.Text()) {
			var sb strings.Builder
			for _, r := range word {
				if unicode.IsLetter(r) {
					sb.WriteRune(unicode.ToLower(r))
				}
			}
			w := sb.String()
			if w != "" {
				counts[w]++
			}
		}
	}

	type pair struct {
		word  string
		count int
	}
	pairs := make([]pair, 0, len(counts))
	for w, c := range counts {
		pairs = append(pairs, pair{w, c})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.word, p.count)
	}
}