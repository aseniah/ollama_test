using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");
var rows = new List<object>();

for (int i = 1; i < lines.Length; i++)
{
    var parts = lines[i].Split(',');
    if (parts.Length < 4) continue;
    rows.Add(new
    {
        Name = parts[0].Trim(),
        Age = int.Parse(parts[1].Trim()),
        Email = parts[2].Trim(),
        Score = float.Parse(parts[3].Trim())
    });
}

Console.WriteLine(JsonSerializer.Serialize(rows));
