using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvPath = "input/input.csv";
var formatPath = "input/expected_format.json";
var birthDateColumn = null;
var outputFormat = null;

// Parse expected format to understand structure
if (File.Exists(formatPath)) {
    var jsonStr = File.ReadAllText(formatPath);
    outputFormat = JsonNode.Parse(jsonStr);
}

// Read CSV file and parse rows
var csvData = File.ReadAllText(csvPath).Split(new[] { '\n' }, StringSplitOptions.RemoveEmptyEntries);

if (csvData.Length == 0) {
    Console.OutputEncoding ??= Encoding.UTF8;
    Array.Empty<Dictionary<string, string>>();
}

var firstRow = csvData[0].Trim().ToLowerInvariant();
var headers = firstRow.Split(new[] { ',', ';' }, StringSplitOptions.RemoveEmptyEntries);

var birthDateColIndex = Array.IndexOf(headers, "date");
var ageColIndex = Array.IndexOf(headers, "age");
var nameColIndex = -1;
var colCount = 0;

for (int i = 0; i < headers.Length; i++) {
    var header = headers[i];
    if (header.ToLowerInvariant().StartsWith("name") || header.ToLowerInvariant().StartsWith("firstname")) {
        nameColIndex = i;
    }
    colCount = i + 1;
}

var results = new Dictionary<string, string>();

for (int i = 1; i < csvData.Length; i++) {
    var line = csvData[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    var row = line.Split(new[] { ',', ';' }, StringSplitOptions.RemoveEmptyEntries);
    if (row.Length < Math.Max(colCount, 2)) continue;
    
    // Find headers in current row that match column index
    var birthDateVal = -1;
    var dateIndex = -1;
    for (int j = 0; j < row.Length; j++) {
        var val = row[j].Trim().ToLowerInvariant();
        if (j == colCount && headers[colCount].ToLowerInvariant().StartsWith("date")) {
            birthDateVal = j + 1; // column index in row
            break;
        }
    }
    
    // Calculate age as of July 1, 2025
    var birthDateStr = "1900-01-01";
    var age = int.MinValue;
    
    if (line.Contains("birthday") || line.Contains("birth date")) {
        // Parse the birth date
        var parsedDate = DateTime.Parse(line.Substring(0, 10) + "-00" + line[6..7] + "-" + line[8..9]);
        var today = new DateTime(2025, 7, 1);
        age = (int)(today - parsedDate).TotalDays;
    } else {
        // Parse the birth date string from CSV
        if (line.Contains(",")) {
            var parts = line.Split(',');
            if (parts.Length >= 2) {
                var birthYear = int.Parse(parts[0].Trim());
                var birthMonth = int.Parse(parts[1].Trim());
                var birthDay = int.Parse(parts[2].Trim() + "1").Substring(0, 1);
                var birthDateStr = $"{birthYear}-{birthMonth:02}-{birthDay:02}";
                
                var parsedBirth = new DateTime(birthYear, birthMonth, birthDay);
                var today = new DateTime(2025, 7, 1);
                age = (int)(today - parsedBirth).TotalDays / 365.25;
            }
        }
    }
    
    if (age == int.MinValue) {
        // Parse as string
        var parts = line.Split(',');
        if (parts.Length >= 1) {
            try {
                var birthYear = int.Parse(parts[0].Trim());
                age = 2025 - birthYear;
                if (age < 0) age = 0;
            } catch { /* skip */ }
        }
    }
    
    // Create result object
    var resultObj = new System.Text.Json.Nodes.Object();
    resultObj["date"] = new DateTimeOffset(DateTime.Parse("1970-01-01").AddMonths(7).AddDay(1));
    resultObj["age"] = age.ToString();
    results[birthDateVal] = JsonNode.CreateObject().ToString(System.Text.Json.Serialization.MemberNameSerialization.OmitD);
}

var resultArray = new System.Collections.Generic.List<string>();
foreach (var entry in results) {
    resultArray.Add(entry.Key);
}

Console.OutputEncoding ??= Encoding.UTF8;
Console.Out.WriteLine(JsonConvert.SerializeObject(resultArray, new JsonSerializerOptions()));