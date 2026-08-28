using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string csvPath = "input/input.csv";
string[] lines = File.ReadAllLines(csvPath);
string header = lines[0].Split(',');
int nameIdx = Array.IndexOf(header, "Name");
int birthdayIdx = Array.IndexOf(header, "Birthday");
int diedIdx = Array.IndexOf(header, "Died");
int fatherIdx = Array.IndexOf(header, "Father");
int motherIdx = Array.IndexOf(header, "Mother");
int brotherIdx = Array.IndexOf(header, "Brother");
int sisterIdx = Array.IndexOf(header, "Sister");

DateTime referenceDate = new DateTime(2025, 7, 1);

var results = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    string[] fields = lines[i].Split(',');
    
    // Parse name
    string fullName = fields[nameIdx].Trim();
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];
    
    // Parse birthday: M/D/YYYY
    string[] bd = fields[birthdayIdx].Trim().Split('/');
    int bMonth = int.Parse(bd[0]);
    int bDay = int.Parse(bd[1]);
    int bYear = int.Parse(bd[2]);
    string birthdayStr = $"{bYear:0000}-{bMonth:00}-{bDay:00}";
    
    // Calculate age
    int age;
    string diedStr = fields[diedIdx].Trim();
    if (diedStr != "null" && !string.IsNullOrEmpty(diedStr))
    {
        // Deceased: DeathYear - BirthYear
        string[] dParts = diedStr.Split('/');
        int dYear = int.Parse(dParts[2]);
        age = dYear - bYear;
    }
    else
    {
        // Alive: as of July 1, 2025
        age = 2025 - bYear;
        if (bMonth > 7 || (bMonth == 7 && bDay > 1))
        {
            age -= 1;
        }
    }
    
    // Parse relatives
    var relatives = new JsonArray();
    
    void AddRelative(string nameField, string relationship)
    {
        string name = nameField.Trim();
        if (name == "null" || string.IsNullOrEmpty(name)) return;
        string[] parts = name.Split(' ');
        var rel = new JsonObject
        {
            ["FirstName"] = parts[0],
            ["LastName"] = parts[parts.Length - 1],
            ["Relationship"] = relationship
        };
        relatives.Add(rel);
    }
    
    AddRelative(fields[fatherIdx], "Father");
    AddRelative(fields[motherIdx], "Mother");
    AddRelative(fields[brotherIdx], "Brother");
    AddRelative(fields[sisterIdx], "Sister");
    
    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayStr,
        ["Age"] = age,
        ["Relatives"] = relatives
    };
    
    results.Add(person);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(results, options));