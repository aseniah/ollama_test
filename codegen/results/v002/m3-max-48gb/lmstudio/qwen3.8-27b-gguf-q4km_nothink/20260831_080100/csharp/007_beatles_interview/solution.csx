using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;

string[] lines = File.ReadAllLines("input/input.csv");

DateTime referenceDate = new DateTime(2025, 7, 1);

List<JsonNode> result = new List<JsonNode>();

// Skip header
for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;
    
    string[] fields = line.Split(',');
    if (fields.Length < 7) continue;
    
    string fullName = fields[0].Trim();
    string birthdayStr = fields[1].Trim();
    string diedStr = fields[2].Trim();
    string father = fields[3].Trim();
    string mother = fields[4].Trim();
    string brother = fields[5].Trim();
    string sister = fields[6].Trim();
    
    // Parse name
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];
    
    // Parse birthday (M/D/YYYY format)
    DateTime birthday;
    try
    {
        birthday = DateTime.ParseExact(birthdayStr, "M/d/yyyy", System.Globalization.CultureInfo.InvariantCulture);
    }
    catch
    {
        birthday = DateTime.Parse(birthdayStr);
    }
    
    // Calculate age as of July 1, 2025
    int age;
    if (diedStr != "null" && !string.IsNullOrWhiteSpace(diedStr))
    {
        DateTime died;
        try
        {
            died = DateTime.ParseExact(diedStr, "M/d/yyyy", System.Globalization.CultureInfo.InvariantCulture);
        }
        catch
        {
            died = DateTime.Parse(diedStr);
        }
        
        DateTime ageRef = died < referenceDate ? died : referenceDate;
        age = ageRef.Year - birthday.Year;
        if ((ageRef.Month, ageRef.Day) < (birthday.Month, birthday.Day))
        {
            age--;
        }
    }
    else
    {
        age = referenceDate.Year - birthday.Year;
        if ((referenceDate.Month, referenceDate.Day) < (birthday.Month, birthday.Day))
        {
            age--;
        }
    }
    
    // Build relatives array
    List<JsonNode> relatives = new List<JsonNode>();
    
    if (father != "null" && !string.IsNullOrWhiteSpace(father))
    {
        string[] fatherParts = father.Split(' ');
        relatives.Add(JsonNode.Parse("{" +
            $"\"FirstName\":\"{fatherParts[0]}\"" +
            $"\"" + $"\"LastName\":\"{fatherParts[fatherParts.Length - 1]}\"" +
            $"" +
            $"\"" +
            $"\"Relationship\":\"Father\"" +
            " }"
        ));
    }
    
    if (mother != "null" && !string.IsNullOrWhiteSpace(mother))
    {
        string[] motherParts = mother.Split(' ');
        relatives.Add(JsonNode.Parse("{" +
            $"\"FirstName\":\"{motherParts[0]}\"" +
            $"\"" + $"\"LastName\":\"{motherParts[motherParts.Length - 1]}\"" +
            $"" +
            $"\"" +
            $"\"Relationship\":\"Mother\"" +
            " }"
        ));
    }
    
    if (brother != "null" && !string.IsNullOrWhiteSpace(brother))
    {
        string[] brotherParts = brother.Split(' ');
        relatives.Add(JsonNode.Parse("{" +
            $"\"FirstName\":\"{brotherParts[0]}\"" +
            $"\"" + $"\"LastName\":\"{brotherParts[brotherParts.Length - 1]}\"" +
            $"" +
            $"\"" +
            $"\"Relationship\":\"Brother\"" +
            " }"
        ));
    }
    
    if (sister != "null" && !string.IsNullOrWhiteSpace(sister))
    {
        string[] sisterParts = sister.Split(' ');
        relatives.Add(JsonNode.Parse("{" +
            $"\"FirstName\":\"{sisterParts[0]}\"" +
            $"\"" + $"\"LastName\":\"{sisterParts[sisterParts.Length - 1]}\"" +
            $"" +
            $"\"" +
            $"\"Relationship\":\"Sister\"" +
            " }"
        ));
    }
    
    // Format birthday as ISO
    string birthdayIso = birthday.ToString("yyyy-MM-dd");
    
    // Build person object
    JsonObject personObj = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayIso,
        ["Age"] = age,
        ["Relatives"] = new JsonArray(relatives.ToArray())
    };
    
    result.Add(personObj);
}

JsonArray output = new JsonArray(result.ToArray());
string json = output.ToJsonString(new JsonSerializerOptions 
{
    WriteIndented = true
});

Console.Write(json);