using System;
using System.Text.Json;
using System.Text.Json.Nodes;

var data = JsonNode.Parse(File.ReadAllText("input/data.json"));
var jsonArray = (JsonArray)data;

var filtered = new List<JsonElement>();

foreach (var item in jsonArray) {
    if (item["active"] is not JsonValue activeObj) continue;
    var active = activeObj is true;

    if (item["age"] is not JsonValue ageObj) continue;
    var age = int.Parse(ageObj.ValueToString());

    if (active && age >= 30) {
        filtered.Add(item);
    }
}

filtered.Sort((a, b) => string.Compare(a["name"].ValueString, b["name"].ValueString, StringComparison.Ordinal));

Console.WriteLine(JsonSerializer.Serialize(filtered, new JsonSerializerOptions { WriteIndented = false }));