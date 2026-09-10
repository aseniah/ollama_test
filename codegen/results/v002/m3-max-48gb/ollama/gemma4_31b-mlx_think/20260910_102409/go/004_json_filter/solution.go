package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Record represents the structure of the objects in the input JSON
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Read the input file
	filePath := "input/data.json"
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	var records []Record
	// Unmarshal the JSON array into a slice of Record structs
	if err := json.Unmarshal(fileBytes, &records); err != nil {
		return
	}

	// Filter records: active must be true and age must be 30 or older
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort the records by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Marshal the filtered and sorted list back to JSON
	// Use MarshalIndent for human-readable output or Marshal for compact output
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		return
	}

	// Output the resulting JSON array to stdout
	fmt.Println(string(output))
}