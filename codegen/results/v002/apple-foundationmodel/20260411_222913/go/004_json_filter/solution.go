package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Read input from file
func readInput(filePath string) ([]map[string]interface{}, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var data []map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return data, nil
}

// Filter and sort records
func filterAndSortRecords(data []map[string]interface{}) []map[string]interface{} {
	filtered := make([]map[string]interface{}, 0, len(data))
	for _, record := range data {
		if record["active"] != nil && record["age"] != nil &&
			*record["active"] == true && *record["age"] >= 30 {
			filtered = append(filtered, record)
		}
	}
	return filtered
}

func main() {
	filePath := "input/data.json"
	data, err := readInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	filteredData := filterAndSortRecords(data)

	// Sort by name ascending
	filteredData = sortByName(filteredData)

	// Serialize and print to stdout
	serializedData, err := json.MarshalIndent(filteredData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(serializedData))
}

// Helper function to sort by name
func sortByName(records []map[string]interface{}) []map[string]interface{} {
	return sort.Slice(records, func(i, j int) bool {
		return records[i]["name"] == records[j]["name"]
	})
}