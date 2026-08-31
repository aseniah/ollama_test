using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string csvPath = "input/input.csv";
var lines = File.ReadAllLines(csvPath);

var referenceDate = new DateTime(2025, 7, 1);
var people = new List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;
    
    var parts = line.Split(',');
    // Name, Birthday, Died, Father, Mother, Brother, Sister
    string name = parts[0].Trim();
    string birthdayStr = parts[1].Trim();
    string diedStr = parts[2].Trim();
    string fatherStr = parts[3].Trim();
    string motherStr = parts[4].Trim();
    string brotherStr = parts[5].Trim();
    string sisterStr = parts[6].Trim();
    
    // Parse birthday - format is M/D/YYYY
    DateTime birthday;
    if (!DateTime.TryParseExact(birthdayStr, new[] { "M/d/yyyy", "M/d/yy", "MM/dd/yyyy" }, CultureInfo.InvariantCulture, DateTimeStyles.None, out birthday))
    {
        if (!DateTime.TryParse(birthdayStr, CultureInfo.InvariantCulture, DateTimeStyles.None, out birthday))
        {
            birthday = DateTime.MinValue;
        }
    }
    
    // Parse died date
    DateTime? diedDate = null;
    if (diedStr != "null" && diedStr != "")
    {
        if (DateTime.TryParseExact(diedStr, new[] { "M/d/yyyy", "M/d/yy", "MM/dd/yyyy" }, CultureInfo.InvariantCulture, DateTimeStyles.None, out var d))
        {
            diedDate = d;
        }
        else
        {
            DateTime.TryParse(diedStr, CultureInfo.InvariantCulture, DateTimeStyles.None, out d);
            diedDate = d;
        }
    }
    
    // Split name into FirstName and LastName
    var nameParts = name.Split(' ');
    string firstName = nameParts.Length > 0 ? nameParts[0] : "";
    string lastName = nameParts.Length > 1 ? nameParts[nameParts.Length - 1] : "";
    
    // Calculate age
    DateTime referenceDateForAge = diedDate.HasValue ? diedDate.Value : referenceDate;
    int age = referenceDateForAge.Year - birthday.Year;
    if (referenceDateForAge.Month < birthday.Month || 
        (referenceDateForAge.Month == birthday.Month && referenceDateForAge.Day < birthday.Day))
    {
        age--;
    }
    
    // Build relatives
    var relatives = new List<JsonNode>();
    
    void AddRelative(string nameStr, string relationship)
    {
        if (nameStr == "null" || nameStr == "") return;
        var np = nameStr.Split(' ');
        string fn = np.Length > 0 ? np[0] : "";
        string ln = np.Length > 1 ? np[np.Length - 1] : "";
        relatives.Add(JsonNode.Parse($"{{\"FirstName\":\"{fn}\",\"LastName\":\"{ln}\",\"Relationship\":\"{relationship}\"}}"));
    }
    
    AddRelative(fatherStr, "Father");
    AddRelative(motherStr, "Mother");
    AddRelative(brotherStr, "Brother");
    AddRelative(sisterStr, "Sister");
    
    var relativeArray = new JsonArray();
    foreach (var r in relatives)
    {
        relativeArray.Add(r);
    }
    
    var personObj = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthday.ToString("yyyy-MM-dd"),
        ["Age"] = age,
        ["Relatives"] = relativeArray
    };
    
    people.Add(personObj);
}

var jsonArray = new JsonArray();
foreach (var p in people)
{
    jsonArray.Add(p);
}

Console.WriteLine(jsonArray.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));