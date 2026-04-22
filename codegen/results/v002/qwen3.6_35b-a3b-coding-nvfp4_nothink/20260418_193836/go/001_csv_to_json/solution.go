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

	var result []Record
	for i, row := range records {
		if i == 0 {
			continue // skip header
		}
		if len(row) < 4 {
			continue
		}
		age, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age: %v\n", err)
			os.Exit(1)
		}
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score: %v\n", err)
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

	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(jsonData))
}