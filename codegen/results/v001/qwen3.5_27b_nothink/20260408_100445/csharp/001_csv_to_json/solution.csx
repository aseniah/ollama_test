using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Serialization;

var lines = File.ReadAllLines("input/data.csv");
var result = new List<object>();

if (lines.Length > 1)
{
    for (int i = 1; i < lines.Length; i++)
    {
        var parts = lines[i].Split(',');
        if (parts.Length >= 4)
        {
            var name = parts[0].Trim();
            int age = int.Parse(parts[1].Trim());
            var email = parts[2].Trim();
            float score = float.Parse(parts[3].Trim(), System.Globalization.CultureInfo.InvariantCulture);

            result.Add(new
            {
                Name = name,
                Age = age,
                Email = email,
                Score = score
            });
        }
    }
}

Console.WriteLine(JsonSerializer.Serialize(result));