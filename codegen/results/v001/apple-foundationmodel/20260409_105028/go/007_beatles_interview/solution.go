package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Person struct {
	Name string `json:"name"`
	DOB string `json:"dob"`
}

func parseDate(dateStr string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(dateStr))
}

func calculateAge(dob string) int {
	year, err := parseDate(dob)
	if err != nil {
		log.Fatalf("Error parsing date: %v", err)
	}
	return 2025 - year
}

func main() {
	var args = &cobra.Command{
		Use:   "process-csv",
		Short: "Process CSV file to calculate ages and output JSON",
	}
	args.PersistentFlags().StringP("input", "i", "", "Input CSV file path")
	args.PersistentFlags().StringP("output", "o", "", "Output JSON file path")

	if err := args.Parse(); err != nil {
		log.Fatalf("Failed to parse arguments: %v", err)
	}

	inputFile := args.PersistentFlags().GetString("input")
	outputFile := args.PersistentFlags().GetString("output")

	var persons []Person

	// Read CSV file
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	reader := csv.NewReader(strings.NewReader(string(data)))
	if err := reader.ReadAll(&persons); err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	// Calculate ages and write to JSON
	var jsonData []byte
	err = json.MarshalIndent(persons, jsonData, "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Write to output file
	if err := ioutil.WriteFile(outputFile, jsonData, 0644); err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	log.Printf("Successfully processed data and written to %s", outputFile)
}