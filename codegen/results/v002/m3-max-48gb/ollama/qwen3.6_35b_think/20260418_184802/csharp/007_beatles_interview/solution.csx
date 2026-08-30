using System;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/input.csv");
JsonArray result = new JsonArray();

DateTime refDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    string[] parts = line.Split(',').Select(p => p.Trim()).ToArray();
    string fullName = parts[0];
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];
    
    string birthdayStr = parts[1];
    DateTime birthday = DateTime.ParseExact(birthdayStr, new[] { "M/d/yyyy", "MM/dd/yyyy" }, CultureInfo.InvariantCulture);
    
    string diedStr = parts[2];
    DateTime? died = null;
    if (!string.IsNullOrEmpty(diedStr) && diedStr != "null")
    {
        died = DateTime.ParseExact(diedStr, new[] { "M/d/yyyy", "MM/dd/yyyy" }, CultureInfo.InvariantCulture);
    }
    
    DateTime calcDate = died ?? refDate;
    int age = calcDate.Year - birthday.Year;
    if (calcDate.Month < birthday.Month || (calcDate.Month == birthday.Month && calcDate.Day < birthday.Day))
    {
        age--;
    }
    
    JsonArray relatives = new JsonArray();
    string[] relTypes = { "Father", "Mother", "Brother", "Sister" };
    
    for (int j = 0; j < 4; j++)
    {
        string relName = parts[3 + j];
        if (!string.IsNullOrEmpty(relName) && relName != "null")
        {
            string[] relParts = relName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
            if (relParts.Length > 0)
            {
                var relObj = new JsonObject
                {
                    ["FirstName"] = relParts[0],
                    ["LastName"] = relParts[relParts.Length - 1],
                    ["Relationship"] = relTypes[j]
                };
                relatives.Add(relObj);
            }
        }
    }
    
    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = relatives
    };
    result.Add(person);
}

var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(result, options));