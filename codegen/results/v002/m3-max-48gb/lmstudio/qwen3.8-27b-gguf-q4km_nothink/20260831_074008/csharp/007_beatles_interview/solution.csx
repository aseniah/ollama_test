using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

var lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');

var people = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    var name = parts[0];
    
    // Split name into first and last
    var nameParts = name.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];
    
    // Parse birthday
    var birthdayParts = parts[1].Split('/');
    var birthDate = new DateTime(int.Parse(birthdayParts[2]), int.Parse(birthdayParts[0]), int.Parse(birthdayParts[1]));
    
    // Calculate age as of July 1, 2025
    var referenceDate = new DateTime(2025, 7, 1);
    int age = referenceDate.Year - birthDate.Year;
    if (referenceDate.Month < birthDate.Month || (referenceDate.Month == birthDate.Month && referenceDate.Day < birthDate.Day))
    {
        age--;
    }
    
    // Format birthday as ISO 8601
    var birthdayStr = birthDate.ToString("yyyy-MM-dd");
    
    var relatives = new List<Dictionary<string, string>>();
    
    // Father
    if (parts[3] != "null")
    {
        var fParts = parts[3].Split(' ');
        relatives.Add(new Dictionary<string, string>
        {
            ["FirstName"] = fParts[0],
            ["LastName"] = fParts[fParts.Length - 1],
            ["Relationship"] = "Father"
        });
    }
    
    // Mother
    if (parts[4] != "null")
    {
        var mParts = parts[4].Split(' ');
        relatives.Add(new Dictionary<string, string>
        {
            ["FirstName"] = mParts[0],
            ["LastName"] = mParts[mParts.Length - 1],
            ["Relationship"] = "Mother"
        });
    }
    
    // Brother
    if (parts[5] != "null")
    {
        var bParts = parts[5].Split(' ');
        relatives.Add(new Dictionary<string, string>
        {
            ["FirstName"] = bParts[0],
            ["LastName"] = bParts[bParts.Length - 1],
            ["Relationship"] = "Brother"
        });
    }
    
    // Sister
    if (parts[6] != "null")
    {
        var sParts = parts[6].Split(' ');
        relatives.Add(new Dictionary<string, string>
        {
            ["FirstName"] = sParts[0],
            ["LastName"] = sParts[sParts.Length - 1],
            ["Relationship"] = "Sister"
        });
    }
    
    var person = new Dictionary<string, object>
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayStr,
        ["Age"] = age,
        ["Relatives"] = relatives
    };
    
    people.Add(person);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

var json = JsonSerializer.Serialize(people, options);
Console.WriteLine(json);