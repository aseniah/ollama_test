using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/data.csv");
var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (line.Length == 0) continue;

    var fields = line.Split(',');
    var obj = new JsonObject
    {
        ["Name"] = fields[0].Trim(),
        ["Age"] = int.Parse(fields[1].Trim()),
        ["Email"] = fields[2].Trim(),
        ["Score"] = float.Parse(fields[3].Trim(), System.Globalization.CultureInfo.InvariantCulture)
    };
    jsonArray.Add(obj);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.Write(jsonArray.ToJsonString(options));