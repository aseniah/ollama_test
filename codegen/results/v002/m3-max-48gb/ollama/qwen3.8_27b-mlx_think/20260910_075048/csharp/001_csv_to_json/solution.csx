using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;

var lines = File.ReadAllLines("input/data.csv");

var jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;

    var fields = line.Split(',');
    var obj = new JsonObject
    {
        ["Name"] = fields[0].Trim(),
        ["Age"] = int.Parse(fields[1].Trim()),
        ["Email"] = fields[2].Trim(),
        ["Score"] = float.Parse(fields[3].Trim(), CultureInfo.InvariantCulture)
    };
    jsonArray.Add(obj);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.WriteLine(jsonArray.ToJsonString(options));