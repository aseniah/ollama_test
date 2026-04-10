package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Birthdate string `json:"birthdate"`
}

func calculateAge(birthdate string) int {
	// Assuming the birthdate is in YYYY-MM-DD format
	year, month, day, err := strconv.Atoi(birthdate[:4]), 0, 0, nil
	if err != nil {
		log.Printf("Error parsing birthdate: %s", err)
		return 0
	}

	// July 1, 2025 is 07/01/2025
	currentYear := 2025
	currentMonth := 7
	currentDay := 1

	// Calculate age
	if year < currentYear || (year == currentYear && month < currentMonth) ||
		(year == currentYear && month == currentMonth && day < currentDay) {
		return currentYear - year
	} else {
		return 0
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <input_file>", os.Args[0])
	}

	inputFile := os.Args[1]
	outputFile := "output.json"

	// Read CSV file and process each entry
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Delimiter = ","
	var people []Person

	for {
		record, err := reader.ReadRecord()
		if err != nil && err != io.EOF {
			log.Fatalf("Error reading record: %v", err)
		}
		if record == nil {
			break
		}

		name := strings.TrimSpace(record[0])
		birthdate := strings.TrimSpace(record[1])

		age := calculateAge(birthdate)
		people = append(people, Person{Name: name, Birthdate: birthdate})
	}

	// Write JSON output
	var jsonData []byte
	err = json.Marshal(people)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	fmt.Println(string(jsonData))
}