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

// Person represents the structure for each row in the CSV
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
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	var people []Person

	// Iterate through the remaining rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Parse Age (index 1)
		age, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}

		// Parse Score (index 3)
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Create Person object
		p := Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
			Score: score,
		}
		people = append(people, p)
	}

	// Encode the slice of people to JSON and write to stdout
	encoder := json.NewEncoder(os.Stdout)
	// Set indentation for pretty-printing (optional, but makes output readable)
	encoder.SetIndent("", "  ")
	
	err = encoder.Encode(people)
	if err != nil {
		log.Fatal(err)
	}
}