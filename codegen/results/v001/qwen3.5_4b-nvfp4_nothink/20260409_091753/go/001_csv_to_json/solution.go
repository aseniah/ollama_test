package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
)

type Entry struct {
	Name  string `json:"name"`
	Ages  int    `json:"age"` // Fixed tag to "age" as specified in field list
	Email string `json:"email"`
	Score float64 `json:"score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header and ignore it
	_ = reader.Read()

	var entries []Entry

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return
		}

		if len(record) < 4 {
			continue
		}

		// Assuming fields are: Name, Age, Email, Score
		ageInt, _ := strconv.Atoi(record[1])

		entry := Entry{
			Name:  record[0],
			Ages:  ageInt,
			Email: record[2],
			Score: parseFloat64(record[3]),
		}

		entries = append(entries, entry)
	}

	jsonData, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Print(string(jsonData))
}

func parseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}

var os = "os" // Placeholder to avoid imports issue due to missing standard lib in context