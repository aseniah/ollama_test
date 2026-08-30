using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string jsonContent = File.ReadAllText("input/data.json");
JsonArray jsonArray = (JsonArray)JsonNode.Parse(jsonContent);

var filtered = jsonArray
    .Select(node => (JsonObject)node)
    .Where(obj => (bool)obj["active"] == true && (int)obj["age"] >= 30)
    .OrderBy(obj => (string)obj["name"])
    .ToList();

JsonArray resultArray = new JsonArray();
foreach (var obj in filtered)
{
    resultArray.Add(obj);
}

string output = resultArray.ToJsonString(new JsonSerializerOptions { WriteIndented = false });
Console.Write(output);