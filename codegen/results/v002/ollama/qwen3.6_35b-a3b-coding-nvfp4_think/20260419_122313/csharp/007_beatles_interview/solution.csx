using System;
using System.IO;
using System.Globalization;
using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/input.csv");
var result = new JsonArray();

DateTime reference = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    string[] cols = line.Split(',');
    
    string name = cols[0].Trim();
    string[] nameParts = name.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];
    
    string birthdayStr = cols[1].Trim();
    DateTime birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", CultureInfo.InvariantCulture);
    string formattedBirthday = birthday.ToString("yyyy-MM-dd");
    
    int age = reference.Year - birthday.Year;
    if (reference.Month < birthday.Month || (reference.Month == birthday.Month && reference.Day < birthday.Day))
        age--;
        
    var relatives = new JsonArray();
    string[] relationships = { "Father", "Mother", "Brother", "Sister" };
    string[] relativeCols = { cols[3].Trim(), cols[4].Trim(), cols[5].Trim(), cols[6].Trim() };
    
    for (int r = 0; r < relationships.Length; r++)
    {
        string relName = relativeCols[r];
        if (relName != "null" && !string.IsNullOrEmpty(relName))
        {
            string[] relParts = relName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            relatives.Add(new JsonObject
            {
                ["FirstName"] = relParts[0],
                ["LastName"] = relParts[relParts.Length - 1],
                ["Relationship"] = relationships[r]
            });
        }
    }
    
    result.Add(new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = formattedBirthday,
        ["Age"] = age,
        ["Relatives"] = relatives
    });
}

Console.WriteLine(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = true }));