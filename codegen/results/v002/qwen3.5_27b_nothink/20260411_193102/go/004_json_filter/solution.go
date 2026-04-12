package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Read the file
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal JSON into a slice of Records
	var allRecords []Record
	if err := json.Unmarshal(data, &allRecords); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Filter records where active is true and age is >= 30
	var filtered []Record
	for _, r := range allRecords {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Marshal the result back to JSON
	output, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Output to stdout
	fmt.Println(string(output))
}