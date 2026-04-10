using System;
using System.IO;
using System.Text.Json;
using System.Collections.Generic;

var dataPath = "input/data.csv";
if (!File.Exists(dataPath))
{
    Console.Error.WriteLine($"File not found: {dataPath}");
    Environment.Exit(1);
}

var lines = File.ReadAllLines(dataPath);
if (lines.Length <= 1)
{
    Console.WriteLine($"No data found in {dataPath}.");
    Environment.Exit(1);
}

var data = new List<JsonElement>();
using var reader = new StreamReader(dataPath);
var headerLine = reader.ReadLine();

if (string.IsNullOrWhiteSpace(headerLine))
{
    Console.Error.WriteLine("No header line found.");
    Environment.Exit(1);
}

// Parse header (assuming order: Name, Age, Email, Score)
var headerColumns = headerLine.Split(',');

if (headerColumns.Length < 4)
{
    Console.Error.WriteLine($"Invalid header format. Expected at least 4 columns.");
    Environment.Exit(1);
}

for (int i = 1; i < lines.Length; i++)
{
    var rowLine = lines[i];
    if (string.IsNullOrWhiteSpace(rowLine))
    {
        continue;
    }

    var values = rowLine.Split(',');
    if (values.Length < 4)
    {
        continue;
    }

    var name = values[0].Trim();
    if (string.IsNullOrWhiteSpace(name))
    {
        continue;
    }

    int age;
    if (!int.TryParse(values[1].Trim(), out age))
    {
        continue;
    }

    var email = values[2].Trim();
    float score;
    if (!float.TryParse(values[3].Trim(), out score))
    {
        continue;
    }

    var obj = JsonNode.Create(new {
        Name = name,
        Age = age,
        Email = email,
        Score = score
    });
    data.Add(obj);
}

var options = new JsonSerializerOptions
{
    WriteIndented = true
};

Console.WriteLine(JsonSerializer.Serialize(data, options));