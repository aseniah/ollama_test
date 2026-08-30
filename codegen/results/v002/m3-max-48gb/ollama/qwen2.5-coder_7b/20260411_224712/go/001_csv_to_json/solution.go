package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Open the input file
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Skip the header row
	data := make([]map[string]interface{}, 0, len(records)-1)
	for i := 1; i < len(records); i++ {
		row := records[i]
		entry := map[string]interface{}{
			"Name":  row[0],
			"Age":   row[1],
			"Email": row[2],
			"Score": row[3],
		}
		data = append(data, entry)
	}

	// Output the JSON array
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}