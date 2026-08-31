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
	if len(records) < 2 {
		fmt.Print("[]\n")
		return
	}

	var data []Record
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 4 {
			continue
		}
		age, err := strconv.Atoi(row[1])
		if err != nil {
			age = 0
		}
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			score = 0.0
		}
		data = append(data, Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	if data == nil {
		data = []Record{}
	}

	output, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}