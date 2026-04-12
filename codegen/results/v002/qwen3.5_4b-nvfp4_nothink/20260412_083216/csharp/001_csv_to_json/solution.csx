using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
if (lines.Length == 0) {
    throw new FileNotFoundException("input/data.csv not found");
}

var header = lines[0].Split(',');
var records = new List<Dictionary<string, string>>();

for (int i = 1; i < lines.Length; i++) {
    var row = lines[i].Split(',');
    var dict = new Dictionary<string, string>();
    for (int j = 0; j < header.Length; j++) {
        dict[header[j]] = row[j];
    }
    records.Add(dict);
}

var outputList = new List<JsonElement>();

foreach (var record in records) {
    var obj = JsonArray.Parse("{}");
    if (!record.TryGetValue("Name", out var name)) throw new InvalidOperationException("Missing Name");
    if (!record.TryGetValue("Age", out var age)) throw new InvalidOperationException("Missing Age");
    if (!record.TryGetValue("Email", out var email)) throw new InvalidOperationException("Missing Email");
    if (!record.TryGetValue("Score", out var score)) throw new InvalidOperationException("Missing Score");

    obj["Name"] = name.ValueToString();
    obj["Age"] = int.Parse(age);
    obj["Email"] = email.ValueToString();
    obj["Score"] = float.Parse(score);
    
    outputList.Add(obj);
}

System.Console.Write(JsonSerializer.Serialize(outputList));