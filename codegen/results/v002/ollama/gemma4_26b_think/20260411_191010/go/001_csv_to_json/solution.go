package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
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
	filePath := "input/data.csv"
	file, err := os.Open(filePath)
	if err != nil {
		// We use log.Fatal which writes to stderr, 
		// keeping stdout clean for the required JSON output.
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize CSV reader
	reader := csv.NewReader(file)

	// Read the header row first to skip it
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Read all remaining rows
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var people []Person

	// Iterate through records and convert them to Person structs
	for _, record := range records {
		// Record indices based on input: 0:Name, 1:Age, 2:Email, 3:Score
		if len(record) < 4 {
			continue
		}

		name := record[0]
		email := record[2]

		// Convert Age string to int
		age, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}

		// Convert Score string to float64
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		person := Person{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		}

		people = append(people, person)
	}

	// Marshal the slice of Person into a JSON array
	// Using MarshalIndent for readability, though json.Marshal would also be valid
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Output the JSON array to stdout
	fmt.Println(string(jsonData))
}