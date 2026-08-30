using System.Text.Json;

var lines = File.ReadAllLines("input/data.csv");
var rows = lines.Skip(1).Select(line => {
    var parts = line.Split(',');
    return new {
        Name = parts[0],
        Age = int.Parse(parts[1]),
        Email = parts[2],
        Score = float.Parse(parts[3])
    };
}).ToList();

Console.WriteLine(JsonSerializer.Serialize(rows));
