package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Birth string `json:"birth"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// Skip header if present (assuming first row is header based on typical CSV)
	// If the CSV doesn't have a header, adjust the starting index.
	startIndex := 0
	if len(records) > 0 && (records[0][0] == "Name" || records[0][0] == "name") {
		startIndex = 1
	}

	people := make([]Person, 0)

	// Target date for age calculation: July 1, 2025
	targetDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	// Expected date format in CSV: MM/DD/YYYY (inferred from standard US formats)
	dateLayout := "1/2/2006" // or "01/02/2006" to be safe for single digits

	for i := startIndex; i < len(records); i++ {
		row := records[i]
		if len(row) < 2 {
			continue // Skip malformed rows
		}

		name := row[0]
		birthDateStr := row[1]

		// Parse the birth date
		birthDate, err := time.Parse(dateLayout, birthDateStr)
		if err != nil {
			// Try alternative format just in case
			birthDate, err = time.Parse("01/02/2006", birthDateStr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Could not parse date '%s' for %s\n", birthDateStr, name)
				continue
			}
		}

		// Calculate Age as of July 1, 2025
		age := targetDate.Year() - birthDate.Year()

		// Adjust if birthday hasn't occurred yet in the target year
		if targetDate.YearDay() < birthDate.YearDay() {
			age--
		}

		// Re-format date to match common JSON output expectations (YYYY-MM-DD)
		birthDateFormatted := birthDate.Format("2006-01-02")

		people = append(people, Person{
			Name:  name,
			Age:   age,
			Birth: birthDateFormatted,
		})
	}

	// Output JSON to stdout
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(people); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}