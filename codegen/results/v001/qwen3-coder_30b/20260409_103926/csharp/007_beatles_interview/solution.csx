using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var jsonArray = new JsonArray();

// Skip header line
for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    
    // Parse date of birth (assuming format is MM/dd/yyyy)
    var dob = DateTime.ParseExact(fields[2], "MM/dd/yyyy", CultureInfo.InvariantCulture);
    
    // Calculate age as of July 1, 2025
    var age = 2025 - dob.Year;
    if (dob.Date > new DateTime(2025, 7, 1))
        age--;
    
    var obj = new JsonObject
    {
        ["id"] = int.Parse(fields[0]),
        ["name"] = fields[1],
        ["date_of_birth"] = dob.ToString("yyyy-MM-dd"),
        ["age"] = age,
        ["email"] = fields[3],
        ["phone"] = fields[4]
    };
    
    jsonArray.Add(obj);
}

Console.WriteLine(jsonArray.ToJsonString());