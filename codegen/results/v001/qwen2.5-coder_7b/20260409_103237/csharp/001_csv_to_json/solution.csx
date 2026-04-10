using System;
using System.Collections.Generic;
using System.IO;

string[] lines = File.ReadAllLines("input/data.csv");
List<Dictionary<string, object>> data = new List<Dictionary<string, object>>();

for (int i = 1; i < lines.Length; i++)
{
    string[] row = lines[i].Split(',');
    if (row.Length == 4)
    {
        Dictionary<string, object> item = new Dictionary<string, object>
        {
            { "Name", row[0] },
            { "Age", int.Parse(row[1]) },
            { "Email", row[2] },
            { "Score", float.Parse(row[3]) }
        };
        data.Add(item);
    }
}

string json = JsonSerializer.Serialize(data, new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(json);