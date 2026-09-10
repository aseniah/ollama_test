using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/data.csv");

JsonArray jsonArray = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    string line = lines[i].Trim();
    if (line.Length == 0) continue;

    string[] fields = line.Split(',');

    JsonObject obj = new JsonObject
    {
        ["Name"] = fields[0].Trim(),
        ["Age"] = int.Parse(fields[1].Trim()),
        ["Email"] = fields[2].Trim(),
        ["Score"] = float.Parse(fields[3].Trim())
    };

    jsonArray.Add(obj);
}

JsonSerializerOptions options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.WriteLine(jsonArray.ToJsonString(options));