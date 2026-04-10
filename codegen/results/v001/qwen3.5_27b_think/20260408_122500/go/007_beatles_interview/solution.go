package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	targetDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	headers := records[0]
	dataRows := records[1:]

	nameIdx := -1
	dobIdx := -1
	dateFormat := "2006-01-02" // Default format

	// Identify column indices
	for i, header := range headers {
		lowerHeader := strings.ToLower(strings.TrimSpace(header))
		if lowerHeader == "name" || lowerHeader == "full name" || lowerHeader == "person" {
			nameIdx = i
		}
		if lowerHeader == "dob" || lowerHeader == "date of birth" || lowerHeader == "birth date" || lowerHeader == "birthdate" || lowerHeader == "dateborn" {
			dobIdx = i
		}
	}

	// Fallback if headers not found
	if nameIdx == -1 {
		nameIdx = 0
	}
	if dobIdx == -1 {
		dobIdx = 1
	}

	output := make([]map[string]interface{}, 0, len(dataRows))

	for _, row := range dataRows {
		record := make(map[string]interface{})

		// Map original columns to lowercase keys
		for i, header := range headers {
			if i == len(row) {
				break
			}
			key := strings.ToLower(strings.Trim(strings.TrimSpace(header), " "))
			record[key] = strings.TrimSpace(row[i])
		}

		// Calculate and add age
		var age int = -1
		if dobIdx < len(row) {
			birthStr := strings.TrimSpace(row[dobIdx])
			// Try common formats
			var birthDate time.Time
			formats := []string{
				"2006-01-02",
				"01/02/2006",
				"02/01/2006",
				"2006/01/02",
			}
			for _, f := range formats {
				parsed, err := time.Parse(f, birthStr)
				if err == nil {
					birthDate = parsed
					break
				}
			}

			if !birthDate.IsZero() {
				age = targetDate.Year() - birthDate.Year()
				// Adjust if birthday hasn't occurred yet in the target year
				bday := time.Date(targetDate.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, birthDate.Location())
				if targetDate.Before(bday) {
					age--
				}
			}
		}

		record["age"] = age
		output = append(output, record)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(output); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}