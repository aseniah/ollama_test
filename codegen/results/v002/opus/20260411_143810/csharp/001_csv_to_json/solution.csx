using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");
var nodes = new System.Text.Json.Nodes.JsonArray();

foreach (var line in lines.Skip(1))
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    var parts = line.Split(',');
    var obj = new System.Text.Json.Nodes.JsonObject
    {
        ["Name"] = parts[0].Trim(),
        ["Age"] = int.Parse(parts[1].Trim()),
        ["Email"] = parts[2].Trim(),
        ["Score"] = float.Parse(parts[3].Trim())
    };
    nodes.Add(obj);
}

Console.WriteLine(nodes.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));
