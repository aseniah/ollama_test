using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var filePath = Args[0];
var lines = File.ReadAllLines(filePath);

var jsonDoc = JsonDocument.Parse(new string(lines).TrimEnd('\r\n')).Value;

// Parse the array and filter/sort records
var filteredRecords = new List<JsonElement>();
foreach (var doc in jsonDoc.EnumerateArray())
{
    var nameNode = doc["name"];
    var ageNode = doc["age"];
    var activeNode = doc["active"];

    var name = nameNode.GetValueKind() == JsonValueKind.String ? (string)nameNode : null;
    var age = int.TryParse(ageNode.ToString() out var a, out _) ? a : 0;
    var active = bool.TryParse(activeNode.ToString() out var b, out _) && b;

    if (active && age >= 30)
    {
        filteredRecords.Add(doc);
    }
}

// Sort by name ascending
filteredRecords.sort((x, y) => string.Compare(x[1]["name"].GetValueKind() == JsonValueKind.String ? (string)x[1]["name"] : null, y[1]["name"].GetValueKind() == JsonValueKind.String ? (string)y[1]["name"] : null, StringComparison.Ordinal));

var jsonArray = new List<JsonElement>();
foreach (var doc in filteredRecords)
{
    jsonArray.Add(doc);
}

Console.WriteLine(JsonSerializer.Serialize(jsonArray, options => {
    options.WriteIndented = true;
    options.DiswriteNullValues = false;
    options.UseStringTwiceEncodingForDateTime = false;
    options.ErrorWhenCaseMismatch = false;
    options.IgnoreUnreadableCharacters = true;
}));