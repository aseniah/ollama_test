package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// Record represents the structure of the objects in the input JSON array.
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Define the path to the input file
	inputPath := "input/data.json"

	// Read the file contents
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		// If the file cannot be read, we exit silently to avoid 
		// outputting anything else to stdout per requirements.
		os.Exit(1)
	}

	var records []Record
	// Unmarshal the JSON array into the slice of Record structs
	if err := json.Unmarshal(data, &records); err != nil {
		os.Exit(1)
	}

	// Filter the records: active must be true and age must be >= 30
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

	// Marshal the filtered and sorted slice back to JSON
	output, err := json.Marshal(filtered)
	if err != nil {
		os.Exit(1)
	}

	// Output the resulting JSON array to stdout
	fmt.Print(string(output))
}