using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;
using System.IO;

var data = File.ReadAllText("input/data.json");
var doc = JsonDocument.Parse(data);
var root = doc.RootElement;

var filtered = new List<JsonElement>();
foreach (var element in root.EnumerateArrayElements())
{
    var age = element.GetProperty("age").GetInt32();
    var active = element.GetProperty("active").GetBoolean();
    if (active && age >= 30)
    {
        filtered.Add(element);
    }
}

filtered.Sort((a, b) => a.GetProperty("name").GetString().CompareTo(b.GetProperty("name").GetString()));

Console.WriteLine(JsonSerializer.Serialize(filtered));