package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Active bool   `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	file, err := os.Open("input/data.json")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var records []Person
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&records); err != nil {
		os.Exit(1)
	}

	var filtered []Person
	for _, p := range records {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.Marshal(filtered)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print(string(output))
}