package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Read and parse the JSON input
func readData() ([]map[string]interface{}, error) {
	file, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// Filter and sort the records
func filterAndSortRecords(data []map[string]interface{}) ([]map[string]interface{}, error) {
	filtered := make([]map[string]interface{}, 0, len(data))

	for _, record := range data {
		name := record["name"].(string)
		age, ok := record["age"].(int)
		if ok && age >= 30 && record["active"].(bool) {
			filtered = append(filtered, record)
		}
	}

	filtered = strings.FieldsFunc(filtered, func(s string) bool {
		_, ok := s["active"].(bool)
		return ok && (s["age"].(int) >= 30)
	})

	return filtered, nil
}

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	filteredAndSorted, err := filterAndSortRecords(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(json.MarshalIndent(filteredAndSorted, "", "  "))
}