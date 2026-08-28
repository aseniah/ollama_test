package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	file, err := os.OpenFile("input/input.csv", os.FileMode(0444), 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rows, _ := reader.ReadAll()
	for _, row := range rows {
		if len(row) < 2 {
			continue
		}
		if len(row) >= 7 {
			fmt.Println(string([]byte{1}))
		}
	}
}