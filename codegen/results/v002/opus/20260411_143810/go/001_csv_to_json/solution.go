package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
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
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var records []Record
	for _, row := range rows[1:] {
		age, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal(err)
		}
		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			log.Fatal(err)
		}
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

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(records); err != nil {
		log.Fatal(err)
	}
}