package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Record represents the structure of the objects in the JSON array.
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Path to the input file
	filePath := "input/data.json"

	// Read the file contents
	data, err := os.ReadFile(filePath)
	if err != nil {
		// If the file cannot be read, we exit without printing to stdout 
		// to adhere to the "Do not output anything else" constraint.
		os.Exit(1)
	}

	// Parse the JSON array
	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		os.Exit(1)
	}

	// Filter records: active must be true and age must be 30 or older
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort the filtered records by name in ascending order
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Encode the resulting slice back to JSON
	// We use MarshalIndent for clean, readable output, but standard Marshal works too.
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		os.Exit(1)
	}

	// Write the JSON array to stdout
	fmt.Println(string(output))
}