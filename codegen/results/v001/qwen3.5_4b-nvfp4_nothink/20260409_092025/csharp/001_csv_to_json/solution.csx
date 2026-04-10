using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

if (!File.Exists("input/data.csv")) {
    throw new FileNotFoundException("input/data.csv");
}

string content = File.ReadAllText("input/data.csv");
var lines = content.Split(new[] { Environment.NewLine }, StringSplitOptions.None);
if (lines.Length == 0) throw new InvalidOperationException("Empty file");

var names = new List<string>();
var ages = new List<int>();
var emails = new List<string>();
var scores = new List<float>();

foreach (string line in lines.Skip(1)) {
    if (string.IsNullOrWhiteSpace(line)) continue;
    var parts = line.Split(',', Environment.NewLine, StringSplitOptions.RemoveEmptyEntries);
    if (parts.Length < 4) continue;

    names.Add(parts[0]);
    ages.Add(int.Parse(parts[1], CultureInfo.InvariantCulture));
    emails.Add(parts[2]);
    scores.Add(float.Parse(parts[3], CultureInfo.InvariantCulture));
}

var results = new List<object>();
for (int i = 0; i < names.Count; i++) {
    var name = names[i];
    var age = ages[i];
    var email = emails[i];
    var score = scores[i];

    var obj = new JsonObject() {
        ["Name"] = name,
        ["Age"] = age,
        ["Email"] = email,
        ["Score"] = score,
    };
    results.Add(obj);
}

var jsonSerializerOptions = new JsonSerializerOptions { WriteIndented = false };
Console.WriteLine(JsonSerializer.Serialize(results, jsonSerializerOptions));