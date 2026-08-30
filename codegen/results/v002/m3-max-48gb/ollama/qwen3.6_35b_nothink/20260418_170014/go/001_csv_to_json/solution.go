package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Record struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
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

	// Skip header row
	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "No data rows found\n")
		os.Exit(1)
	}

	var result []Record
	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		if len(row) != 4 {
			fmt.Fprintf(os.Stderr, "Invalid row length at row %d\n", i)
			os.Exit(1)
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid age value: %v\n", err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid score value: %v\n", err)
			os.Exit(1)
		}

		record := Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		result = append(result, record)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}