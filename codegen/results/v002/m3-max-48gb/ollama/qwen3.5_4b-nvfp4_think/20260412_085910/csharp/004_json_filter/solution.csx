using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var jsonContent = File.ReadAllText("input/data.json");
var jsonDoc = JsonDocument.Parse(jsonContent);

var records = new List<JsonNode>();

foreach (var item in jsonDoc.RootElement.EnumerateArray().Select(j => j))
{
    var name = ((string)item["name"]);
    var age = ((int)item["age"]);
    var active = ((bool)item["active"]);

    if (active && age >= 30)
    {
        records.Add(item);
    }
}

var sortedRecords = records.OrderBy(r => (string)r["name"]).ToList();

var outputArray = new List<JsonNode>();
foreach (var item in sortedRecords)
{
    var outputItem = JsonDocument.Parse(@"{"name": """ + item["name"] + @", age": " + item["age"] + @" , active": " + ((bool)item["active"]) + @", score": " + item["score"] + @"}")
                .Value;
    outputArray.Add(outputItem);
}

var outputJson = new JsonArray();
foreach (var item in outputArray)
{
    outputArray.Add(item.ToString());
}

Console.Write(outputArray.ToArray().ToString());