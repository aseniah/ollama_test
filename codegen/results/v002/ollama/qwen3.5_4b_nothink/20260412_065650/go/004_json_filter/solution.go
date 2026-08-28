package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Record struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Active bool   `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	// Read input file
	data, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		log.Fatal("Error reading input file:", err)
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	// Filter records: active is true and age >= 30
	var filteredRecords []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filteredRecords = append(filteredRecords, r)
		}
	}

	// Sort by name ascending
	sort.Slice(filteredRecords, func(i, j int) bool {
		return filteredRecords[i].Name < filteredRecords[j].Name
	})

	// Output to stdout as JSON
	output, err := json.Marshal(filteredRecords)
	if err != nil {
		log.Fatal("Error marshaling output:", err)
	}

	fmt.Println(string(output))
}