package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// Record represents the structure of the JSON objects in the input array.
type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Define the path to the input file.
	filePath := "input/data.json"

	// Read the file content.
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// If the file doesn't exist or can't be read, we exit.
		// Using os.Exit(1) to ensure no extra output is printed to stdout.
		os.Exit(1)
	}

	var allRecords []Record
	// Parse the JSON array.
	err = json.Unmarshal(data, &allRecords)
	if err != nil {
		os.Exit(1)
	}

	// Filter records: active must be true AND age must be 30 or older.
	var filteredRecords []Record
	for _, r := range allRecords {
		if r.Active && r.Age >= 30 {
			filteredRecords = append(filteredRecords, r)
		}
	}

	// Sort the filtered records by name ascending.
	sort.Slice(filteredRecords, func(i, j int) bool {
		return filteredRecords[i].Name < filteredRecords[j].Name
	})

	// Encode the filtered and sorted list back to JSON.
	output, err := json.MarshalIndent(filteredRecords, "", "  ")
	if err != nil {
		os.Exit(1)
	}

	// Output the result to stdout.
	fmt.Println(string(output))
}