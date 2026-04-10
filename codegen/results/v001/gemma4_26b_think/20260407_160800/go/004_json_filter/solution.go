package main

import (
	"encoding/json"
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
	// Open the input file.
	file, err := os.Open("input/data.json")
	if err != nil {
		return
	}
	defer file.Close()

	// Decode the JSON array from the file.
	var records []Record
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&records); err != nil {
		return
	}

	// Filter records: active must be true and age must be 30 or older.
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort the filtered records by name in ascending order.
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Output the resulting JSON array to stdout.
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(filtered); err != nil {
		return
	}
}