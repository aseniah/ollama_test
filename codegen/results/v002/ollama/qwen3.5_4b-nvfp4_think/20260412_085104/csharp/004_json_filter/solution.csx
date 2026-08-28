using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.IO;
using System.Collections.Generic;

var jsonContent = File.ReadAllText("input/data.json");
var documents = JsonDocument.Parse(jsonContent).EnumerateArray().ToList();

var filteredRecords = documents.Where(r => 
    (bool)r["active"] == true && 
    ((int)r["age"]) >= 30)
    .OrderBy(r => r["name"].GetValue<string>())
    .ToList();

var results = new List<JsonElement>();
foreach(var doc in filteredRecords) {
    var obj = new JsonObject(doc);
    results.Add(obj);
}

Console.WriteLine(JsonSerializer.Serialize(results));