using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var array = new JsonArray();

foreach (var line in lines.Skip(1))
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    var parts = line.Split(',');
    var obj = new JsonObject
    {
        ["Name"] = parts[0],
        ["Age"] = int.Parse(parts[1]),
        ["Email"] = parts[2],
        ["Score"] = float.Parse(parts[3])
    };
    array.Add(obj);
}

using var writer = new System.Text.StringWriter();
array.WriteTo(writer, new System.Text.Json.JsonSerializerOptions { WriteIndented = true });
Console.Write(writer.ToString());