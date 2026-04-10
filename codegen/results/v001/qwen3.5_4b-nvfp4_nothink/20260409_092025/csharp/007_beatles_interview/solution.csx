using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputPath = "input/input.csv";
var expectedPath = "input/expected_format.json";

// Read the reference format to infer column structure and logic
string expectedContent = File.ReadAllText(expectedPath);
JsonElement expectedRef = JsonDocument.Parse(expectedContent).RootElement;

// Infer columns from the expected JSON structure
List<string> inferredColumns = new List<string>();
if (expectedRef is JsonObject) {
    foreach (var kvp in ((JsonObject)expectedRef).EnumerateObject()) {
        // Assuming keys represent person names or identifiers, and we need to map them
        // Since input.csv likely has headers like "Name,DOB", let's assume the first few columns are names
        // Actually, looking at the typical structure: 
        // Column 0 might be Name (or multiple names), Column 1..N are DOBs.
        // However, without seeing the actual content, I'll assume a standard mapping:
        // The CSV file likely has rows with "Name,DateOfBirth" or similar.
        // Let's infer based on the structure. If expectedRef is an array of objects, each object represents a person.
        
        // Correction: Usually, 'input/expected_format.json' contains one line example of a single record.
        // So inferredColumns should be the keys of this JSON object (minus any internal metadata).
    }
}

// Since I cannot see the file content to literally infer column indices perfectly without visual inspection,
// I will assume a standard CSV format: Name, DOB (one date per row usually, or multiple if needed).
// Given the instruction "Calculate ages as of July 1, 2025", the transformation is simply: Age = Year(DOB) - Year(RangeStart) plus days.

var csvContent = File.ReadAllText(inputPath);
List<string> lines = csvContent.Split('\n').Where(l => !string.IsNullOrWhiteSpace(l)).ToList();

// If the first line contains headers, shift by 1. Else assume data starts immediately.
// Assuming header is column 0.
var records = new List<JsonElement>();

for (int i = 0; i < lines.Count; i++) {
    if (i == 0 && !string.IsNullOrEmpty(lines[i]) && (lines[i].Contains('"') || lines[i].Contains(','))) {
        // Likely header row
        continue; 
    }
    
    var line = lines[i];
    try {
        // Basic CSV parsing for simple comma-separated values.
        // Handling potential quotes around strings? Simple split first.
        string[] parts = line.Split(',', StringSplitOptions.None);
        
        if (parts.Length < 2) continue;

        // Assuming Column 0 is Name, Column 1 is DOB.
        // If the data format changed and names are multiple or structure differs, logic might fail.
        // But to produce a JSON array based on standard formats:
        string? name = null;
        string? dobStr = parts[1];

        if (i > 0) {
            // If first row is data, try to parse as date.
        } else {
            // First row might be header? 
        }
    } catch {
        // Skip
        continue;
    }
    
    // Correct parsing approach:
    // Re-parse carefully based on standard CSV.
    // Let's assume the first line (header) is skipped or handled.
    // If i == 0, assume header.
    if (i < lines.Length - 1) {
        // Assuming Header at index 0
        // Data starts at 1? Or if there are no headers, i=0 is data.
        // Let's assume `input/expected_format.json` shows a single entry with specific keys.
        // Common case: CSV has header "Name,DateOfBirth". 
        // So we parse parts[0] as Name, parts[1] as DOB.
    }
}

// Correct implementation of logic
string csvPath = "input/input.csv";
var allLines = System.IO.File.ReadAllLines(csvPath);
var dataLines = new List<string>();
bool skipHeader = true;
var outputArray = new System.Collections.Generic.List<JsonElement>();

foreach (var line in allLines) {
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    // Check if it's a header row: typically contains strings like "Name", "DateOfBirth", etc.
    // Since we can't see the content, I assume the first line is a header based on typical CSV standards unless otherwise specified.
    if (skipHeader) {
        skipHeader = false;
        continue;
    }
    
    dataLines.Add(line);
}

var referenceObj = JsonDocument.Parse(File.ReadAllText("input/expected_format.json")).RootElement;
// Assuming 'referenceObj' defines the structure of output items.
// Example structure: [{"Name": "...", "Age": ...}]
// We need to parse DOB and calculate age as of July 1, 2025.

var results = new System.Collections.Generic.List<Dictionary<string, string>>();
foreach (var line in dataLines) {
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    // Simple CSV split
    var parts = line.Split(',');
    if (parts.Length < 2) continue;
    
    var name = parts[0].Trim('"').Replace("'", "").Trim();
    var dobStr = parts[1].Trim();
    
    // Handle potential quotes around DOB or Name
    if (dobStr.StartsWith("\"") && dobStr.EndsWith("\"")) {
        dobStr = dobStr.Substring(1, dobStr.Length - 2);
    }
    
    try {
        DateTime? dob = null;
        
        // Try parsing date formats (YYYY-MM-DD, MM-DD-YYYY, DD/MM/YYYY)
        if (DateTime.TryParseExact(dobStr, "yyyy-MM-dd", null, System.Globalization.DateTimeStyles.None, out var d)) {
            dob = d;
        } else if (DateTime.TryParseExact(dobStr, "dd-MM-yyyy", null, System.Globalization.DateTimeStyles.None, out var d2)) {
            dob = d2;
        } else if (DateTime.TryParseExact(dobStr, "MM-dd-yyyy", null, System.Globalization.DateTimeStyles.None, out var d3)) {
            dob = d3;
        } else {
            // Fallback for DD-MM-YYYY or YYYY-DD-MM? 
            // Let's try common ISO (yyyy-MM-dd) again with specific patterns.
            if (DateTime.TryParseExact(dobStr, "dd-MMM-yyyy", null, System.Globalization.DateTimeStyles.None, out var d4)) {
                dob = d4;
            }
        }
        
        if (dob != null) {
            // Calculate age as of July 1, 2025
            var referenceDate = new DateTime(2025, 7, 1);
            
            // Calculate age logic: 
            // 1. Subtract DOB from Reference Date to get difference in days/years
            // 2. Add exact days if needed? Usually age is just years.
            // But sometimes age includes months/days if calculated precisely.
            // The prompt asks for "ages", implying integer years.
            int age = referenceDate.Year - dob.Value.Year;
            
            if (referenceDate < dob.Value) {
                // If person hasn't been born yet? Or simply logic for past dates.
            } else {
                // Check if birthday passed in this year
                if (dob.Value.Month > 7 || (dob.Value.Month == 7 && dob.Value.Day > 1)) {
                    age--; 
                }
                
                // Wait, reference date is July 1.
                // If DOB is Feb 20, 1950:
                // Reference: 2025-07-01. DOB: 1950-02-20.
                // Difference: 75 years (minus one day if birthday hasn't passed yet? No, birthday passed).
                // If DOB is July 20, 1950:
                // Reference: 2025-07-01. Birthday is August... No, Birthday is July 20.
                // Reference date (July 1) is BEFORE Birthday (July 20). So they haven't turned 76 yet. 
                // Age = 2025 - 1950 - 1 = 74.
            }
            
            var resultDict = new Dictionary<string, string> { { "Name", name }, { "Age", age.ToString() } };
            outputArray.Add(resultDict);
        }
    } catch {
        // Skip invalid rows
    }
}

var jsonSerializerOptions = new System.Text.Json.JsonSerializerOptions();
jsonSerializerOptions.WriteIndented = false;
jsonSerializerOptions.AllowTrailingCommas = true;

var jsonString = System.Text.Json.JsonSerializer.Serialize(outputArray, jsonSerializerOptions);
System.Console.WriteLine(jsonString);