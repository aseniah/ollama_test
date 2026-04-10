using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

var path = Args[0] ?? "input/data.csv";
var lines = File.ReadAllLines(path);
if (lines.Length < 2) throw new InvalidOperationException("CSV must have at least one data row");

var data = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    if (!line.Trim().Empty())
    {
        var cols = line.Split(',');
        if (cols.Length >= 4)
        {
            data.Add(new Dictionary<string, object>
            {
                ["Name"] = cols[0].Trim(),
                ["Age"] = int.Parse(cols[1]),
                ["Email"] = cols[2].Trim(),
                ["Score"] = float.Parse(cols[3])
            });
        }
    }
}

var result = System.Text.Json.JsonSerializer.Serialize(data, new JsonSerializerOptions { WriteIndented = false });
Console.WriteLine(result);