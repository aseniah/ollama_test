package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Row represents a single row from the CSV file
type Row struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Email string `json:"email"`
	Score float64 `json:"score"`
}

// processCSV reads the CSV file and processes each row to create Row objects,
// then writes them to JSON format.
func processCSV(filePath string) error {
	// Read the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var rows []Row
	var currentRow Row
	reader := csv.NewReader(file)

	// Skip header
	reader.ReadHeader()

	// Process each row
	for {
		row := reader.Read()
		if row == nil {
			break
		}

		if len(row) < 4 {
			log.Printf("Skipping incomplete row: %#v", row)
			continue
		}

		// Parse each field
		currentRow.Name = row[0]
		currentRow.Age = int(row[1])
		currentRow.Email = row[2]
		currentRow.Score = float64(row[3])

		// Append to rows
		rows = append(rows, currentRow)
	}

	// Write JSON array
	if err := json.SetIndent(json.Indent("  "), 2); err != nil {
		return fmt.Errorf("json.SetIndent failed: %w", err)
	}

	if err := json.Marshal(rows); err != nil {
		return fmt.Errorf("json.Marshal failed: %w", err)
	}

	return nil
}

func main() {
	filePath := "input/data.csv"
	if err := processCSV(filePath); err != nil {
		log.Fatalf("Error processing CSV: %v", err)
	}
	fmt.Println(string([]byte(string([]byte([]byte([]byte([]byte("{\n"))))))))
}