package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Struct to hold the data
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Score   float64 `json:"score"`
}

func main() {
	// Read the CSV file
	file, err := os.Open("input/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read all data from the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Split the data into rows
	rows := strings.Split(string(data), "\n")
	rows = rows[1:] // Exclude the header row

	// Initialize the JSON array
	var persons []Person

	// Process each row to create Person objects
	for _, row := range rows {
		fields := strings.Split(row, ",")
		if len(fields) < 4 {
			continue // Skip malformed rows
		}

		name := fields[0]
		ageStr, _ := strconv.Atoi(fields[1])
		email := fields[2]
		scoreStr, _ := strconv.ParseFloat(fields[3], 64)

		person := Person{
			Name:    name,
			Age:     ageStr,
			Email:   email,
			Score:   scoreStr,
		}

		persons = append(persons, person)
	}

	// Write the JSON output
	err = json.Indent(os.Stdout, persons, "  ")
	if err != nil {
		log.Fatal(err)
	}
}