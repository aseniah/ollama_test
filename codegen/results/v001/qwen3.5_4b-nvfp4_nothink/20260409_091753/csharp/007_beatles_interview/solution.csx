using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json.Nodes;

// Read and parse the expected format to infer transformation rules
string expectedFormatPath = "input/expected_format.json";
using (var reader = new StreamReader(expectedFormatFormatPath))
{
    var expectedJsonString = reader.ReadToEnd();
    // Parse the expected JSON string to understand the structure and rules
    dynamic expected = JsonNode.Parse(expectedJsonString);
}

// Function to parse a date from a string in the format "YYYY-MM-DD"
string? parseDate(string? dateString)
{
    if (string.IsNullOrWhiteSpace(dateString)) return null;
    
    var dateComponents = dateString.Split('-');
    if (dateComponents.Length != 3) return null;

    int year, month, day;
    try {
        year = int.Parse(dateComponents[0]);
        month = int.Parse(dateComponents[1]);
        day = int.Parse(dateComponents[2]);
        
        // Basic validation
        if (month < 1 || month > 12 || day < 1 || day > 31) return null;
    }
    catch {
        return null;
    }
    
    return $"{year}-{month:00}-{day:00}"; // Return a normalized date string
}

// Main logic to process the input CSV file and produce output JSON array
var inputFilePath = "input/input.csv";
string inputFileContent = File.ReadAllText(inputFilePath);

var lines = inputFileContent.Split(new[] { "\r\n", "\n" }, StringSplitOptions.RemoveEmptyEntries);

var resultList = new List<object>();

foreach (var line in lines)
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    var values = line.Split(',', StringSplitOptions.TrimEntries | StringSplitOptions.RemoveEmptyStrings).Select(s => s?.Trim()).ToArray();
    
    // Assuming the CSV has 6 columns: Name, Birth Date, Address, Age Category, Contact Status, Notes
    if (values.Length != 6) continue;

    string? name = values[0];
    string? birthDateStr = values[1];
    string? address = values[2];
    string? ageCategory = values[3];
    string? contactStatus = values[4];
    string? notes = values[5];

    // Parse the birth date string into a DateTime object
    var parsedDateStr = parseDate(birthDateStr);
    
    if (string.IsNullOrEmpty(parsedDateStr)) {
        continue; 
    }

    var birthDate = new DateTime(int.Parse(parsedDateStr.Replace('-', '1234-'))); // Adjust for the date string as per the CSV input format
    
    // If the date is invalid, skip this row
    if (birthDate.Day == 0 || birthDate.Month == 0) continue;

    var referenceDate = new DateTime(2025, 7, 1);
    
    var age = int.Parse(referenceDate.Subtract(birthDate.Date).TotalDays.ToString()); // Adjust for the date string as per the CSV input format
    
    var formattedAge = $"{age}"; // Adjust for the date string as per the CSV input format

    resultList.Add(new { name, formattedAge });
}

// Output the result list as a JSON array to stdout
Console.WriteLine(System.Text.Json.Serialization.JsonIgnoreConverters.EncodeObject(resultList.ToArray()));