using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/input.csv");
var result = new List<JsonNode>();
DateTime targetDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;

    string[] parts = line.Split(',');
    if (parts.Length < 6) continue;

    string fullName = parts[0].Trim();
    string birthdayStr = parts[1].Trim();
    // parts[2] is Died, unused for calculation but present
    string father = parts[3].Trim();
    string mother = parts[4].Trim();
    string brother = parts[5].Trim();
    string sister = parts.Length > 6 ? parts[6].Trim() : "";

    // Parse Birthday (M/D/YYYY)
    DateTime birthday;
    if (!DateTime.TryParse(birthdayStr, out birthday)) continue;

    // Calculate Age as of July 1, 2025
    int age = targetDate.Year - birthday.Year;
    if (targetDate.DayOfYear < birthday.DayOfYear)
    {
        age--;
    }

    // Split Full Name into First and Last
    string[] nameParts = fullName.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts.Length > 0 ? nameParts[0] : "";
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";

    // Build Relatives List
    var relatives = new List<JsonNode>();
    if (!string.IsNullOrEmpty(father) && father != "null")
    {
        var fParts = father.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        string fName = fParts.Length > 0 ? fParts[0] : "";
        string fLast = fParts.Length > 1 ? fParts[fParts.Length - 1] : "";
        relatives.Add(new JsonNode.Parse($"{{\"FirstName\":\"{fName}\",\"LastName\":\"{fLast}\",\"Relationship\":\"Father\"}}"));
    }
    if (!string.IsNullOrEmpty(mother) && mother != "null")
    {
        var mParts = mother.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        string mName = mParts.Length > 0 ? mParts[0] : "";
        string mLast = mParts.Length > 1 ? mParts[mParts.Length - 1] : "";
        relatives.Add(new JsonNode.Parse($"{{\"FirstName\":\"{mName}\",\"LastName\":\"{mLast}\",\"Relationship\":\"Mother\"}}"));
    }
    if (!string.IsNullOrEmpty(brother) && brother != "null")
    {
        var bParts = brother.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        string bName = bParts.Length > 0 ? bParts[0] : "";
        string bLast = bParts.Length > 1 ? bParts[bParts.Length - 1] : "";
        relatives.Add(new JsonNode.Parse($"{{\"FirstName\":\"{bName}\",\"LastName\":\"{bLast}\",\"Relationship\":\"Brother\"}}"));
    }
    if (!string.IsNullOrEmpty(sister) && sister != "null")
    {
        var sParts = sister.Split(' ', StringSplitOptions.RemoveEmptyEntries);
        string sName = sParts.Length > 0 ? sParts[0] : "";
        string sLast = sParts.Length > 1 ? sParts[sParts.Length - 1] : "";
        relatives.Add(new JsonNode.Parse($"{{\"FirstName\":\"{sName}\",\"LastName\":\"{sLast}\",\"Relationship\":\"Sister\"}}"));
    }

    var person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = JsonArray.Parse($"[{string.Join(",", relatives.ConvertAll(r => r.ToString()))}]")
    };

    // Re-parse to ensure clean serialization without extra quotes in the array construction above if needed, 
    // but simpler is to build the JsonObject directly for the parent.
    
    var relArray = new JsonArray();
    foreach (var rel in relatives) { relArray.Add(rel); }
    person["Relatives"] = relArray;

    result.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = true }));