package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// agePerson represents a person's name, date of birth, and calculated age.
type agePerson struct {
	Name     string `json:"name"`
	Dob      string `json:"dob"`
	Age      int    `json:"age"`
}

func main() {
	// Define the reference date for age calculation: July 1, 2025.
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0)

	// Open and read the input CSV file.
	inputFile, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input.csv: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	var people []agePerson

	// Read the CSV file line by line.
	var line string
	lineCount := 0

	for {
		line, err = inputFile.ReadLine()
		if err == io.EOF {
			break