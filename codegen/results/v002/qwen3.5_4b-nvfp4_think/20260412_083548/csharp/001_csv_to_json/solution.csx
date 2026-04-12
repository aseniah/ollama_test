using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
if (lines.Length < 2) {
    Console.WriteLine("[]");
    return;
}

var output = new List<JsonNode>();
for (int i = 1; i < lines.Length; i++) {
    var row = lines[i].Split(',');
    if (row.Length != 4) continue;
    
    var obj = JsonNode.Parse("{}") as JObject;
    obj["Name"] = row[0];
    obj["Age"] = int.Parse(row[1]);
    obj["Email"] = row[2];
    obj["Score"] = float.Parse(row[3]);
    output.Add(obj);
}

var jsonDocument = new JsonArray(output.ToArray());
Console.WriteLine(JsonSerializer.Serialize(jsonDocument, new JsonSerializerOptions { WriteIndented = false }));