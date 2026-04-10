package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Person represents the structure of each row in the CSV
type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Open the input file
	file, err := os.Open("input/data.csv")
	if err != nil {
		// If file doesn't exist, we exit silently as per "Do not output anything else" 
		// constraint for valid JSON output requirement, but in a real tool 
		// you'd handle this error.
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row
	_, err = reader.Read()
	if err != nil {
		// If file is empty or error reading header, output empty array
		fmt.Print("[]")
		return
	}

	var people []Person

	// Iterate through the remaining rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		// Ensure we have enough columns to avoid index out of bounds
		if len(record) < 4 {
			continue
		}

		// Parse Age (integer)
		age, err := strconv.Atoi(record[1])
		if err != nil {
			age = 0
		}

		// Parse Score (float)
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			score = 0.0
		}

		person := Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		}

		people = append(people, person)
	}

	// Encode the slice to JSON and write to stdout
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(people)
	if err != nil {
		return
	}
}