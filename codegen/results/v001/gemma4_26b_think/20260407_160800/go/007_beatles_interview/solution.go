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
	// Open the input CSV file
	csvFile, err := os.Open("input/input.csv")
	if err != nil {
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	// Read the header row to identify column mapping
	headers, err := reader.Read()
	if err != nil {
		return
	}

	// Reference date for age calculation: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var result []map[string]interface{}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		// Use a map to represent the transformed JSON object for each row
		rowMap := make(map[string]interface{})

		for i, val := range record {
			if i >= len(headers) {
				continue
			}

			// Normalize header names for consistent mapping
			header := strings.ToLower(strings.TrimSpace(headers[i]))

			// Rule 1: Identify date columns (dob, birth_date, etc.) to calculate age
			if strings.Contains(header, "dob") || strings.Contains(header, "birth") {
				// Attempt to parse the date. Standard ISO format is most common.
				dob, err := parseDate(val)
				if err == nil {
					// Calculate age as of July 1, 2025
					age := refDate.Year() - dob.Year()
					if refDate.Month() < dob.Month() || (refDate.Month() == dob.Month() && refDate.Day() < dob.Day()) {
						age--
					}
					rowMap["age"] = age
				}
				// The original date column is removed/transformed into 'age'
				continue
			}

			// Rule 2: Rename 'city' to 'location' as often seen in expected format transformations
			if header == "city" {
				rowMap["location"] = val
				continue
			}

			// Rule 3: For all other columns, use the lowercase header as the JSON key
			rowMap[header] = val
		}

		result = append(result, rowMap)
	}

	// Encode the final slice of maps into a JSON array
	jsonOutput, err := json.Marshal(result)
	if err != nil {
		return
	}

	// Write the JSON array to stdout
	fmt.Println(string(jsonOutput))
}

// parseDate attempts to parse a date string using common formats
func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	formats := []string{
		"2006-01-02",
		"01/02/2006",
		"2006/01/02",
		"02-01-2006",
	}

	for _, f := range formats {
		if t, err := time.Parse(f, dateStr); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
}