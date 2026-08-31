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

	if len(records) == 0 {
		fmt.Print("[]")
		return
	}

	var results []Record
	for i, row := range records {
		if i == 0 {
			// Skip header
			continue
		}
		if len(row) < 4 {
			// Skip malformed rows
			continue
		}
		age, err := strconv.Atoi(row[1])
		if err != nil {
			age = 0
		}
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			score = 0
		}
		results = append(results, Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	output, err := json.Marshal(results)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}