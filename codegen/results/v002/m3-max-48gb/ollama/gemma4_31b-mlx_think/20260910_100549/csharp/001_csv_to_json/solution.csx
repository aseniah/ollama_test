using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Globalization;

// Read all lines from the specified CSV file
string filePath = "input/data.csv";
string[] lines = File.ReadAllLines(filePath);

// Skip the header and map each row to an anonymous object
var data = lines.Skip(1)
    .Where(line => !string.IsNullOrWhiteSpace(line))
    .Select(line => 
    {
        string[] columns = line.Split(',');
        return new 
        {
            Name = columns[0],
            Age = int.Parse(columns[1]),
            Email = columns[2],
            Score = float.Parse(columns[3], CultureInfo.InvariantCulture)
        };
    })
    .ToList();

// Serialize the list to a JSON string and output to stdout
string jsonOutput = JsonSerializer.Serialize(data);
Console.WriteLine(jsonOutput);