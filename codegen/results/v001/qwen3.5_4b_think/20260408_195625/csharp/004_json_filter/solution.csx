using System.Text.Json;
using System.Text.Json.Nodes;
using System.IO;
using System.Collections.Generic;

var data = await File.ReadAllTextAsync("input/data.json");
var root = JsonNode.Parse(data);

var filtered = new List<JsonNode>();

foreach (var item in root) {
    var jsonObject = item as JsonObject;
    if (jsonObject == null) continue;

    var active = jsonObject["active"]?.GetValue<bool>() ?? false;
    var age = jsonObject["age"]?.GetValue<int>() ?? 0;

    if (active && age >= 30) {
        filtered.Add(item);
    }
}

filtered.Sort((a, b) => string.Compare(a["name"].Value<string>(), b["name"].Value<string>(), StringComparison.Ordinal));

await System.Text.Json.JsonSerializer.SerializeAsync(filtered);