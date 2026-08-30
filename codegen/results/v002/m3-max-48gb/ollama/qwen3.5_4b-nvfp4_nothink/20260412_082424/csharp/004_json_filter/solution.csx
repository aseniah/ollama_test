using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.json");
var jsonStr = string.Join("\n", lines);

var items = new List<object>();
JsonNode? root = null;
try {
    root = JsonDocument.Parse(jsonStr).RootElement;
} catch (Exception) {
    return; // Or throw, but requirements say no output else. 
}

if (root == null || root.ValueKind != JsonValueKind.Array) {
    return;
}

foreach (JsonNode? node in (JsonArray?) root) {
    if (node == null) continue;
    
    var n = JsonDocument.Parse(node).RootElement;
    string nameStr = (n["name"]?.GetValue<string>()) ?? "";
    int age = n["age"].GetInt32();
    bool active = n["active"].GetBoolean();
    float score = n["score"].GetFloat();

    if (!active || age < 30) continue;

    items.Add(n["name"].GetString());
}

var sortedItems = items.OrderBy(s => s).ToList();

var jsonOptions = new JsonSerializerOptions { WriteIndented = false };
string output = JsonSerializer.Serialize(sortedItems, jsonOptions);
Console.WriteLine(output);