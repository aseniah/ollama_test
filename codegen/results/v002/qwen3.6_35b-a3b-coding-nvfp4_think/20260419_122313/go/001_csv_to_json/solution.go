package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
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
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var records []Record

	// Skip header row
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if len(row) < 4 {
			continue
		}

		age, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}

		score, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			panic(err)
		}

		records = append(records, Record{
			Name:  row[0],
			Age:   age,
			Email: row[2],
			Score: score,
		})
	}

	jsonData, err := json.Marshal(records)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(jsonData)
	os.Stdout.Write([]byte("\n"))
}