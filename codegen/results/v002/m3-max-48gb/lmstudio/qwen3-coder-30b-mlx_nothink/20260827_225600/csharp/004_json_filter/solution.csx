using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonData = File.ReadAllText("input/data.json");
var jsonArray = JsonNode.Parse(jsonData).AsArray();

var filteredRecords = jsonArray
    .Where(obj => obj["active"].GetValue<bool>() && obj["age"].GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"].GetValue<string>())
    .ToArray();

Console.WriteLine(JsonSerializer.Serialize(filteredRecords));