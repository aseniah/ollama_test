using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var refDate = new DateTime(2025, 7, 1);
var result = new JsonArray();

string[] GetNames(string n) => n.Trim().Split(' ', StringSplitOptions.RemoveEmptyEntries);

for (int i = 1; i < lines.Length; i++)
{
    var cols = lines[i].Split(',');
    if (cols.Length < 7) continue;
    
    var nameParts = GetNames(cols[0]);
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];
    
    var bday = DateTime.ParseExact(cols[1].Trim(), "M/d/yyyy", CultureInfo.InvariantCulture);
    int age = refDate.Year - bday.Year;
    if (refDate.Month < bday.Month || (refDate.Month == bday.Month && refDate.Day < bday.Day)) age--;
    
    var person = new JsonObject {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = bday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = new JsonArray()
    };

    var relatives = new List<(string First, string Last, string Rel)>();
    if (cols[3].Trim() != "null") 
    {
        var n = GetNames(cols[3]);
        relatives.Add((n[0], n[1], "Father"));
    }
    if (cols[4].Trim() != "null") 
    {
        var n = GetNames(cols[4]);
        relatives.Add((n[0], n[1], "Mother"));
    }
    if (cols[5].Trim() != "null") 
    {
        var n = GetNames(cols[5]);
        relatives.Add((n[0], n[1], "Brother"));
    }
    if (cols[6].Trim() != "null") 
    {
        var n = GetNames(cols[6]);
        relatives.Add((n[0], n[1], "Sister"));
    }

    foreach (var (f, l, rel) in relatives)
    {
        person["Relatives"]!.Add(new JsonObject {
            ["FirstName"] = f,
            ["LastName"] = l,
            ["Relationship"] = rel
        });
    }

    result.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = true }));