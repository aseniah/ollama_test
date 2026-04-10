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
	// Target date for age calculation: July 1, 2025
	targetDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	// Open the input CSV file
	f, err := os.Open("input/input.csv")
	if err != nil {
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)

	// Read the header row
	headers, err := reader.Read()
	if err != nil {
		return
	}

	// Prepare to store the resulting objects
	var results []map[string]interface{}

	// Identify which column contains the date for age calculation
	// We look for common date-related headers
	dateColIndex := -1
	for i, h := range headers {
		lowerH := strings.ToLower(h)
		if lowerH == "birthdate" || lowerH == "dob" || lowerH == "date_of_birth" || lowerH == "birth_date" {
			dateColIndex = i
			break
		}
	}

	// Iterate through the CSV rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		rowMap := make(map[string]interface{})

		for i, value := range record {
			if i >= len(headers) {
				continue
			}
			header := headers[i]

			// If this is the date column, calculate age and replace/add 'age'
			if i == dateColIndex {
				// Attempt to parse common date formats
				parsedDate, parseErr := parseDate(value)
				if parseErr == nil {
					age := calculateAge(parsedDate, targetDate)
					rowMap["age"] = age
				} else {
					// If parsing fails, we might fall back to the original value or skip
					// Based on typical transformation rules, we aim for the 'age' field.
					// We'll skip the original date field if it's the date column.
					continue
				}
			} else {
				// For all other columns, keep the value as a string
				rowMap[header] = value
			}
		}

		// If we didn't find a date column but the logic requires one, 
		// the record might still be valid if it contains other data.
		// However, we must ensure we don't add the date column as its original value 
		// if the transformation rules (implied by expected_format.json) suggest 
		// replacing birthdate with age.
		
		results = append(results, rowMap)
	}

	// Encode the resulting slice to JSON
	encoder := json.NewEncoder(os.Stdout)
	// Use indent for readability if required, but standard JSON array is requested.
	// The prompt implies a standard JSON array output.
	err = encoder.Encode(results)
	if err != nil {
		return
	}
}

// parseDate attempts to parse a date string using common formats
func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	formats := []string{
		"2006-01-02",
		"01/02/2006",
		"1/2/2006",
		"2006/01/02",
		"02-01-2006",
	}

	for _, f := range formats {
		t, err := time.Parse(f, dateStr)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
}

// calculateAge computes the age of a person as of the target date
func calculateAge(birthDate, targetDate time.Time) int {
	years := targetDate.Year() - birthDate.Year()

	// Adjust if the birthday has not occurred yet in the target year
	if targetDate.Month() < birthDate.Month() || 
	   (targetDate.Month() == birthDate.Month() && targetDate.Day() < birthDate.Day()) {
		years--
	}

	return years
}