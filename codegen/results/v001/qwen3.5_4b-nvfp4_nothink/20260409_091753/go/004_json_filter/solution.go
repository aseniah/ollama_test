package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// Read the input file
	data, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	var rawRecords json.RawMessage
	if len(data) > 0 {
		_ = json.Unmarshal(data, &rawRecords) // Just to validate it's JSON if needed, but we trust the format based on prompt description below
	}

	var records []Record
	err = json.Unmarshal(data, &records)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error unmarshaling JSON:", err)
		os.Exit(1)
	}

	// Filter records: active is true and age is 30 or older
	var filteredRecords []Record
	for _, record := range records {
		if record.Active && record.Age >= 30 {
			filteredRecords = append(filteredRecords, record)
		}
	}

	// Sort by name ascending
	sort.Slice(filteredRecords, func(i, j int) bool {
		return filteredRecords[i].Name < filteredRecords[j].Name
	})

	// Output the result as JSON array
	output, err := json.Marshal(filteredRecords)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling output:", err)
		os.Exit(1)
	}

	_, err = fmt.Dprintf("%s\n", output)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing to stdout:", err)
		os.Exit(1)
	}
}