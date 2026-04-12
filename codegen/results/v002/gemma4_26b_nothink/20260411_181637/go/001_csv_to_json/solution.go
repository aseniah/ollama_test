package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Person represents the structure of each row in the CSV
type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	// Open the input file
	file, err := os.Open("input/data.csv")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row first to skip it
	_, err = reader.Read()
	if err != nil {
		log.Fatalf("failed to read header: %v", err)
	}

	var people []Person

	// Iterate through the remaining rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading row: %v", err)
		}

		// Parse Age (string to int)
		age, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("error parsing age for %s: %v", record[0], err)
			continue
		}

		// Parse Score (string to float64)
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Printf("error parsing score for %s: %v", record[0], err)
			continue
		}

		// Create person object
		person := Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		}

		people = append(people, person)
	}

	// Encode the slice of people to JSON and write to stdout
	encoder := json.NewEncoder(os.Stdout)
	// Set indentation for pretty-printing (optional, but makes it valid and readable)
	encoder.SetIndent("", "  ")
	
	err = encoder.Encode(people)
	if err != nil {
		log.Fatalf("failed to encode json: %v", err)
	}
}