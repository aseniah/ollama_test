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
	f, err := os.Open("input/data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	var records []Record
	for _, row := range rows[1:] {
		age, _ := strconv.Atoi(row[1])
		score, _ := strconv.ParseFloat(row[3], 64)
		records = append(records, Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	if records == nil {
		records = []Record{}
	}

	json.NewEncoder(os.Stdout).Encode(records)
}
