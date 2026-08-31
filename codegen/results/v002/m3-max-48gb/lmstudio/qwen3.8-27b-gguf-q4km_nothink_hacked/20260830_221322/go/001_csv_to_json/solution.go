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
	
	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Println("[]")
		return
	}

	// Skip header (first row)
	result := make([]Record, 0, len(records)-1)
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 4 {
			continue
		}
		
		age, err := strconv.Atoi(row[1])
		if err != nil {
			// Skip invalid rows
			continue
		}
		
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			// Skip invalid rows
			continue
		}
		
		record := Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		result = append(result, record)
	}

	// Marshal to JSON
	jsonOutput, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Output to stdout
	fmt.Println(string(jsonOutput))
}