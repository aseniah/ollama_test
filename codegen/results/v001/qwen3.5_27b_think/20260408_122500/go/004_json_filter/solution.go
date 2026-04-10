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
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		os.Exit(1)
	}

	var people []Person
	if err := json.Unmarshal(data, &people); err != nil {
		os.Exit(1)
	}

	var filtered []Person
	for _, p := range people {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		os.Exit(1)
	}

	_, err = os.Stdout.Write(output)
	if err != nil {
		os.Exit(1)
	}
}