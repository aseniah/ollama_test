package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

// Person represents the structure of each row in the CSV file.
type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Open the specific input file
	file, err := os.Open("input/data.csv")
	if err != nil {
		// If the file cannot be opened, we exit silently to avoid 
		// polluting stdout with non-JSON error messages.
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	var people []Person

	// If the file only contains a header or is empty, 
	// the slice remains an empty slice (which encodes to []).
	if len(records) > 1 {
		// Iterate through records starting from index 1 to skip the header row
		for _, record := range records[1:] {
			// Ensure the row has the required number of columns to prevent index out of bounds
			if len(record) < 4 {
				continue
			}

			// Parse Age (index 1) as an integer
			age, _ := strconv.Atoi(record[1])

			// Parse Score (index 3) as a float64
			score, _ := strconv.ParseFloat(record[3], 64)

			person := Person{
				Name:  record[0],
				Age:   age,
				Email: record[2],
				Score: score,
			}
			people = append(people, person)
		}
	} else {
		// Initialize as an empty slice to ensure output is [] instead of null
		people = []Person{}
	}

	// Encode the slice of persons directly to stdout as a JSON array
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(people)
	if err != nil {
		return
	}
}