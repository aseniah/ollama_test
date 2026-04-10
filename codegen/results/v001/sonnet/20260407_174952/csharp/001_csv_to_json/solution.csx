using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var headers = lines[0].Split(',');

var array = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;

    var fields = lines[i].Split(',');
    var obj = new JsonObject
    {
        ["Name"] = fields[0].Trim(),
        ["Age"] = int.Parse(fields[1].Trim()),
        ["Email"] = fields[2].Trim(),
        ["Score"] = float.Parse(fields[3].Trim(), System.Globalization.CultureInfo.InvariantCulture)
    };
    array.Add(obj);
}

Console.WriteLine(array.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));
