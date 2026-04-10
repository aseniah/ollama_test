using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var content = File.ReadAllText("input/data.json");
var rootArray = JsonArray.Parse(content);

var filteredRecords = rootArray
    .OfType<JsonObject>()
    .Where(x => {
        bool? isActive = x["active"]?.GetValue<bool>();
        int? age = x["age"]?.GetValue<int>();
        return isActive == true && (age ?? 0) >= 30;
    })
    .OrderBy(x => x["name"]?.ToString() ?? "")
    .ToList();

Console.WriteLine(JsonSerializer.Serialize(filteredRecords));