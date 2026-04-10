using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var array = new JsonArray();

foreach (var line in lines.Skip(1))
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    var cols = line.Split(',');
    var obj = new JsonObject
    {
        ["Name"] = cols[0].Trim(),
        ["Age"] = int.Parse(cols[1].Trim()),
        ["Email"] = cols[2].Trim(),
        ["Score"] = float.Parse(cols[3].Trim())
    };
    array.Add(obj);
}

Console.WriteLine(array.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));
