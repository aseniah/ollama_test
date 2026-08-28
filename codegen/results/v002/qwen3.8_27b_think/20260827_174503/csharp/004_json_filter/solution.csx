using System;
using System.IO;
using System.Linq;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
var array = JsonArray.Parse(json);

var filtered = array
    .Where(obj => {
        var o = (JsonObject)obj;
        return (bool)o["active"] && (int)o["age"] >= 30;
    })
    .OrderBy(obj => (string)((JsonObject)obj)["name"])
    .ToList();

var result = new JsonArray();
foreach (var obj in filtered)
    result.Add(obj);

Console.WriteLine(result.ToJsonString());