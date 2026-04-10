package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
)

const referenceDate = time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

type Person struct {
	Name   string `json:"name"`
	Dob    string `json:"dob"`
	Age    int    `json:"age"`
	Format string `json:"format"`
}

func loadFormat(filename string) map[string]interface{} {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("error loading format file: %w", err))
	}

	var output interface{}
	if err := json.Unmarshal(data, &output); err != nil {
		panic(fmt.Errorf("error unmarshaling format file: %w", err))
	}

	return output.(map[string]interface{})
}

func calculateAge(birthDateStr string, birthTimeStr string) int {
	dob, err := time.Parse("2006-01-02 15:04:05", birthDateStr+birthTimeStr)
	if err != nil {
		dob, err = time.Parse("2006-01-02", birthDateStr)
		if err != nil {
			return -1
		}
	}

	birthTime, _ := time.Parse("15:04:05", birthTimeStr)
	fullDob := time.Date(dob.Year(), dob.Month(), dob.Day(), birthTime.Hour(), 0, 0, 0, dob.Location())

	ageResult := referenceDate.Sub(fullDob).Round(-time.Hour).Minutes() / 60
	return int(ageResult)
}

func main() {
	// Load expected format and extract field names from JSON structure
	outputFormat := loadFormat("input/expected_format.json")
	
	type FormatSchema struct {
		Field map[string]string `json:"field"`
	}
	
	if schema, ok := outputFormat["field"].(map[string]interface{}); ok {
		colName := ""
		for k, v := range schema {
			if f, ok := v.(string); ok {
				colName = f
			}
		}
		
		cols, _ := colName.(string)
	}
	
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading header:", err)
		os.Exit(1)
	}
	
	var records []Person
	
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading record:", err)
			continue
		}
		
		cols := make([]string, len(header))
		for i := range header {
			cols[i] = header[i]
		}
		
		for i := range header {
			if cols[i] == "name" || cols[i] == "dob" || cols[i] == "date_of_birth" || cols[i] == "age_calculated" {
				continue
			}
		}
		
		var p Person
		
		for i, col := range header {
			switch col {
			case "name":
				p.Name = record[i]
			case "dob", "date_of_birth":
				p.Dob = record[i]
			default:
				continue
			}
		}
		
		if p.Dob != "" {
			p.Age, _ = strconv.Atoi(record[4])
			format := outputFormat["format"].(string)
			p.Format = format
		}
		
		records = append(records, p)
	}
	
	outputData := make([]Person, 0, len(records))
	for i := range records {
		outputData[i] = records[i]
	}
	
	jsonBytes, err := json.MarshalIndent(outputData, "", " ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error marshaling to JSON:", err)
		os.Exit(1)
	}
	
	fmt.Println(string(jsonBytes))
}