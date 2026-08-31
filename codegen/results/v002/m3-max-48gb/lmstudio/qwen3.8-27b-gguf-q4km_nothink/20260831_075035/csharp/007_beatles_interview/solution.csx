using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var referenceDate = new DateTime(2025, 7, 1);

var lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');
var dataLines = lines.Skip(1).ToList();

var people = new List<JsonObject>();

foreach (var line in dataLines)
{
    var fields = line.Split(',');
    var dict = new Dictionary<string, string>();
    for (int i = 0; i < header.Length && i < fields.Length; i++)
    {
        dict[header[i]] = fields[i];
    }

    // Parse full name
    var fullName = dict["Name"].Trim();
    var nameParts = fullName.Split(' ');
    string firstName, lastName;
    if (nameParts.Length == 2)
    {
        firstName = nameParts[0];
        lastName = nameParts[1];
    }
    else
    {
        firstName = nameParts[0];
        lastName = string.Join(" ", nameParts.Skip(1));
    }

    // Parse birthday
    var birthdayStr = dict["Birthday"].Trim();
    var birthdayParts = birthdayStr.Split('/');
    int month = int.Parse(birthdayParts[0]);
    int day = int.Parse(birthdayParts[1]);
    int year = int.Parse(birthdayParts[2]);
    var birthday = new DateTime(year, month, day);
    
    // Calculate age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.Month < birthday.Month || 
       (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
    {
        age--;
    }

    // Parse relatives
    var relatives = new ArrayNode();
    
    void AddRelative(string name, string relationship)
    {
        if (string.IsNullOrEmpty(name) || name == "null") return;
        var nameParts2 = name.Trim().Split(' ');
        string fName = nameParts2[0];
        string lName = nameParts2.Length > 1 ? string.Join(" ", nameParts2.Skip(1)) : "";
        var relObj = new JsonObject
        {
            ["FirstName"] = fName,
            ["LastName"] = lName,
            ["Relationship"] = relationship
        };
        relatives.Add(relObj);
    }

    AddRelative(dict["Father"], "Father");
    AddRelative(dict["Mother"], "Mother");
    AddRelative(dict["Brother"], "Brother");
    AddRelative(dict["Sister"], "Sister");

    var personObj = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    people.Add(personObj);
}

var arrayNode = new ArrayNode();
foreach (var p in people)
{
    arrayNode.Add(p);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.WriteLine(arrayNode.ToJsonString(options));