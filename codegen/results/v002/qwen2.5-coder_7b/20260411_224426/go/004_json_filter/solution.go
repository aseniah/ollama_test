package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// Open the file
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file contents into a byte slice
	byteValue, _ := ioutil.ReadAll(file)

	// Unmarshal the JSON data into an array of Person structs
	var people []Person
	json.Unmarshal(byteValue, &people)

	// Filter and sort the people based on the given conditions
	filteredPeople := make([]Person, 0)
	for _, person := range people {
		if person.Active && person.Age >= 30 {
			filteredPeople = append(filteredPeople, person)
		}
	}
	sort.Slice(filteredPeople, func(i, j int) bool {
		return filteredPeople[i].Name < filteredPeople[j].Name
	})

	// Marshal the filtered and sorted data back into JSON format
	jsonData, _ := json.MarshalIndent(filteredPeople, "", "  ")

	// Output the JSON data to stdout
	fmt.Println(string(jsonData))
}