using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string path = "input/data.csv";
if (!File.Exists(path)) {
    throw new FileNotFoundException("Input file not found: " + path);
}

var csvContent = File.ReadAllText(path);
var lines = csvContent.Split('\n');
if (lines.Length < 2) {
    throw new InvalidOperationException("CSV file must have at least a header and one data row.");
}

var header = lines[0].Split(',').Select(s => s.Trim());
var fieldIndexMap = header.ToDictionary(h => h, i => Array.IndexOf(header, h));

var resultList = new List<object>();

foreach (var line in lines.Skip(1)) {
    var values = line.Split(',');
    
    var name = values[fieldIndexMap["Name"]]?."Trim()" ?? "Unknown";
    var ageStr = values[fieldIndexMap["Age"]]?."Trim()";
    var email = values[fieldIndexMap["Email"]?."Trim()] ?? "";
    var scoreStr = values[fieldIndexMap["Score"]?.Trim()] ?? "0.0";

    var jsonNode = JsonNode.Parse("""
        {
            "Name": @name,
            "Age": 0,
            "Email": @email,
            "Score": 0.0
        }
    """); // Placeholder; reconstructing cleanly below

    int.TryParse(ageStr, out int age);
    float.Parse(scoreStr, CultureInfo.InvariantCulture); // Use explicit cast to avoid culture issues in parsing directly, though Parse is robust
    
    resultList.Add(new {
        Name = name,
        Age = age,
        Email = email,
        Score = float.Parse($"{scoreStr}", CultureInfo.InvariantCulture) 
    });
}

Console.OutputEncoding = Encoding.UTF8;
JsonSerializerOptions options = new JsonSerializerOptions { WriteIndented = false };
Console.WriteLine(JsonSerializer.Serialize(resultList, options));