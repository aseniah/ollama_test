package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// parseCSV reads and parses the input CSV file.
func parseCSV(path string) ([]map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return []map[string]string{}, nil
	}

	// The first line is the header.
	header := lines[0]
	headers := strings.Split(header, ",")

	var rows []map[string]string
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		fields := strings.Split(line, ",")
		record := make(map[string]string)
		for i, h := range headers {
			if i < len(fields) {
				record[h] = fields[i]
			}
		}
		rows = append(rows, record)
	}

	return rows, nil
}

// transformCSVToJSON transforms the parsed CSV rows into a JSON array.
// It includes: id, name, age (calculated as of July 1, 2025), and date_format.
func transformCSVToJSON(rows []map[string]string) ([]map[string]interface{}, error) {
	if len(rows) == 0 {
		return []map[string]interface{}{}, nil
	}

	// Determine the expected format based on the first row's columns.
	expectedFormat := rows[0]
	var colOrder []string
	for key := range expectedFormat {
		colOrder = append(colOrder, key)
	}
	sort.Strings(colOrder) // Sorting ensures consistent column order if needed, but JSON keys are not ordered in spec strictly. 
                           // However, the example shows specific columns: "id", "name", "age", "date_format". Let's assume these are present and mapped by name.

	type OutputRecord struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
	Age           int    `json:"age"`
DateFormat     string `json:"date_format"`
	}

	var result []OutputRecord

	for _, row := range rows {
		var id, name string
		if ageVal, ok := row["age_str"]; ok {
			id = row["id"]
			name = row["name"]
			var age int
			fmt.Sscanf(ageVal, "%d", &age)
			result = append(result, OutputRecord{ID: id, Name: name, Age: age})
		} else if colAge := findKeyInMap(row, "age"); colAge != "" {
			id = row["id"]
			name = row["name"]
			var age int
			fmt.Sscanf(colAge, "%d", &age)
			result = append(result, OutputRecord{ID: id, Name: name, Age: age})
		} else {
			id = row["id"]
			name = row["name"]
			var age int
			if colAge := findKeyInMap(row, "age"); colAge != "" {
				fmt.Sscanf(colAge, "%d", &age)
			}
			result = append(result, OutputRecord{ID: id, Name: name, Age: -1})
		}

		// If age is not explicitly provided, we might need to calculate it. 
		// But the problem says "Calculate ages as of July 1, 2025." 
		// So we assume the file has DOB (Date of Birth) and we calculate age based on that.
		if dobVal, ok := row["date_of_birth"]; ok {
			dobStr := dobVal
			dob := parseDate(dobStr)
			if !dob.IsZero() {
				result = append(result, OutputRecord{ID: id, Name: name, Age: int(dob.Year())}) 
			}
		} else if colDOB := findKeyInMap(row, "date_of_birth"); colDOB != "" {
			dobStr := row[colDOB]
			dob := parseDate(dobStr)
			if !dob.IsZero() {
				result = append(result, OutputRecord{ID: id, Name: name, Age: int(dob.Year())}) 
			}
		}

		var dateFormat string
		if colFormat := findKeyInMap(row, "date_format"); colFormat != "" {
			dateFormat = row[colFormat]
		} else {
			dateFormat = "YYYY-MM-DD" // Default format if not specified
		}

		result[len(result)-1].DateFormat = dateFormat
	}

	return result, nil
}

func findKeyInMap(m map[string]string, key string) string {
	if v, ok := m[key]; ok {
		return v
	}
	return ""
}

func parseDate(str string) time.Time {
	t1, _ := time.Parse("2006-01-02", str)
	return t1
}

func main() {
	// Read the CSV file.
	rows, err := parseCSV("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input.csv:", err)
		os.Exit(1)
	}

	// Transform the CSV rows to JSON format.
	result, err := transformCSVToJSON(rows)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error transforming data:", err)
		os.Exit(1)
	}

	// Output the result as a JSON array to stdout.
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}
	fmt.Print(string(jsonBytes))
}