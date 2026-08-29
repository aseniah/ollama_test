package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(records) == 0 {
		fmt.Print("[]")
		return
	}

	// Use header to map column indices
	header := records[0]
	nameIdx, ageIdx, emailIdx, scoreIdx := -1, -1, -1, -1
	for i, h := range header {
		switch h {
		case "Name":
			nameIdx = i
		case "Age":
			ageIdx = i
		case "Email":
			emailIdx = i
		case "Score":
			scoreIdx = i
		}
	}

	people := make([]Person, 0, len(records)-1)
	for i := 1; i < len(records); i++ {
		row := records[i]
		p := Person{}
		if nameIdx >= 0 && nameIdx < len(row) {
			p.Name = row[nameIdx]
		}
		if ageIdx >= 0 && ageIdx < len(row) {
			age, err := strconv.Atoi(row[ageIdx])
			if err != nil {
				age = 0
			}
			p.Age = age
		}
		if emailIdx >= 0 && emailIdx < len(row) {
			p.Email = row[emailIdx]
		}
		if scoreIdx >= 0 && scoreIdx < len(row) {
			score, err := strconv.ParseFloat(row[scoreIdx], 64)
			if err != nil {
				score = 0
			}
			p.Score = score
		}
		people = append(people, p)
	}

	out, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(string(out))
}