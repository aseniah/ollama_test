package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "time"
)

// person represents an individual in the CSV input
type person struct {
    Name        string  `json:"name"`
    Age         int     `json:"age"`
    BirthDate   time.Time `json:"birthdate"`
}

func main() {
    // Open and read CSV file
    file, err := os.Open("input/input.csv")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
        os.Exit(1)
    }
    
    var results []person
    
    // Process each row (skip header row)
    for i, record := range records {
        // Skip header row
        if i == 0 {
            continue
        }
        
        // Skip empty records
        if len(record) < 2 {
            continue
        }
        
        name := record[0]
        
        // Parse birth date - expect format like 2000-05-25
        birthDateStr := record[1]
        birthDate, err := time.Parse("2006-01-02", birthDateStr)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error parsing date %s: %v\n", birthDateStr, err)
            continue
        }
        
        // Calculate age as of July 1, 2025
        refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
        age := calculateAge(birthDate, refDate)
        
        // Create person struct
        p := person{
            Name:        name,
            Age:         age,
            BirthDate:   birthDate,
        }
        results = append(results, p)
    }
    
    // Output as JSON array to stdout
    jsonData, err := json.MarshalIndent(results, "", "  ")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println(string(jsonData))
}

// calculateAge computes the age of a person as of the reference date
func calculateAge(birthDate time.Time, refDate time.Time) int {
    age := refDate.Year() - birthDate.Year()
    
    // Check if birthday has already occurred in the reference year
    if (birthDate.Month() > refDate.Month() ||
        (birthDate.Month() == refDate.Month() && birthDate.Day() > refDate.Day())) {
        age--
    }
    
    return age
}