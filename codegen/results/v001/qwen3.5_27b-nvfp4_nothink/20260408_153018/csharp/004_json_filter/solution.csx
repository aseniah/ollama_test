#r "System.Text.Json"

using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var path = Path.Combine("input", "data.json");
var jsonContent = File.ReadAllText(path);
var root = JsonNode.Parse(jsonContent)!;
var array = root!.AsArray();

var result = array
    .Where(x =>
    {
        var obj = x!.AsObject();
        if (obj.TryGetValue("active", out var activeVal) &&
            activeVal?.GetValue<bool>() == true)
        {
            if (obj.TryGetValue("age", out var ageVal) &&
                ageVal != null &&
                ageVal.GetValue<int>() >= 30)
            {
                return true;
            }
        }
        return false;
    })
    .Select(x => x!.AsObject())
    .OrderBy(obj => obj["name"]?.GetValue<string>())
    .ToList();

var output = JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(output);