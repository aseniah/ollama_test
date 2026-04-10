package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Person represents the transformed structure for the output JSON.
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// The target date for age calculation: July 1, 2025.
	targetDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	// Open the input CSV file.
	file, err := os.Open("input/input.csv")
	if err != nil {
		// If the file cannot be opened, output an empty JSON array.
		fmt.Print("[]")
		return
	}
	defer file.Close()

	// Initialize the CSV reader.
	reader := csv.NewReader(file)

	// Read all records from the CSV.
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Print("[]")
		return
	}

	// If there's no data (only header or empty), output an empty JSON array.
	if len(records) <= 1 {
		fmt.Print("[]")
		return
	}

	// Identify the column indices for name and date of birth (dob).
	// We check common naming conventions.
	header := records[0]
	nameIdx, dobIdx := -1, -1
	for i, col := range header {
		switch col {
		case "name", "Name", "full_name", "FullName":
			nameIdx = i
		case "dob", "birth_date", "date_of_birth", "birthdate", "BirthDate":
			dobIdx = i
		}
	}

	// Fallback to default indices (0 and 1) if headers are not specifically identified.
	if nameIdx == -1 {
		nameIdx = 0
	}
	if dobIdx == -1 {
		dobIdx = 1
	}

	var output []Person

	// Iterate through the CSV rows (skipping the header row).
	for i := 1; i < len(records); i++ {
		row := records[i]

		// Ensure the row has enough columns to access the required indices.
		if len(row) <= nameIdx || len(row) <= dobIdx {
			continue
		}

		name := row[nameIdx]
		dobStr := row[dobIdx]

		// Attempt to parse the date string using common ISO and standard formats.
		dateFormats := []string{
			"2006-01-02",
			"01/02/2006",
			"2006/01/02",
			"1/2/2006",
			"2/1/2006",
		}

		var dob time.Time
		var parseErr error
		for _, fmtStr := range dateFormats {
			dob, parseErr = time.Parse(fmtStr, dobStr)
			if parseErr == nil {
				break
			}
		}

		// If parsing the date fails, skip this record.
		if parseErr != nil {
			continue
		}

		// Calculate the age as of July 1, 2025.
		// Age = Difference in years, adjusted if the birthday hasn't occurred yet in the target year.
		age := targetDate.Year() - dob.Year()
		if targetDate.Month() < dob.Month() || (targettargetDate.Month() == dob.Month() && targetDate.Day() < dob.Day()) {
			age--
		}

		output = append(output, Person{
			Name: name,
			Age:  age,
		})
	}

	// Encode the resulting slice to a JSON array string.
	jsonBytes, err := json.Marshal(output)
	if err != nil {
		fmt.Print("[]")
		return
	}

	// Print the JSON array to stdout.
	fmt.Print(string(jsonBytes))
}