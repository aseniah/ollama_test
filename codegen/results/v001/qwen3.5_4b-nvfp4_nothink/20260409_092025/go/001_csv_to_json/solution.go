package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type DataRow struct {
	Name  string   `json:"name"`
	Age   *int     `json:"age"` // Using pointer to handle unexported int field issue in json.Marshal
	Email string   `json:"email"`
	Score *float64 `json:"score"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row (skip it)
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	var rows []DataRow

	// Read data rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading row: %v\n", err)
			os.Exit(1)
		}

		if len(record) < 4 {
			continue // Skip incomplete rows
		}

		row := DataRow{
			Name: record[0],
			Email: record[2],
		}

		var ageInt int
		_, err = fmt.Sscanf(record[1], "%d", &ageInt)
		if err != nil {
			continue // Skip rows with invalid age
		}
		row.Age = &ageInt

		var scoreFloat float64
		_, err = fmt.Sscanf(record[3], "%f", &scoreFloat)
		if err != nil {
			continue // Skip rows with invalid score
		}
		row.Score = &scoreFloat

		rows = append(rows, row)
	}

	// Marshal to JSON and write to stdout
	jsonData, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	_, err = os.Stdout.WriteString(string(jsonData) + "\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to stdout: %v\n", err)
		os.Exit(1)
	}
}