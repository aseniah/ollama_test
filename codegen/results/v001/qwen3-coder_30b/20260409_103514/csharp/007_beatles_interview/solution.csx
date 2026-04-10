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
    
    var person = new JsonObject
    {
        ["id"] = fields[0],
        ["name"] = fields[1],
        ["email"] = fields[2],
        ["birthDate"] = fields[3]
    };
    
    // Calculate age as of July 1, 2025
    if (DateTime.TryParse(fields[3], out DateTime birthDate))
    {
        var targetDate = new DateTime(2025, 7, 1);
        var age = targetDate.Year - birthDate.Year;
        
        // Adjust if birthday hasn't occurred yet this year
        if (birthDate.Date > targetDate.AddYears(-age))
            age--;
            
        person["age"] = age;
    }
    
    jsonArray.Add(person);
}

Console.WriteLine(jsonArray.ToJsonString());