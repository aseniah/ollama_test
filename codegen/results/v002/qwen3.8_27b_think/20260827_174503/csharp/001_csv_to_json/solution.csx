using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/data.csv");

var array = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    string[] fields = lines[i].Split(',');
    if (fields.Length < 4) continue;

    var obj = new JsonObject
    {
        ["Name"] = fields[0],
        ["Age"] = int.Parse(fields[1]),
        ["Email"] = fields[2],
        ["Score"] = float.Parse(fields[3])
    };
    array.Add(obj);
}

string json = JsonSerializer.Serialize(array, new JsonSerializerOptions
{
    WriteIndented = false
});

Console.Write(json);