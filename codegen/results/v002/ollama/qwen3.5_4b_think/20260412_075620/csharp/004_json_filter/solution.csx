using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");

var document = JsonDocument.Parse(json);
var root = document.RootElement;

var records = new List<object>();

if (root.ValueKind == JsonValueKind.Array) {
    foreach (var element in root.EnumerateArray()) {
        var name = element.GetProperty("name").GetString();
        var age = element.GetProperty("age").GetInt32();
        var active = element.GetProperty("active").GetBoolean();
        var score = element.GetProperty("score").GetDouble();

        if (active && age >= 30) {
            records.Add(new { name, age, active, score });
        }
    }
}

records = records.OrderBy(x => ((dynamic)x).name).ToList();

var options = new JsonSerializerOptions();
var output = JsonSerializer.Serialize(records, options);

Console.WriteLine(output);