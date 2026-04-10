using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputPath = "input/data.json";
var content = File.ReadAllText(inputPath);
var doc = JsonDocument.Parse(content);
var records = doc.RootElement.ValueSequence()
    .Select(record => new {
        name = record.GetProperty("name").GetText(),
        age = int.Parse(record.GetProperty("age").GetText()),
        active = record.GetProperty("active").GetBool(),
        score = float.Parse(record.GetProperty("score").GetText())
    })
    .Where(r => r.active && r.age >= 30)
    .OrderBy(r => r.name);

var result = records.Select(r => new { name = r.name, age = r.age, active = r.active, score = r.score }).ToArray();

Console.WriteLine(JsonSerializer.Serialize(result));