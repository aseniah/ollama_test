using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonContent = File.ReadAllText("input/data.json");
var array = JsonArray.Parse(jsonContent);

var filtered = array
    .Cast<JsonObject>()
    .Where(obj => obj["active"]?.GetValue<bool>() == true && obj["age"]?.GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"]?.GetValue<string>(), StringComparer.Ordinal)
    .ToArray();

var result = new JsonArray(filtered);
var output = result.ToJsonString(new JsonSerializerOptions { WriteIndented = true });
Console.Write(output);