using System;
using System.IO;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read the CSV file
var content = File.ReadAllText("input/data.csv");

// Parse lines, handle both Windows and Unix line endings
var lines = content.Split(new[] { "\r\n", "\n" }, StringSplitOptions.None);

// Skip the header row
if (lines.Length > 0)
{
    lines = lines.Skip(1).ToList();
}

// Create JSON array
var jsonArray = new List<JsonElement>();

// Process each line
foreach (var line in lines)
{
    if (string.IsNullOrWhiteSpace(line))
        continue;

    var values = line.Split(',');
    
    // Validate we have at least 4 columns
    if (values.Length < 4)
        continue;

    var obj = new JsonObject();
    obj["Name"] = values[0];
    obj["Age"] = int.Parse(values[1]);
    obj["Email"] = values[2];
    obj["Score"] = float.Parse(values[3]);
    
    jsonArray.Add(obj);
}

// Output as JSON array
Console.WriteLine(JsonArrayToString(jsonArray));