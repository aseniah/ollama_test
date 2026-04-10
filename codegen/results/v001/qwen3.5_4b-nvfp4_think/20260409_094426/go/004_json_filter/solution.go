package main

import (
	"encoding/json"
	"os"
	"sort"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	file, err := os.Open("input/data.json")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var records []Person
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&records); err != nil {
		os.Exit(1)
	}

	filtered := make([]Person, 0, len(records))
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.Marshal(filtered)
	if err != nil {
		os.Exit(1)
	}

	os.Stdout.Write(output)
}