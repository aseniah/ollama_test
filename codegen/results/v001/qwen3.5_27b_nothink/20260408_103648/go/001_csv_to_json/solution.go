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
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
		os.Exit(1)
	}

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	// Skip header
	dataRows := records[1:]
	var results []Record

	for _, row := range dataRows {
		if len(row) < 4 {
			continue // Skip malformed rows
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			age = 0 // Default to 0 on parse error
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			score = 0.0 // Default to 0.0 on parse error
		}

		results = append(results, Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}