package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Record struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Score float64 `json:"score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if reader.CheckFieldCount == nil {
		reader.FieldsPerRecord = 4
	}
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading header: %v\n", err)
		os.Exit(1)
	}

	var records []Record
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading row: %v\n", err)
			os.Exit(1)
		}

		if len(record) != 4 || record[0] != header[0] || record[1] != header[1] || record[2] != header[2] || record[3] != header[3] {
			continue
		}

		records = append(records, Record{
			Name:  record[0],
			Age:   parseAge(record[1]),
			Email: record[2],
			Score: parseScore(record[3]),
		})
	}

	output, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	_, err = os.Stdout.Write(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing to stdout: %v\n", err)
		os.Exit(1)
	}
}

func parseAge(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}

func parseScore(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return val
}