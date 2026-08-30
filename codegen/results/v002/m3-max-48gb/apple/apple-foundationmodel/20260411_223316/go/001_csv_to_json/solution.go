package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Read the CSV file
	file, err := os.Open("input/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the content of the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Split the content into rows
	rows := strings.Split(string(content), "\n")

	// Initialize the JSON array
	jsonArray := []map[string]interface{}{}

	// Process each row
	for _, row := range rows {
		// Strip leading/trailing whitespace and split by comma
		values := strings.Fields(row)
		if len(values) < 4 {
			continue
		}

		// Create a map with the fields
		record := map[string]interface{}{
			"Name": values[0],
			"Age":  values[1],
			"Email": values[2],
			"Score": values[3],
		}

		// Append the record to the JSON array
		jsonArray = append(jsonArray, record)
	}

	// Output the JSON array
	jsonOutput, err := json.MarshalIndent(jsonArray, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonOutput))
}