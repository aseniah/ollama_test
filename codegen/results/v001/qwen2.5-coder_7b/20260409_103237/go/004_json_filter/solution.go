package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// Person represents an individual record in the JSON array.
type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Active  bool    `json:"active"`
	Score   float64 `json:"score"`
}

func main() {
	// Open and read the input file
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Filter and sort the data
	filtered := make([]Person, 0)
	for _, person := range people {
		if person.Active && person.Age >= 30 {
			filtered = append(filtered, person)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Output the result to stdout as a JSON array
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(output))
}