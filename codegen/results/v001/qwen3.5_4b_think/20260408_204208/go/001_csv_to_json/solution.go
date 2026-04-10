package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Read records (skip header by using a separate field or skipping first row)
	var data []Person
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		var p Person
		if row[0] != "" && row[0] != "Name" {
			p.Name = row[0]
		}
		if len(row) > 1 && row[1] != "" {
			a, err := strconv.Atoi(strings.TrimSpace(row[1]))
			if err != nil {
				continue
			}
			p.Age = a
		}
		if len(row) > 2 && row[2] != "" {
			p.Email = strings.TrimSpace(row[2])
		}
		if len(row) > 3 && row[3] != "" {
			f, err := strconv.ParseFloat(strings.TrimSpace(row[3]), 64)
			if err != nil {
				continue
			}
			p.Score = f
		}
		data = append(data, p)
	}

	// Output JSON array
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}