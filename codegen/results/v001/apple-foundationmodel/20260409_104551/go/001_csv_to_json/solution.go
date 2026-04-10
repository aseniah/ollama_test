package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Function to process the CSV file and output JSON
func processCSV() {
	// Read the CSV file
	file, err := os.Open("input/data.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Read all data from the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Split the data into rows
	rows := strings.Split(string(data), "\n")

	// Skip the header
	rows = rows[1:]

	// Initialize a slice to hold the JSON objects
	jsonArray := make([]map[string]interface{}, 0, len(rows))

	// Process each row
	for _, row := range rows {
		values := strings.Fields(row)
		if len(values) < 4 {
			log.Printf("Skipping incomplete row: %s", row)
			continue
		}

		var name, age, email, score interface{}
		err := strconv.Atoi(values[0]) // Age
		if err != nil {
			log.Printf("Failed to parse age: %s", values[0])
			continue
		}
		name = values[0]
		age = age

		err = strconv.Atoi(values[1]) // Age
		if err != nil {
			log.Printf("Failed to parse age: %s", values[1])
			continue
		}
		email = values[1]
		age = age

		err = strconv.Atoi(values[2]) // Age
		if err != nil {
			log.Printf("Failed to parse age: %s", values[2])
			continue
		}
		score = values[2]
		email = email

		jsonObject := map[string]interface{}{
			"Name": name,
			"Age":   age,
			"Email": email,
			"Score": score,
		}
		jsonArray = append(jsonArray, jsonObject)
	}

	// Output the JSON array
	fmt.Println(json.MarshalIndent(jsonArray, "", "  "))
}

func main() {
	processCSV()
}