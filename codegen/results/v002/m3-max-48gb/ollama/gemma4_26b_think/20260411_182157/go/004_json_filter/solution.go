package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Record represents the structure of each object in the input JSON array.
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Define the path to the input file.
	inputPath := "input/data.json"

	// Read the entire content of the file.
	data, err := os.ReadFile(inputPath)
	if err != nil {
		// If the file cannot be read, we exit without printing anything to stdout
		// to comply with the "Do not output anything else" requirement.
		return
	}

	// Unmarshal the JSON data into a slice of Record objects.
	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		return
	}

	// Create a slice to hold the filtered records.
	var filtered []Record

	// Filter records: active must be true AND age must be 30 or older.
	for _, record := range records {
		if record.Active && record.Age >= 30 {
			filtered = append(filtered, record)
		}
	}

	// Sort the filtered records by name in ascending order.
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Marshal the filtered and sorted records back into JSON.
	// Using MarshalIndent to produce a human-readable JSON array.
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		return
	}

	// Output the resulting JSON array to stdout.
	fmt.Println(string(output))
}