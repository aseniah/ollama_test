package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Read the JSON file
	data, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the JSON data into a slice of Person structs
	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Filter the people based on the conditions
	var filteredPeople []Person
	for _, person := range people {
		if person.Active && person.Age >= 30 {
			filteredPeople = append(filteredPeople, person)
		}
	}

	// Sort the filtered people by name
	sort.Slice(filteredPeople, func(i, j int) bool {
		return filteredPeople[i].Name < filteredPeople[j].Name
	})

	// Output the result as JSON
	jsonData, err := json.MarshalIndent(filteredPeople, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}