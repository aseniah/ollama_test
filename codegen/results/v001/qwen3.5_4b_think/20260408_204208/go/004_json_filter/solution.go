package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Record struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Active bool   `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	// Open the input file
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the entire file content
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	var records []Record
	err = json.Unmarshal(data, &records)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error unmarshaling JSON:", err)
		os.Exit(1)
	}

	// Filter records where active is true and age >= 30
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Output to stdout as JSON
	output, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Print(string(output))
}