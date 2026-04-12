package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Define the output format for JSON
	var output []map[string]interface{}

	// Read the CSV file
	file, err := os.Open("input/data.csv")
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	// Read the content of the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Split the content into rows based on newlines
	rows := strings.Split(string(data), "\n")

	// Skip the header
	headerIndex := 0
	for i := 0; i < len(rows); i++ {
		if rows[i][0] == ',' {
			headerIndex = i
			break
		}
	}

	// Process each row
	for i := headerIndex + 1; i < len(rows); i++ {
		// Split each row into fields
		fields := strings.Split(rows[i], ",")
		if len(fields) != 4 {
			log.Printf("Skipping malformed row: %s", rows[i])
			continue
		}

		// Create a map for the row
		row := make(map[string]interface{})
		row["Name"] = fields[0]
		row["Age"] = int(fields[1])
		row["Email"] = fields[2]
		row["Score"] = float64(fields[3])

		// Append the row to the output slice
		output = append(output, row)
	}

	// Convert the slice of maps to JSON and print
	err = json.NewEncoder(os.Stdout).Encode(output)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
}