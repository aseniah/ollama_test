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
	// Define the path to the input file
	inputPath := "input/data.json"

	// Read the file content
	data, err := os.ReadFile(inputPath)
	if err != nil {
		// If the file doesn't exist or can't be read, we exit.
		// In a production environment, you might want to handle this differently.
		return
	}

	// Unmarshal the JSON array into a slice of Record
	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		return
	}

	// Filter the records: active must be true and age must be 30 or older
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort the filtered records by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Marshal the filtered and sorted records back to JSON
	// Using MarshalIndent for readable output, but standard Marshal works too.
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		return
	}

	// Output the resulting JSON to stdout
	fmt.Println(string(output))
}