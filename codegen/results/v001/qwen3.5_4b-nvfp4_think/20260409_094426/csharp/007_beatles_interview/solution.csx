using System;
using System.IO;
using System.Text;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputPath = "input/input.csv";
var expectedFormatPath = "input/expected_format.json";

// Check if files exist before reading
if (!File.Exists(inputPath)) {
    throw new FileNotFoundException($"Input file not found: {inputPath}");
}

if (!File.Exists(expectedFormatPath)) {
    throw new FileNotFoundException($"Expected format reference not found: {expectedFormatPath}");
}

// Read CSV file content
string csvContent = File.ReadAllText(inputPath);
var lines = csvContent.Split(new[] { "\r\n", "\n" }, StringSplitOptions.RemoveEmptyEntries);

// Parse header (first line) and data rows
var headers = new List<string>();
var dataRows = new List<List<object>>();

var trimmedLines = lines.Select(line => line.TrimEnd('\r', '\n'));
if (trimmedLines.Length > 0) {
    var headerLine = trimmedLines[0].Split(',');
    headers = headerLine.Select(h => h.Trim()).ToList();
}

for (int i = 1; i < trimmedLines.Length; i++) {
    var values = trimmedLines[i].Split(',');
    var row = new List<object>();
    
    foreach (var value in values) {
        if (!string.IsNullOrWhiteSpace(value)) {
            // Trim whitespace
            var trimmedValue = value.Trim();
            
            // Try to parse as DateTime
            try {
                DateTime dt;
                if (DateTime.TryParseExact(trimmedValue, "M/d/yyyy", null, System.Globalization.DateTimeStyles.None, out dt)) {
                    row.Add(dt);
                } else if (int.TryParse(trimmedValue, out int i)) {
                    row.Add(i);
                } else {
                    row.Add(trimmedValue);
                }
            } catch {
                row.Add(trimmedValue);
            }
        } else {
            row.Add(null);
        }
    }
    
    dataRows.Add(row);
}

// Build JSON array based on expected format structure
var resultArray = new List<Dictionary<string, object>>();

foreach (var rowData in dataRows) {
    var dictionary = new Dictionary<string, object>();
    
    for (int j = 0; j < headers.Count && j < rowData.Count; j++) {
        dictionary.Add(headers[j], rowData[j]);
    }
    
    // Calculate age as of July 1, 2025
    if (j + 1 < rowData.Count) {
        var birthDate = rowData[j + 1] as DateTime?;
        if (birthDate != null) {
            var birthYear = birthDate.Value.Year;
            var calculatedAge = DateTime.Parse("7/1/2025").Year - birthYear - 1;
            dictionary.Add("Age", Math.Max(0, calculatedAge));
        }
    }
    
    resultArray.Add(dictionary);
}

// Output JSON array to stdout
var options = new JsonSerializerOptions();
options.WriteIndented = false;
System.Console.WriteLine(JsonSerializer.Serialize(resultArray, options));