package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Record struct {
	Name  string      `json:"name"`
	Age   int         `json:"age"`
	Active bool       `json:"active"`
	Score float64     `json:"score"`
}

func main() {
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var records []Record
	if err := json.Unmarshal(content, &records); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	filtered := make([]Record, 0, len(records))
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}