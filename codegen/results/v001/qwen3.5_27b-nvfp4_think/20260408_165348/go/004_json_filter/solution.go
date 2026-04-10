package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		os.Exit(1)
	}

	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		os.Exit(1)
	}

	filtered := []Person{}
	for _, person := range people {
		if person.Active && person.Age >= 30 {
			filtered = append(filtered, person)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(output))
}