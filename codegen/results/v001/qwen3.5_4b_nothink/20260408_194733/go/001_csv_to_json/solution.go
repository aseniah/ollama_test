package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type UserData struct {
	Name  string
	Age   int
	Email string
	Score float64
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var users []UserData

	lineNum := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading CSV at line %d: %v\n", lineNum, err)
			os.Exit(1)
		}

		lineNum++
		if lineNum == 1 {
			// Skip header
			continue
		}

		if len(record) < 4 {
			fmt.Fprintf(os.Stderr, "Error: insufficient fields in line %d\n", lineNum)
			os.Exit(1)
		}

		name := record[0]
		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing age in line %d: %v\n", lineNum, err)
			os.Exit(1)
		}
		email := record[2]
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing score in line %d: %v\n", lineNum, err)
			os.Exit(1)
		}

		user := UserData{
			Name:  name,
			Age:   age,
			Email: email,
			Score: score,
		}
		users = append(users, user)
	}

	output, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}