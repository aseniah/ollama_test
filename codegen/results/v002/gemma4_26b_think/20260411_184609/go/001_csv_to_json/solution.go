package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Person represents the structure of each row in the CSV file as defined in the requirements.
type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Define the path to the input file.
	filePath := "input/data.csv"

	// Open the input file.
	file, err := os/Open(filePath)
	if err != nil {
		// If the file cannot be opened, we exit. 
		// We do not print to stdout to ensure we don't pollute the required JSON output.
		os.Exit(1)
	}
	defer file.Close()

	// Initialize a CSV reader.
	reader := csv.NewReader(file)

	// Read all records from the CSV.
	records, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}

	// We initialize the slice with an empty slice to ensure that if the CSV 
	// only contains a header, the output is an empty array [] instead of null.
	people := make([]Person, 0)

	// Iterate through the records, skipping the first row (the header).
	for i := 1; i < len(records); i++ {
		row := records[i]

		// Ensure the row has the expected number of columns.
		if len(row) < 4 {
			continue
		}

		// Extract and parse Name (string).
		name := row[0]

		// Extract and parse Age (int).
		age, err := strconv.Atoi(row[1])
		if err != nil {
			continue
		}

		// Extract Email (string).
		email := row[2]

		// Extract and parse Score (float64).
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			continue
		}

		// Append the parsed data to the people slice.
		people = append(people, Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		})
	}

	// Marshal the slice of Person objects into a JSON array.
	jsonData, err := json.Marshal(people)
	if err != nil {
		os.Exit(1)
	}

	// Print the resulting JSON string to stdout.
	fmt.Println(string(jsonData))
}