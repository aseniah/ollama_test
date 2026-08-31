using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Linq;

var jsonContent = File.ReadAllText("input/data.json");
var jsonArray = JsonArray.Parse(jsonContent);

var filtered = jsonArray
    .Select(item => item.AsObject())
    .Where(obj => obj["active"].GetValue<bool>() && obj["age"].GetValue<int>() >= 30)
    .OrderBy(obj => obj["name"].GetValue<string>())
    .Select(obj => new JsonObject
    {
        ["name"] = obj["name"].ToString(),
        ["age"] = obj["age"].ToString(),
        ["active"] = obj["active"].ToString(),
        ["score"] = obj["score"].ToString()
    });

Console.WriteLine(new JsonArray(filtered.ToArray()));