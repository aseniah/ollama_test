using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string csvPath = "input/input.csv";
string[] lines = File.ReadAllLines(csvPath);

// Parse header
string[] headers = lines[0].Split(',');

// Build list of persons
var persons = new List<JsonObject>();

foreach (string line in lines.Skip(1))
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    string[] fields = line.Split(',');
    
    string fullName = fields[0].Trim();
    string birthdayStr = fields[1].Trim();
    string diedStr = fields[2].Trim();
    string fatherStr = fields[3].Trim();
    string motherStr = fields[4].Trim();
    string brotherStr = fields[5].Trim();
    string sisterStr = fields[6].Trim();
    
    // Parse name
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];
    
    // Parse birthday M/D/YYYY -> YYYY-MM-DD
    DateTime birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", System.Globalization.CultureInfo.InvariantCulture);
    string birthdayFormatted = birthday.ToString("yyyy-MM-dd");
    
    // Calculate age
    int age;
    if (diedStr != "null" && !string.IsNullOrWhiteSpace(diedStr))
    {
        DateTime died = DateTime.ParseExact(diedStr, "M/d/yyyy", System.Globalization.CultureInfo.InvariantCulture);
        // Age at death
        age = died.Year - birthday.Year;
        if (died.Month < birthday.Month || (died.Month == birthday.Month && died.Day < birthday.Day))
        {
            age--;
        }
    }
    else
    {
        // Age as of July 1, 2025
        DateTime referenceDate = new DateTime(2025, 7, 1);
        age = referenceDate.Year - birthday.Year;
        if (referenceDate.Month < birthday.Month || (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
        {
            age--;
        }
    }
    
    // Build relatives
    var relatives = new JsonArray();
    
    void AddRelative(string nameStr, string relationship)
    {
        if (nameStr == "null" || string.IsNullOrWhiteSpace(nameStr)) return;
        string[] relNameParts = nameStr.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        string relFirstName = relNameParts[0];
        string relLastName = relNameParts[relNameParts.Length - 1];
        
        relatives.Add(new JsonObject
        {
            ["FirstName"] = relFirstName,
            ["LastName"] = relLastName,
            ["Relationship"] = relationship
        });
    }
    
    AddRelative(fatherStr, "Father");
    AddRelative(motherStr, "Mother");
    AddRelative(brotherStr, "Brother");
    AddRelative(sisterStr, "Sister");
    
    var personObj = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = relatives
    };
    
    persons.Add(personObj);
}

var result = new JsonArray(persons);
string jsonOutput = result.ToJsonString(new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(jsonOutput);