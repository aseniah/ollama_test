using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string csvPath = "input/data.csv";
string content = File.ReadAllText(csvPath);
string[] lines = content.Split(new[] { "\r\n", "\n" }, StringSplitOptions.RemoveEmptyEntries);

// Skip header (line 0)
List<JsonNode> rows = new List<JsonNode>();

for (int i = 1; i < lines.Length; i++)
{
    string[] parts = lines[i].Split(',');
    
    if (parts.Length < 4) continue;

    // Trim and parse values
    string name = parts[0].Trim();
    int age = int.Parse(parts[1].Trim());
    string email = parts[2].Trim();
    float score = float.Parse(parts[3].Trim(), System.Globalization.CultureInfo.InvariantCulture);

    JsonNode row = new JsonObject { ["Name"] = name, ["Age"] = age, ["Email"] = email, ["Score"] = score };
    rows.Add(row);
}

var jsonArray = new JsonArray(rows);
string jsonOutput = JsonSerializer.Serialize(jsonArray);

Console.WriteLine(jsonOutput);