using System;
using System.Collections.Generic;
using System.IO;

var data = new List<Dictionary<string, object>>();

try
{
    var lines = File.ReadAllLines("input/data.csv");
    
    if (lines.Length > 1)
    {
        for (int i = 1; i < lines.Length; i++)
        {
            var columns = lines[i].Split(',');
            if (columns.Length == 4)
            {
                data.Add(new Dictionary<string, object>
                {
                    { "Name", columns[0] },
                    { "Age", int.Parse(columns[1]) },
                    { "Email", columns[2] },
                    { "Score", float.Parse(columns[3]) }
                });
            }
        }
    }

    var json = JsonSerializer.Serialize(new JsonArray(data), new JsonSerializerOptions { WriteIndented = true });
    Console.WriteLine(json);
}
catch (Exception ex)
{
    Console.WriteLine($"Error: {ex.Message}");
}