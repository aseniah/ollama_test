using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string jsonContent = File.ReadAllText("input/data.json");
JsonDocument jsonDoc = JsonDocument.Parse(jsonContent);
JsonArray dataArray = jsonDoc.RootElement.GetProperty("elements"); // Assuming the root is an object with an "elements" array, but let's handle if the root is the array directly.
// Re-evaluating: The input example shows the root of the JSON is the array itself.
// So dataArray = jsonDoc.RootElement.AsArray();

JsonArray dataArray = jsonDoc.RootElement.AsArray();

List<dynamic> filteredRecords = new List<dynamic>();

foreach (JsonElement element in dataArray)
{
    if (element.TryGetProperty("name", out JsonProperty nameProp) &&
        element.TryGetProperty("age", out JsonProperty ageProp) &&
        element.TryGetProperty("active", out JsonProperty activeProp) &&
        element.TryGetProperty("score", out JsonProperty scoreProp))
    {
        string name = nameProp.GetString()!;
        int age = ageProp.GetInt32();
        bool active = activeProp.GetBoolean();
        double score = scoreProp.GetDouble();

        if (active && age >= 30)
        {
            filteredRecords.Add(new
            {
                Name = name,
                Age = age,
                Active = active,
                Score = score
            });
        }
    }
}

var sortedRecords = filteredRecords.OrderBy(r => r.Name).ToList();

string output = JsonNode.Serialize(sortedRecords);
Console.WriteLine(output);