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
    person["LastName"] = nameParts.Length > 1 ? string.Join(" ", nameParts.Skip(1)) : "";
    
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
    
    // Father
    if (!string.IsNullOrEmpty(values[3]))
    {
        var father = new JsonObject();
        var fatherParts = values[3].Split(' ', StringSplitOptions.RemoveEmptyEntries);
        father["FirstName"] = fatherParts[0];
        father["LastName"] = fatherParts.Length > 1 ? string.Join(" ", fatherParts.Skip(1)) : "";
        father["Relationship"] = "Father";
        relatives.Add(father);
    }
    
    // Mother
    if (!string.IsNullOrEmpty(values[4]))
    {
        var mother = new JsonObject();
        var motherParts = values[4].Split(' ', StringSplitOptions.RemoveEmptyEntries);
        mother["FirstName"] = motherParts[0];
        mother["LastName"] = motherParts.Length > 1 ? string.Join(" ", motherParts.Skip(1)) : "";
        mother["Relationship"] = "Mother";
        relatives.Add(mother);
    }
    
    // Brother
    if (!string.IsNullOrEmpty(values[5]) && values[5] != "null")
    {
        var brother = new JsonObject();
        var brotherParts = values[5].Split(' ', StringSplitOptions.RemoveEmptyEntries);
        brother["FirstName"] = brotherParts[0];
        brother["LastName"] = brotherParts.Length > 1 ? string.Join(" ", brotherParts.Skip(1)) : "";
        brother["Relationship"] = "Brother";
        relatives.Add(brother);
    }
    
    // Sister
    if (!string.IsNullOrEmpty(values[6]) && values[6] != "null")
    {
        var sister = new JsonObject();
        var sisterParts = values[6].Split(' ', StringSplitOptions.RemoveEmptyEntries);
        sister["FirstName"] = sisterParts[0];
        sister["LastName"] = sisterParts.Length > 1 ? string.Join(" ", sisterParts.Skip(1)) : "";
        sister["Relationship"] = "Sister";
        relatives.Add(sister);
    }
    
    person["Relatives"] = relatives;
    people.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(people, new JsonSerializerOptions { WriteIndented = true }));