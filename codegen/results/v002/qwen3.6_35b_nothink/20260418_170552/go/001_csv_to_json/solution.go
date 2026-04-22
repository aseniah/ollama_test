package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	// Skip header
	if len(records) < 1 {
		fmt.Println("[]")
		return
	}

	var result []Record
	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatalf("Error parsing age: %v", err)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			log.Fatalf("Error parsing score: %v", err)
		}

		record := Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		}
		result = append(result, record)
	}

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	fmt.Println(string(output))
}