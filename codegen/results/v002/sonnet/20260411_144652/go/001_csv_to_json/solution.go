package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Row struct {
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
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var rows []Row
	for _, rec := range records[1:] {
		age, err := strconv.Atoi(rec[1])
		if err != nil {
			log.Fatal(err)
		}
		score, err := strconv.ParseFloat(rec[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		rows = append(rows, Row{
			Name:  rec[0],
			Age:   age,
			Email: rec[2],
			Score: score,
		})
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(rows); err != nil {
		log.Fatal(err)
	}
}
