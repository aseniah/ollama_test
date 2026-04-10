#r "System.Text.Json"
#r "System.IO"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the input CSV
string[] lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');

// Create a list to hold the output JSON objects
var result = new List<object>();

// Reference date for age calculation
var referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    string[] values = lines[i].Split(',');
    
    // Basic assumption: CSV structure based on common patterns and expected JSON
    // Typically: Name, DateOfBirth, City, etc.
    // Adjust indices based on actual input.csv header if known, but inferring from expected_format.json
    // Assuming: Name, DateOfBirth, City, State, Zip
    
    var name = values.Length > 0 ? values[0].Trim() : "";
    var dobStr = values.Length > 1 ? values[1].Trim() : "";
    var city = values.Length > 2 ? values[2].Trim() : "";
    var state = values.Length > 3 ? values[3].Trim() : "";
    var zip = values.Length > 4 ? values[4].Trim() : "";

    int age = 0;
    if (!string.IsNullOrEmpty(dobStr) && DateTime.TryParse(dobStr, out DateTime dob))
    {
        age = referenceDate.Year - dob.Year;
        if (referenceDate.DayOfYear < dob.DayOfYear)
        {
            age--;
        }
    }

    var address = new JsonObject();
    address["city"] = city;
    address["state"] = state;
    address["zip"] = zip;

    var person = new JsonObject();
    person["name"] = name;
    person["age"] = age;
    person["address"] = address;

    result.Add(person);
}

var jsonOutput = JsonSerializer.Serialize(result, new JsonSerializerOptions 
{ 
    WriteIndented = false, 
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase 
});

Console.Write(jsonOutput);