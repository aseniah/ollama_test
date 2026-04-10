#r "System.Text.Json"
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputPath = "input/data.json";
var jsonContent = File.ReadAllText(inputPath);
var rootArray = JsonNode.Parse(jsonContent)![""] as JsonArray ?? new JsonArray();

// Parse and filter records
var filteredRecords = new List<JsonElement>();

foreach (var node in rootArray)
{
    if (node is not JsonObject obj) continue;

    var activeNode = obj["active"];
    var ageNode = obj["age"];

    // Check if active is true and age is >= 30
    if (activeNode != null && ageNode != null)
    {
        try
        {
            var isActive = activeNode.GetValueKind() == JsonValueKind.True 
                          ? (activeNode is JsonValue valT && valT.GetValue<bool>())
                          : (activeNode is JsonElement je && je.GetBoolean());

            // Robust age parsing
            int age;
            bool hasAge = false;
            
            if (ageNode is JsonValue valA)
            {
                try { age = valA.GetInt32(); hasAge = true; } catch { }
            }
            else if (ageNode is JsonElement jeAge && jeAge.TryGetInt32(out age))
            {
                hasAge = true;
            }

            if (hasAge && isActive && age >= 30)
            {
                filteredRecords.Add(obj);
            }
        }
        catch { continue; }
    }
}

// Sort by name ascending
var sortedRecords = filteredRecords
    .Select(j => j.EnumerateObject()
         .FirstOrDefault(o => o.Name == "name")?.Value.ToString() ?? string.Empty)
    .ToList(); // Get names for sorting reference if needed, but we need to sort the objects

// Proper sorting of the objects by name
var sortedJsonElements = filteredRecords
    .Select(j => new 
    { 
        Name = j.EnumerateObject().FirstOrDefault(o => o.Name == "name")?.Value.ToString() ?? string.Empty, 
        Item = j 
    })
    .OrderBy(x => x.Name)
    .Select(x => x.Item);

// Output result as a JSON array
var outputArray = new JsonArray();
foreach (var rec in sortedJsonElements)
{
    outputArray.Add(JsonNode.Parse(rec.GetRawText()));
}

Console.WriteLine(outputArray.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));