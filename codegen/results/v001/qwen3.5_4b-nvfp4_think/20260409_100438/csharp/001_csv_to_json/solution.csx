using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var filePath = "input/data.csv";
if (!File.Exists(filePath)) {
    throw new FileNotFoundException($"File not found: {filePath}");
}

var lines = File.ReadAllLines(filePath).ToList();
if (lines.Count == 0) {
    throw new Exception("CSV file is empty");
}

var headerLine = lines[0].TrimEnd('\n', '\r');
var cols = headerLine.Split(',').Select(c => c.Trim()).ToArray();

var dataObjects = lines.Skip(1).Select(line => {
    var values = line.Split(',');
    return new ObjectReader {
        Name = values[0],
        Age = int.Parse(values[1]),
        Email = values[2],
        Score = float.Parse(values[3])
    };
});

var jsonData = JsonDocument.CreateArray(
    dataObjects.Select(obj => new {
        Name = obj.Name,
        Age = obj.Age,
        Email = obj.Email,
        Score = obj.Score
    }).ToArray()
);

JsonSerializer.Serialize(new OutputFormatter(data: jsonData), System.Console.Out);