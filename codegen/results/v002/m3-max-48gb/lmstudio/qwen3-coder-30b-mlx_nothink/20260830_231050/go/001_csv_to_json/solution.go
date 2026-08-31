package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Person struct {
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
	
	var people []Person
	
	// Read the header row
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}
	
	// Verify expected headers
	expectedHeaders := []string{"Name", "Age", "Email", "Score"}
	for i, expected := range expectedHeaders {
		if i >= len(header) || header[i] != expected {
			fmt.Fprintf(os.Stderr, "Unexpected header: expected %s, got %s\n", expected, header[i])
			os.Exit(1)
		}
	}
	
	// Process each data row
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading record: %v\n", err)
			os.Exit(1)
		}
		
		// Validate record has expected number of fields
		if len(record) < 4 {
			fmt.Fprintf(os.Stderr, "Record has insufficient fields: %v\n", record)
			os.Exit(1)
		}
		
		// Parse Age
		age, err := strconv.Atoi(strings.TrimSpace(record[1]))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age: %v\n", err)
			os.Exit(1)
		}
		
		// Parse Score
		score, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score: %v\n", err)
			os.Exit(1)
		}
		
		person := Person{
			Name:  strings.TrimSpace(record[0]),
			Age:   age,
			Email: strings.TrimSpace(record[2]),
			Score: score,
		}
		
		people = append(people, person)
	}
	
	// Output JSON
	jsonBytes, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("%s", jsonBytes)
}