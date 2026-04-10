package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type Record struct {
	Name  string  `json:"Name"`
	Age   int     `json:"Age"`
	Email string  `json:"Email"`
	Score float64 `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	var results []Record
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 4 {
			continue
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			age = 0
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			score = 0.0
		}

		results = append(results, Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	encoder := json.NewEncoder(os.Stdout)
	_ = encoder.Encode(results)
}