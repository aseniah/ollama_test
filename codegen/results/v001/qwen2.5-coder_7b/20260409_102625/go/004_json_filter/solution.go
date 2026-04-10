package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"os"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Open the input file
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the entire file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal JSON into a slice of Person structs
	var people []Person
	err = json.Unmarshal(content, &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Filter and sort the people
	filtered := make([]Person, 0)
	for _, person := range people {
		if person.Active && person.Age >= 30 {
			filtered = append(filtered, person)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Marshal the filtered and sorted slice back to JSON
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the output to stdout
	fmt.Println(string(output))
}