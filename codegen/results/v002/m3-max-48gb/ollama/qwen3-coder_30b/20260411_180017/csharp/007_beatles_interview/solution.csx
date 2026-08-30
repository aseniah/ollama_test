using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var people = new List<JsonObject>();

var header = lines[0].Split(',');
for (int i = 1; i < lines.Length; i++)
{
    var values = lines[i].Split(',');
    var person = new JsonObject();
    
    // Parse name
    var nameParts = values[0].Split(' ', StringSplitOptions.RemoveEmptyEntries);
    person["FirstName"] = nameParts[0];
    person["LastName"] = string.Join(" ", nameParts.Skip(1));
    
    // Parse birthday
    var birthday = DateTime.Parse(values[1]);
    person["Birthday"] = birthday.ToString("yyyy-MM-dd");
    
    // Calculate age as of July 1, 2025
    var age = 2025 - birthday.Year;
    if (birthday.AddYears(age) > new DateTime(2025, 7, 1))
        age--;
    person["Age"] = age;
    
    // Parse relatives
    var relatives = new JsonArray();
    var relativesMap = new Dictionary<string, string> 
    { 
        { "Father", values[3] },
        { "Mother", values[4] },
        { "Brother", values[5] },
        { "Sister", values[6] }
    };
    
    foreach (var (relationship, relativeName) in relativesMap)
    {
        if (!string.IsNullOrEmpty(relativeName) && relativeName != "null")
        {
            var relative = new JsonObject();
            var relativeParts = relativeName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            relative["FirstName"] = relativeParts[0];
            relative["LastName"] = string.Join(" ", relativeParts.Skip(1));
            relative["Relationship"] = relationship;
            relatives.Add(relative);
        }
    }
    
    person["Relatives"] = relatives;
    people.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(people, new JsonSerializerOptions { WriteIndented = true }));