using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var people = new List<JsonObject>();

// Skip header line
for (int i = 1; i < lines.Length; i++)
{
    var fields = lines[i].Split(',');
    
    var person = new JsonObject
    {
        ["FirstName"] = fields[0].Split(' ')[0],
        ["LastName"] = fields[0].Split(' ')[^1],
        ["Birthday"] = ConvertDate(fields[1]),
        ["Age"] = CalculateAge(fields[1], "2025-07-01"),
        ["Relatives"] = new JsonArray()
    };
    
    var relatives = person["Relatives"].AsArray();
    
    // Add father
    if (!string.IsNullOrEmpty(fields[3]))
    {
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fields[3].Split(' ')[0],
            ["LastName"] = fields[3].Split(' ')[^1],
            ["Relationship"] = "Father"
        });
    }
    
    // Add mother
    if (!string.IsNullOrEmpty(fields[4]))
    {
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fields[4].Split(' ')[0],
            ["LastName"] = fields[4].Split(' ')[^1],
            ["Relationship"] = "Mother"
        });
    }
    
    // Add brother
    if (!string.IsNullOrEmpty(fields[5]) && fields[5] != "null")
    {
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fields[5].Split(' ')[0],
            ["LastName"] = fields[5].Split(' ')[^1],
            ["Relationship"] = "Brother"
        });
    }
    
    // Add sister
    if (!string.IsNullOrEmpty(fields[6]) && fields[6] != "null")
    {
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fields[6].Split(' ')[0],
            ["LastName"] = fields[6].Split(' ')[^1],
            ["Relationship"] = "Sister"
        });
    }
    
    people.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(people, new JsonSerializerOptions { WriteIndented = true }));

string ConvertDate(string date)
{
    var parts = date.Split('/');
    var month = int.Parse(parts[0]);
    var day = int.Parse(parts[1]);
    var year = int.Parse(parts[2]);
    
    return $"{year}-{month:D2}-{day:D2}";
}

int CalculateAge(string birthDate, string referenceDate)
{
    var birth = DateTime.Parse(ConvertDate(birthDate));
    var reference = DateTime.Parse(referenceDate);
    
    var age = reference.Year - birth.Year;
    
    if (reference.Month < birth.Month || (reference.Month == birth.Month && reference.Day < birth.Day))
    {
        age--;
    }
    
    return age;
}