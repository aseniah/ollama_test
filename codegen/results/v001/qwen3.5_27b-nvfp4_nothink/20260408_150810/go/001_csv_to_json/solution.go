package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Person struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Skip header
	_, err = reader.Read()
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Failed to read header: %v\n", err)
		os.Exit(1)
	}

	var people []Person
	
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read row: %v\n", err)
			os.Exit(1)
		}

		// Assuming the CSV structure is: Name, Age, Email, Score
		if len(record) < 4 {
			fmt.Fprintf(os.Stderr, "Invalid row format: expected 4 columns, got %d\n", len(record))
			os.Exit(1)
		}

		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse Age for row %v: %v\n", record, err)
			os.Exit(1)
		}

		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse Score for row %v: %v\n", record, err)
			os.Exit(1)
		}

		person := Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		}
		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}