package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Name  string
	Age   int
	Email string
	Score float64
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	record, err := reader.Read() // Skip header row
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	var dataObjects []Data
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		var age float64
		var score float64

		// Parse Age
		if len(record) > 1 {
			age, err = fmt.Sscanf(record[1], "%d", &age)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing age: %v\n", err)
				os.Exit(1)
			}
		}

		// Parse Score
		if len(record) > 3 {
			score, err = fmt.Sscanf(record[3], "%f", &score)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing score: %v\n", err)
				os.Exit(1)
			}
		}

		dataObject := Data{
			Name:  record[0],
			Age:   int(age),
			Email: record[2],
			Score: score,
		}
		dataObjects = append(dataObjects, dataObject)
	}

	output, err := json.MarshalIndent(dataObjects, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}