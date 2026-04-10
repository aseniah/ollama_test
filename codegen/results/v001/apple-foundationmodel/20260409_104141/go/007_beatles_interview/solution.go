package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/v2"
	"github.com/spf13/cobra/terminalui"
)

var (
	inputFile = "input/input.csv"
	outputFile = "output.json"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthYear int    `json:"birth_year"`
}

func main() {
	// Create the main command
	cmd := &cobra.Command{
		Use:   "read-csv-to-json",
		Short: "Convert CSV input to JSON output",
		Long:  `Read a CSV file and produce a JSON array with each person's first name, last name, and age as of July 1, 2025.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := readInput(args[0]); err != nil {
				return err
			}
			if err := writeOutputJSON(args[1]); err != nil {
				return err
			}
			return nil
		},
	}

	// Set command-line options
	cmd.PersistentFlags().StringVar(&inputFile, "input", "", "Input CSV file")
	cmd.PersistentFlagSet.StringVar(&outputFile, "output", "", "Output JSON file")

	// Parse command-line arguments
	if err := cmd.Execute(); err != nil {
		return
	}

	// Calculate ages as of July 1, 2025
	log.Printf("Calculating ages as of July 1, 2025...\n")
	calculateAges(inputFile)

	log.Printf("Output written to %s\n", outputFile)
}

func readInput(inputFile string) error {
	// Read the CSV file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var persons []Person
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("could not read line: %w", err)
		}
		if line == "" {
			continue
		}
		persons = append(persons, parsePerson(line))
	}

	return nil
}

func parsePerson(line string) Person {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return Person{FirstName: "", LastName: "", BirthYear: 0}
	}

	firstName := fields[0]
	lastName := fields[1]
	birthYear, err := strconv.Atoi(fields[2])
	if err != nil {
		return Person{FirstName: firstName, LastName: lastName, BirthYear: 0}
	}

	return Person{FirstName: firstName, LastName: lastName, BirthYear: birthYear}
}

func calculateAges(inputFile string) {
	// Read the input data
	var persons []Person
	if err := readInput(inputFile); err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Calculate ages as of July 1, 2025
	for _, person := range persons {
		age := 2025 - person.BirthYear
		fmt.Printf("Person: %s %s, Age: %d\n", person.FirstName, person.LastName, age)
	}
}

func writeOutputJSON(outputFile string) error {
	// Create JSON output
	var jsonData []byte
	err := json.MarshalIndent(persons, jsonData, "  ")
	if err != nil {
		return err
	}

	// Write JSON to file
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("could not open output file: %w", err)
	}
	defer file.Close()

	_, err = io.WriteString(file, `[` + string(jsonData) + `]`)
	if err != nil {
		return err
	}

	return nil
}