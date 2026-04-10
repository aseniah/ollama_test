using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var inputPath = "input/input.csv";
var examplePath = "input/expected_format.json";

if (!File.Exists(inputPath)) throw new FileNotFoundException("Input file not found", inputPath);

var example = JsonDocument.Parse(File.ReadAllText(examplePath)).Root;
if (example == null) throw new JsonException("Could not parse expected format file");

var csvContent = File.ReadAllText(inputPath);

var records = new List<Record>();
using var reader = new StreamReader(csvContent);
var header = reader.ReadLine()!.Split(',');

for (var i = 1; reader.ReadLine() is string rowLine; i++) {
    var values = rowLine.Split(',');
    var record = new Record {
        Id = int.Parse(values[0]!),
        Name = values[1]!,
        Date = DateTime.Parse(values[2]!),
        City = values[3]!
    };
    records.Add(record);
}

if (records.Count == 0) {
    Console.WriteLine("[]");
    return;
}

var referenceDate = new DateTime(2025, 7, 1);
var results = new List<JsonNode>();

foreach (var r in records) {
    var ageDays = referenceDate - r.Date;
    var years = Math.Floor(ageDays / 365.25);
    var months = (int)(ageDays / 30.44) % 12;
    var days = (int)(ageDays % 30.44);

    results.Add(JsonDocument.Parse($"{{\"id\":{r.Id}, \"name\":\"{r.Name}\", \"date\":\"{r.Date:yyyy-MM-dd}\", \"city\":\"{r.City}\", \"age\":{\"days\":{years},{months},{days}}}}")).Root;
}

Console.WriteLine(JsonDocument.Parse("[" + string.Join(",", results.Select(j => (JsonElement)j))) + "]");

class Record {
    public int Id { get; set; }
    public string Name { get; set; }
    public DateTime Date { get; set; }
    public string City { get; set; }
}