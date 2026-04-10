using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read input.csv
string inputPath = "input/input.csv";
string expectedFormatPath = "input/expected_format.json";

List<string[]> records = new List<string[]>();
var csvReader = new StreamReader(inputPath);

while (csvReader.Peek() >= 0) {
    string line = csvReader.ReadLine();
    if (line == null || line.Trim().IsEmpty) break;
    records.Add(line.Split(',').SelectTrimmedToCharArray());
}
csvReader.Dispose();

// Read expected format to infer rules
var expectedJson = JsonNode.Parse(File.ReadAllText(expectedFormatPath))?.GetValue<double>();

List<Dictionary<string, object>> output = new List<Dictionary<string, object>>();

foreach (var record in records) {
    if (record.Length < 3) continue;
    
    string name = record[0].Trim();
    string dateStr = record[1].Trim();
    string city = record[2].Trim(); // Assuming 3rd column is city based on CSV structure
    
    DateTime birthDate;
    if (dateStr.Length == 4) {
        int year = Convert.ToInt32(dateStr);
        birthDate = new DateTime(year, 1, 1);
    } else if (dateStr.Length == 7) {
        int yearStart = dateStr.Substring(0, 4);
        int day = int.Parse(dateStr.Substring(4, 2));
        birthDate = new DateTime(Convert.ToInt32(yearStart), 1, day);
    } else if (dateStr.Length == 8) {
        int yearStart = dateStr.Substring(0, 4);
        int month = int.Parse(dateStr.Substring(4, 2));
        int day = int.Parse(dateStr.Substring(6, 2));
        birthDate = new DateTime(Convert.ToInt32(yearStart), month, day);
    } else {
        try {
            birthDate = DateTime.Parse(dateStr);
        } catch {
            birthDate = DateTime.MinValue; // Default if parsing fails
        }
    }

    DateTime referenceDate = new DateTime(2025, 7, 1);
    int ageYears = Convert.ToInt32((referenceDate.Year - birthDate.Year));
    
    int dayDiff = Convert.ToInt32((referenceDate.DayOfYear - birthDate.DayOfYear));
    if (dayDiff < 0) {
        ageYears -= 1;
    }
    
    output.Add(new Dictionary<string, object> {
        {"name", name},
        {"city", city},
        {"age", ageYears}
    });
}

var jsonString = JsonSerializer.Serialize(output);
Console.WriteLine(jsonString);